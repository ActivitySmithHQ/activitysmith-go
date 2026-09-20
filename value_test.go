package activitysmith

import (
	"encoding/json"
	"testing"
)

func TestValueSerialization(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()
	client, _ := New("test")
	overrideHostForTests(client, server.URL)
	for _, value := range []any{"$1,240", "0007", "", 0, -12.75} {
		operations := []func() error{
			func() error {
				_, err := client.LiveActivities.Start(LiveActivityStartInput{Title: "Revenue", Type: LiveActivityTypeValue, Value: value})
				return err
			},
			func() error {
				_, err := client.LiveActivities.Update(LiveActivityUpdateInput{ActivityID: "a", Title: "Revenue", Type: LiveActivityTypeValue, Value: value})
				return err
			},
			func() error {
				_, err := client.LiveActivities.End(LiveActivityEndInput{ActivityID: "a", Title: "Revenue", Type: LiveActivityTypeValue, Value: value})
				return err
			},
			func() error {
				_, err := client.LiveActivities.Stream("prod-web-1", LiveActivityStreamInput{Title: "Revenue", Type: LiveActivityTypeValue, Value: value})
				return err
			},
			func() error {
				_, err := client.LiveActivities.EndStream("prod-web-1", LiveActivityStreamEndInput{Title: "Revenue", Type: LiveActivityTypeValue, Value: value})
				return err
			},
		}
		for _, operation := range operations {
			if err := operation(); err != nil {
				t.Fatal(err)
			}
			var body map[string]any
			if err := json.Unmarshal([]byte((*requests)[len(*requests)-1].Body), &body); err != nil {
				t.Fatal(err)
			}
			got, _ := json.Marshal(body["content_state"].(map[string]any)["value"])
			want, _ := json.Marshal(value)
			if string(got) != string(want) {
				t.Fatalf("got %s; want %s", got, want)
			}
		}
	}
}
