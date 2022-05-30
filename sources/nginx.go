package sources

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/overmindtech/discovery"
	"github.com/overmindtech/nginx-source/crossplane"
	"github.com/overmindtech/nginx-source/triggers"
	"github.com/overmindtech/sdp-go"
	"google.golang.org/protobuf/types/known/durationpb"
)

type NginxSource struct {
	Engine *discovery.Engine
}

// Type The type of items that this source is capable of finding
func (s *NginxSource) Type() string {
	return "nginx"
}

// Descriptive name for the source, used in logging and metadata
func (s *NginxSource) Name() string {
	return "nginx-source"
}

// List of contexts that this source is capable of find items for. If the
// source supports all contexts the special value `AllContexts` ("*")
// should be used
func (s *NginxSource) Contexts() []string {
	return []string{
		sdp.WILDCARD,
	}
}

// Get This method does not work for Nginx as there iosn't really any specific
// thing that differentiates one Nginx instance from another, the
// UniqueAttribute I'm using as the SHA1 hash of all of the server's confuration
// as this is all I could think of that would ensure that a server was unique
func (s *NginxSource) Get(ctx context.Context, itemContext string, query string) (*sdp.Item, error) {
	return nil, &sdp.ItemRequestError{
		ErrorType:   sdp.ItemRequestError_OTHER,
		ErrorString: "Get is not supported for this source since nginx instances don't have a truly unique identifier. Use Search instead",
		Context:     itemContext,
	}
}

// Find Not supported for this source
func (s *NginxSource) Find(ctx context.Context, itemContext string) ([]*sdp.Item, error) {
	return []*sdp.Item{}, nil
}

// Search Looks for nginx instance in a given context. The query should be one
// of the following:
//
// * `service-linux`: This looks for nginx on a linux server which has the
// nginx service running
//
func (s *NginxSource) Search(ctx context.Context, itemContext string, query string) ([]*sdp.Item, error) {
	var triggerData triggers.TriggerData

	err := json.Unmarshal([]byte(query), &triggerData)

	if err != nil {
		return nil, &sdp.ItemRequestError{
			ErrorType:   sdp.ItemRequestError_OTHER,
			ErrorString: err.Error(),
			Context:     itemContext,
		}
	}

	switch triggerData.TriggerType {
	case triggers.SERVICE:
		// I possibly don't even need to look for the config file etc. I could
		// literally just run the following and parse out the info:
		//
		// * `nginx -V`: Version info and many arguments
		// * `nginx -Tq`: All config concatenated perfectly for crossplane

		// Extract the arguments that are being passed to nginx from the service
		// as well as the path to the binary itself
		versionCommand := fmt.Sprintf("%v -V %v", triggerData.ServiceData.Binary, strings.Join(triggerData.ServiceData.Args, " "))
		configCommand := fmt.Sprintf("%v -Tq", triggerData.ServiceData.Binary)
		versionUUID := uuid.New()
		configUUID := uuid.New()

		versionRequest := sdp.ItemRequest{
			Type:            "command",
			Method:          sdp.RequestMethod_GET,
			Query:           versionCommand,
			LinkDepth:       0,
			Context:         itemContext,
			Timeout:         durationpb.New(10 * time.Second),
			ItemSubject:     discovery.NewItemSubject(),
			ResponseSubject: discovery.NewResponseSubject(),
			IgnoreCache:     false,
			UUID:            versionUUID[:],
		}

		configRequest := sdp.ItemRequest{
			Type:            "command",
			Method:          sdp.RequestMethod_GET,
			Query:           configCommand,
			LinkDepth:       0,
			Context:         itemContext,
			IgnoreCache:     false,
			UUID:            configUUID[:],
			Timeout:         durationpb.New(10 * time.Second),
			ItemSubject:     discovery.NewItemSubject(),
			ResponseSubject: discovery.NewResponseSubject(),
		}

		wg := sync.WaitGroup{}

		wg.Add(2)

		var versionItems []*sdp.Item
		var versionErr error
		var configItems []*sdp.Item
		var configErr error

		// Execute both requests
		go func() {
			defer wg.Done()
			progress := sdp.RequestProgress{
				StartTimeout: 10 * time.Second,
				Request:      &versionRequest,
				Responders:   make(map[string]*sdp.Responder),
			}

			versionItems, versionErr = progress.Execute(s.Engine.ManagedConnection())
		}()

		go func() {
			defer wg.Done()

			progress := sdp.RequestProgress{
				StartTimeout: 10 * time.Second,
				Request:      &configRequest,
				Responders:   make(map[string]*sdp.Responder),
			}

			configItems, configErr = progress.Execute(s.Engine.ManagedConnection())
		}()

		wg.Wait()

		if configErr != nil || versionErr != nil {
			return []*sdp.Item{}, &sdp.ItemRequestError{
				ErrorType:   sdp.ItemRequestError_OTHER,
				ErrorString: fmt.Sprintf("error gathering nginx info: %v%v", configErr, versionErr),
				Context:     itemContext,
			}
		}

		// Validate that we got one item for each and extract the useful data
		if len(versionItems) != 1 {
			return []*sdp.Item{}, &sdp.ItemRequestError{
				ErrorType:   sdp.ItemRequestError_NOTFOUND,
				ErrorString: fmt.Sprintf("error gathering nginx info: expected 1 version item but got %v", len(versionItems)),
				Context:     itemContext,
			}
		}

		if len(configItems) != 1 {
			return []*sdp.Item{}, &sdp.ItemRequestError{
				ErrorType:   sdp.ItemRequestError_NOTFOUND,
				ErrorString: fmt.Sprintf("error gathering nginx info: expected 1 version item but got %v", len(configItems)),
				Context:     itemContext,
			}
		}

		versionItem := versionItems[0]
		configItem := configItems[0]
		attrMap := make(map[string]interface{})

		if stderr, err := versionItem.Attributes.Get("stderr"); err == nil {
			versionInfo := parseVersionInfo(fmt.Sprint(stderr))

			attrMap["version"] = versionInfo.Version
			attrMap["builtBy"] = versionInfo.BuiltBy
			attrMap["openSSL"] = versionInfo.OpenSSL
			attrMap["configArgs"] = versionInfo.ConfigArgs
		}

		if stdout, err := configItem.Attributes.Get("stdout"); err == nil {
			resp, err := crossplane.Parse(ctx, fmt.Sprint(stdout))

			if err != nil {
				return []*sdp.Item{}, &sdp.ItemRequestError{
					ErrorType:   sdp.ItemRequestError_OTHER,
					ErrorString: fmt.Sprintf("error parsing nginx config: %v", err),
					Context:     itemContext,
				}
			}

			attrMap["config"] = resp.Config

			shaSum := sha1.Sum([]byte(fmt.Sprint(stdout)))
			shaString := base64.URLEncoding.EncodeToString(shaSum[:])

			attrMap["configHash"] = shaString
		}

		attributes, err := sdp.ToAttributes(attrMap)

		if err != nil {
			return []*sdp.Item{}, &sdp.ItemRequestError{
				ErrorType:   sdp.ItemRequestError_OTHER,
				ErrorString: fmt.Sprintf("error converting to attributes: %v", err),
				Context:     itemContext,
			}
		}

		item := sdp.Item{
			Type:            "nginx",
			UniqueAttribute: "configHash", // This is the SHA1 of all of the config
			Attributes:      attributes,
			Context:         itemContext,
		}

		if triggerData.TriggerItemRef != nil {
			item.LinkedItems = append(item.LinkedItems, triggerData.TriggerItemRef)
		}

		return []*sdp.Item{&item}, nil
	default:
		return []*sdp.Item{}, &sdp.ItemRequestError{
			ErrorType:   sdp.ItemRequestError_NOTFOUND,
			ErrorString: fmt.Sprintf("query %v not supported", query),
			Context:     itemContext,
		}
	}
}

// Weight Returns the priority weighting of items returned by this source.
// This is used to resolve conflicts where two sources of the same type
// return an item for a GET request. In this instance only one item can be
// sen on, so the one with the higher weight value will win.
func (s *NginxSource) Weight() int {
	return 100
}

type NginxVersionInfo struct {
	Version    string
	BuiltBy    string
	OpenSSL    string
	ConfigArgs []string
}

var versionRegex = regexp.MustCompile(`nginx version:\s+(\S+)`)
var builtByRegex = regexp.MustCompile(`built by\s+(.*)`)
var opensslRegex = regexp.MustCompile(`built with OpenSSL\s+(\S+)`)
var configArgsRegex = regexp.MustCompile(`configure arguments: (.*)`)
var eachArgRegex = regexp.MustCompile(`(-(\S+'.*?'|\S+)\s??)`)

// parseVersionInfo Parses version information from `nginx -V`
func parseVersionInfo(infoString string) NginxVersionInfo {
	versionInfo := NginxVersionInfo{}

	if matches := versionRegex.FindStringSubmatch(infoString); len(matches) == 2 {
		versionInfo.Version = matches[1]
	}

	if matches := builtByRegex.FindStringSubmatch(infoString); len(matches) == 2 {
		versionInfo.BuiltBy = matches[1]
	}

	if matches := opensslRegex.FindStringSubmatch(infoString); len(matches) == 2 {
		versionInfo.OpenSSL = matches[1]
	}

	if matches := configArgsRegex.FindStringSubmatch(infoString); len(matches) == 2 {
		if argMatches := eachArgRegex.FindAllStringSubmatch(matches[1], -1); len(argMatches) > 0 {
			for _, thisArg := range argMatches {
				if len(thisArg) >= 2 {
					versionInfo.ConfigArgs = append(versionInfo.ConfigArgs, thisArg[1])
				}
			}
		}
	}

	return versionInfo
}
