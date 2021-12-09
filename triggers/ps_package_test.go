package triggers

import (
	"encoding/json"
	"testing"

	"github.com/overmindtech/sdp-go"
)

func TestServiceTrigger(t *testing.T) {
	t.Run("with a matching service item", func(t *testing.T) {
		attr, _ := sdp.ToAttributes(map[string]interface{}{
			"ExecStart": map[string]interface{}{
				"args": []string{
					"-c",
					"/etc/nginx/nginx.conf",
				},
				"binary":  "/usr/sbin/nginx",
				"fullCMD": "/usr/sbin/nginx -c /etc/nginx/nginx.conf",
			},
			"Name": "nginx.service",
		})

		item := sdp.Item{
			Type:            "service",
			UniqueAttribute: "name",
			Attributes:      attr,
			Context:         "test",
		}

		req, err := ServiceTrigger.RequestGenerator(&item)

		if err != nil {
			t.Fatal(err)
		}

		if expected := "nginx"; req.Type != expected {
			t.Errorf("expected req.Type to be %v, got %v", expected, req.Type)
		}

		if expected := sdp.RequestMethod_SEARCH; req.Method != expected {
			t.Errorf("expected req.Method to be %v, got %v", expected, req.Method)
		}

		var td TriggerData

		err = json.Unmarshal([]byte(req.Query), &td)

		if err != nil {
			t.Error(err)
		}

		if expected := SERVICE; td.TriggerType != expected {
			t.Errorf("expected td.TriggerType to be %v, got %v", expected, td.TriggerType)
		}

		if td.ServiceData == nil {
			t.Error("td.ServiceData is nil")
		}

		if expected := "/usr/sbin/nginx"; td.ServiceData.Binary != expected {
			t.Errorf("expected td.ServiceData.Binary to be %v, got %v", expected, td.ServiceData.Binary)
		}

		if len(td.ServiceData.Args) != 2 {
			t.Fatalf("td.ServiceData should have 2 args, got %v", td.ServiceData.Args)
		}

		if expected := "-c"; td.ServiceData.Args[0] != expected {
			t.Errorf("expected td.ServiceData.Args[0] to be %v, got %v", expected, td.ServiceData.Args[0])
		}

		if expected := "/etc/nginx/nginx.conf"; td.ServiceData.Args[1] != expected {
			t.Errorf("expected td.ServiceData.Args[1] to be %v, got %v", expected, td.ServiceData.Args[1])
		}
	})

	t.Run("with an incomplete service item", func(t *testing.T) {
		attr, _ := sdp.ToAttributes(map[string]interface{}{
			"Name": "nginx.service",
		})

		item := sdp.Item{
			Type:            "service",
			UniqueAttribute: "name",
			Attributes:      attr,
			Context:         "test",
		}

		req, err := ServiceTrigger.RequestGenerator(&item)

		if err != nil {
			t.Fatal(err)
		}

		if expected := "nginx"; req.Type != expected {
			t.Errorf("expected req.Type to be %v, got %v", expected, req.Type)
		}

		if expected := sdp.RequestMethod_SEARCH; req.Method != expected {
			t.Errorf("expected req.Method to be %v, got %v", expected, req.Method)
		}

		var td TriggerData

		err = json.Unmarshal([]byte(req.Query), &td)

		if err != nil {
			t.Error(err)
		}

		if expected := SERVICE; td.TriggerType != expected {
			t.Errorf("expected td.TriggerType to be %v, got %v", expected, td.TriggerType)
		}

		if td.ServiceData == nil {
			t.Error("td.ServiceData is nil")
		}

		if expected := "nginx"; td.ServiceData.Binary != expected {
			t.Errorf("expected td.ServiceData.Binary to be %v, got %v", expected, td.ServiceData.Binary)
		}

		if len(td.ServiceData.Args) != 0 {
			t.Fatalf("td.ServiceData should have 0 args, got %v", td.ServiceData.Args)
		}
	})

	t.Run("with a service item that has incompatible attributes", func(t *testing.T) {
		attr, _ := sdp.ToAttributes(map[string]interface{}{
			"ExecStart": map[string]interface{}{
				"args":    "-c /etc/nginx/nginx.conf",
				"binary":  "/usr/sbin/nginx",
				"fullCMD": "/usr/sbin/nginx -c /etc/nginx/nginx.conf",
			},
			"Name": "nginx.service",
		})

		item := sdp.Item{
			Type:            "service",
			UniqueAttribute: "name",
			Attributes:      attr,
			Context:         "test",
		}

		req, err := ServiceTrigger.RequestGenerator(&item)

		if err != nil {
			t.Fatal(err)
		}

		if expected := "nginx"; req.Type != expected {
			t.Errorf("expected req.Type to be %v, got %v", expected, req.Type)
		}

		if expected := sdp.RequestMethod_SEARCH; req.Method != expected {
			t.Errorf("expected req.Method to be %v, got %v", expected, req.Method)
		}

		var td TriggerData

		err = json.Unmarshal([]byte(req.Query), &td)

		if err != nil {
			t.Error(err)
		}

		if expected := SERVICE; td.TriggerType != expected {
			t.Errorf("expected td.TriggerType to be %v, got %v", expected, td.TriggerType)
		}

		if td.ServiceData == nil {
			t.Error("td.ServiceData is nil")
		}

		if expected := "/usr/sbin/nginx"; td.ServiceData.Binary != expected {
			t.Errorf("expected td.ServiceData.Binary to be %v, got %v", expected, td.ServiceData.Binary)
		}

		if len(td.ServiceData.Args) != 0 {
			t.Fatalf("td.ServiceData should have 0 args, got %v", td.ServiceData.Args)
		}
	})
}
