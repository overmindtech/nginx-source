package triggers

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/overmindtech/discovery"
	"github.com/overmindtech/sdp-go"
)

type TriggerType int

const (
	SERVICE TriggerType = iota
)

// Data that will be sent to the Search() method
type TriggerData struct {
	TriggerType TriggerType  `json:"trigger_type,omitempty"`
	ServiceData *ServiceData `json:"service_data,omitempty"`
}

// Data required if the TriggerType is "service"
type ServiceData struct {
	Args   []string `json:"args,omitempty"`
	Binary string   `json:"binary,omitempty"`
}

var AllTriggers = []discovery.Trigger{
	ServiceTrigger,
}

// TODO: At this point I need to modify the trrigger so that is passes the
// relevant information that Search is going to need. The plan is that I pass in
// the command line arguments so that we can then provide those to the `nginx -V
// ` and `nginx -T` commands and use that to get the details that we need

// This trigger is based on service resources such as on linux
//
var ServiceTrigger = discovery.Trigger{
	Type:                      "service",
	UniqueAttributeValueRegex: regexp.MustCompile(`^nginx$`),
	RequestGenerator: func(in *sdp.Item) (*sdp.ItemRequest, error) {
		var args []string
		var binary string
		var argsInterface interface{}
		var binaryInterface interface{}
		var err error

		// Try to pull the arguments from the item
		argsInterface, err = in.Attributes.Get("ExecStart.args")

		if err != nil {
			// Default to empty
			args = make([]string, 0)
		} else {
			if argsInterfaceSlice, ok := argsInterface.([]interface{}); ok {
				for _, arg := range argsInterfaceSlice {
					args = append(args, fmt.Sprint(arg))
				}
			}
		}

		// Try to pull the binary from the service too
		err = nil
		binaryInterface, err = in.Attributes.Get("ExecStart.binary")

		if err != nil {
			binary = "nginx"
		} else {
			binary = fmt.Sprint(binaryInterface)

			if binary == "" {
				binary = "nginx"
			}
		}

		td := TriggerData{
			TriggerType: SERVICE,
			ServiceData: &ServiceData{
				Args:   args,
				Binary: binary,
			},
		}

		// Marshall to JSON and send
		b, err := json.Marshal(td)

		if err != nil {
			return nil, err
		}

		return &sdp.ItemRequest{
			Type:   "nginx",
			Method: sdp.RequestMethod_SEARCH,
			Query:  string(b), // Find details using service mode
		}, nil
	},
}
