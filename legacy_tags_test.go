package activitysmith

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestLegacyTagsSerialization(t *testing.T) {
	for _, tags := range [][]string{nil, {"billing"}, {}} {
		server, requests := newAPITestServer(t)
		client, err := New("test")
		if err != nil {
			t.Fatal(err)
		}
		overrideHostForTests(client, server.URL)
		state := LiveActivityContentStateInput{Title: "Job"}
		_, err = client.LiveActivities.Update(LiveActivityUpdateInput{ActivityID: "activity-1", ContentState: state, Tags: tags})
		if err != nil {
			t.Fatal(err)
		}
		_, err = client.LiveActivities.End(LiveActivityEndInput{ActivityID: "activity-1", ContentState: state, Tags: tags})
		if err != nil {
			t.Fatal(err)
		}
		for _, request := range *requests {
			var body map[string]json.RawMessage
			if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
				t.Fatal(err)
			}
			raw, exists := body["tags"]
			if exists != (tags != nil) {
				t.Fatalf("Tags presence: %s", request.Body)
			}
			if exists {
				var actual []string
				if err := json.Unmarshal(raw, &actual); err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(actual, tags) {
					t.Fatalf("Tags = %s, want %v", raw, tags)
				}
			}
		}
		server.Close()
	}
}
