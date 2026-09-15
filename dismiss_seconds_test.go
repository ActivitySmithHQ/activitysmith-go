package activitysmith

import (
	"encoding/json"
	"testing"
)

func TestStreamDismissSeconds(t *testing.T) {
	cases := []struct {
		name  string
		state LiveActivityContentStateInput
		want  *float64
	}{
		{"omitted", LiveActivityContentStateInput{Title: "Done"}, nil},
		{"field", LiveActivityContentStateInput{Title: "Done", AutoDismissSeconds: 30}, secondsPointer(30)},
		{"zero", (LiveActivityContentStateInput{Title: "Done"}).WithAutoDismissSeconds(0), secondsPointer(0)},
		{"both", (LiveActivityContentStateInput{Title: "Done", AutoDismissMinutes: 5}).WithAutoDismissSeconds(30), secondsPointer(30)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			server, requests := newAPITestServer(t)
			defer server.Close()
			client, err := New("test")
			if err != nil {
				t.Fatal(err)
			}
			overrideHostForTests(client, server.URL)
			_, err = client.LiveActivities.Stream("prod-web-1", LiveActivityStreamInput{ContentState: tc.state})
			if err != nil {
				t.Fatal(err)
			}
			_, err = client.LiveActivities.EndStream("prod-web-1", LiveActivityStreamEndInput{ContentState: tc.state})
			if err != nil {
				t.Fatal(err)
			}
			for _, request := range *requests {
				var body struct {
					ContentState map[string]any `json:"content_state"`
				}
				if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
					t.Fatal(err)
				}
				value, exists := body.ContentState["auto_dismiss_seconds"]
				if tc.want == nil {
					if exists {
						t.Fatalf("unexpected seconds: %v", value)
					}
				} else if !exists || value != *tc.want {
					t.Fatalf("seconds = %v, want %v", value, *tc.want)
				}
				if tc.name == "both" && body.ContentState["auto_dismiss_minutes"] != float64(5) {
					t.Fatal("minutes must remain present when seconds are supplied")
				}
			}
		})
	}
	if !(LiveActivityContentStateInput{}).WithAutoDismissSeconds(0).isSet() {
		t.Fatal("explicit zero must count as a supplied content state")
	}
}

func secondsPointer(v float64) *float64 { return &v }
