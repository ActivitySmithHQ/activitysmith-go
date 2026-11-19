package activitysmith

import (
	"encoding/json"
	"testing"
)

func TestSoundSerialization(t *testing.T) {
	for _, sound := range []string{"", "chime", "default", "none"} {
		server, requests := newAPITestServer(t)
		client, err := New("test")
		if err != nil {
			t.Fatal(err)
		}
		overrideHostForTests(client, server.URL)
		state := LiveActivityContentStateInput{Title: "Deploy", Type: LiveActivityTypeProgress}
		calls := []struct {
			name string
			send func() error
		}{
			{"send", func() error {
				_, err := client.Notifications.Send(PushNotificationInput{Title: "Build Failed", Sound: sound})
				return err
			}},
			{"start", func() error {
				_, err := client.LiveActivities.Start(LiveActivityStartInput{ContentState: state, Sound: sound})
				return err
			}},
			{"update", func() error {
				_, err := client.LiveActivities.Update(LiveActivityUpdateInput{ActivityID: "activity-1", ContentState: state, Sound: sound})
				return err
			}},
			{"end", func() error {
				_, err := client.LiveActivities.End(LiveActivityEndInput{ActivityID: "activity-1", ContentState: state, Sound: sound})
				return err
			}},
			{"stream", func() error {
				_, err := client.LiveActivities.Stream("prod-web-1", LiveActivityStreamInput{ContentState: state, Sound: sound})
				return err
			}},
			{"end stream", func() error {
				_, err := client.LiveActivities.EndStream("prod-web-1", LiveActivityStreamEndInput{Sound: sound})
				return err
			}},
		}
		for _, call := range calls {
			if err := call.send(); err != nil {
				t.Fatalf("%s returned error: %v", call.name, err)
			}
		}
		if len(*requests) != len(calls) {
			t.Fatalf("expected %d requests, got %d", len(calls), len(*requests))
		}
		for index, request := range *requests {
			var body map[string]any
			if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
				t.Fatal(err)
			}
			value, exists := body["sound"]
			if sound == "" {
				if exists {
					t.Fatalf("%s: empty Sound serialized: %s", calls[index].name, request.Body)
				}
			} else if value != sound {
				t.Fatalf("%s: sound = %v, want %q: %s", calls[index].name, value, sound, request.Body)
			}
		}
		server.Close()
	}
}
