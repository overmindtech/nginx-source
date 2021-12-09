package sources

import (
	"regexp"
	"testing"

	"github.com/overmindtech/sdp-go"
)

// This file contains tests for the ColourNameSource source. It is a good idea
// to write as many exhaustive tests as possible at this level to ensure that
// your source responds correctly to certain requests.
func TestGet(t *testing.T) {
	tests := []SourceTest{
		{
			Name:        "get should fail",
			ItemContext: "something.specific",
			Query:       "irrelevant",
			Method:      sdp.RequestMethod_GET,
			ExpectedError: &ExpectedError{
				Type:             sdp.ItemRequestError_OTHER,
				ErrorStringRegex: regexp.MustCompile(`Get is not supported`),
				Context:          "something.specific",
			},
		},
	}

	RunSourceTests(t, tests, &NginxSource{})
}

func TestFind(t *testing.T) {
	tests := []SourceTest{
		{
			Name:          "find returns no items",
			ItemContext:   "something.specific",
			Method:        sdp.RequestMethod_FIND,
			ExpectedError: nil,
			ExpectedItems: &ExpectedItems{
				NumItems: 0,
			},
		},
	}

	RunSourceTests(t, tests, &NginxSource{})
}

func TestSearch(t *testing.T) {

	tests := []SourceTest{
		{
			Name:        "with no query",
			ItemContext: "something.specific",
			Method:      sdp.RequestMethod_SEARCH,
			ExpectedError: &ExpectedError{
				Type:    sdp.ItemRequestError_OTHER,
				Context: "something.specific",
			},
		},
	}

	RunSourceTests(t, tests, &NginxSource{})
}
