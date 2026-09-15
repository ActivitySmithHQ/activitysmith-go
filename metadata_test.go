package activitysmith

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMetadataSerialization(t *testing.T) {
	for _, metadata := range []map[string]any{nil, {}, {"order": "382", "ready": false, "count": float64(0), "empty": "", "ratio": 1.25}} {
		server, requests := newAPITestServer(t)
		client, _ := New("test")
		overrideHostForTests(client, server.URL)
		state := LiveActivityContentStateInput{Title: "Job"}
		operations := []func() error{
			func() error {
				_, e := client.LiveActivities.EndStream("prod-web-1", LiveActivityStreamEndInput{Metadata: metadata})
				return e
			},
			func() error {
				_, e := client.Notifications.Send(PushNotificationInput{Title: "Job", Metadata: metadata})
				return e
			},
			func() error {
				_, e := client.LiveActivities.Start(LiveActivityStartInput{ContentState: state, Metadata: metadata})
				return e
			},
			func() error {
				_, e := client.LiveActivities.Update(LiveActivityUpdateInput{ActivityID: "a", ContentState: state, Metadata: metadata})
				return e
			},
			func() error {
				_, e := client.LiveActivities.End(LiveActivityEndInput{ActivityID: "a", ContentState: state, Metadata: metadata})
				return e
			},
			func() error {
				_, e := client.LiveActivities.Stream("prod-web-1", LiveActivityStreamInput{ContentState: state, Metadata: metadata})
				return e
			},
		}
		for _, operation := range operations {
			if err := operation(); err != nil {
				t.Fatal(err)
			}
		}
		for _, request := range *requests {
			var body map[string]any
			if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
				t.Fatal(err)
			}
			value, exists := body["metadata"]
			if exists != (metadata != nil) {
				t.Fatalf("metadata presence: %s", request.Body)
			}
			if exists && !reflect.DeepEqual(value, metadata) {
				t.Fatalf("metadata = %#v, want %#v", value, metadata)
			}
		}
		server.Close()
	}
}

func TestMetadataRejectsNonScalarValues(t *testing.T) {
	for _, value := range []any{nil, map[string]any{}, []string{"nested"}} {
		if _, err := normalizeLiveActivityUpdateRequest(LiveActivityUpdateInput{Metadata: map[string]any{"field": value}}); err == nil {
			t.Fatalf("accepted %v", value)
		}
	}
}

func TestEndStreamTagsAndExternalPushURLs(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()
	client, _ := New("test")
	overrideHostForTests(client, server.URL)
	for _, tags := range [][]string{nil, {}, {"finished"}} {
		_, err := client.LiveActivities.EndStream("prod-web-1", LiveActivityStreamEndInput{Tags: tags, Metadata: map[string]any{}})
		if err != nil {
			t.Fatal(err)
		}
		var body map[string]any
		json.Unmarshal([]byte((*requests)[len(*requests)-1].Body), &body)
		_, present := body["tags"]
		if present != (tags != nil) {
			t.Fatalf("tags presence: %#v", body)
		}
		if tags != nil && len(body["tags"].([]any)) != len(tags) {
			t.Fatalf("tags: %#v", body)
		}
	}
	for _, url := range []string{"http://example.com", "https://example.com", "shortcuts://run-shortcut?name=Test", "spotify://", "spotify:track:123"} {
		_, err := client.Notifications.Send(PushNotificationInput{Title: "Job", Redirection: url})
		if err != nil {
			t.Fatal(err)
		}
		var body map[string]any
		json.Unmarshal([]byte((*requests)[len(*requests)-1].Body), &body)
		if body["redirection"] != url {
			t.Fatalf("redirection: %#v", body)
		}
	}
}
