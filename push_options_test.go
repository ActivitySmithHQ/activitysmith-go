package activitysmith

import (
	"encoding/json"
	"testing"

	"github.com/ActivitySmithHQ/activitysmith-go/generated"
)

func TestPushOptionsSerialization(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()
	client, err := New("mock-only")
	if err != nil {
		t.Fatal(err)
	}
	client.apiClient.GetConfig().Servers = generated.ServerConfigurations{{URL: server.URL}}
	for _, withIcon := range []bool{false, true} {
		for _, level := range []PushInterruptionLevel{"", PushInterruptionLevelPassive, PushInterruptionLevelActive, PushInterruptionLevelTimeSensitive} {
			input := PushNotificationInput{Title: "GitHub", Subtitle: "Build status", Channels: []string{"ops"}, Tags: []string{"ci"}, InterruptionLevel: level}
			if withIcon {
				input.Icon = "https://example.com/github.png"
			}
			req := input.toGenerated()
			for _, form := range []any{input, &input, req, &req} {
				if _, err := client.Notifications.Send(form); err != nil {
					t.Fatal(err)
				}
				var body map[string]any
				if err := json.Unmarshal([]byte((*requests)[len(*requests)-1].Body), &body); err != nil {
					t.Fatal(err)
				}
				if body["title"] != input.Title || body["subtitle"] != input.Subtitle {
					t.Fatalf("lost fields: %v", body)
				}
				if level == "" {
					if _, ok := body["interruption_level"]; ok {
						t.Fatal("omitted level serialized")
					}
				} else if body["interruption_level"] != string(level) {
					t.Fatalf("wrong options: %v", body)
				}
				if withIcon {
					if body["icon"] != input.Icon {
						t.Fatal("lost icon")
					}
				} else if _, ok := body["icon"]; ok {
					t.Fatal("omitted icon serialized")
				}
				if body["target"].(map[string]any)["channels"].([]any)[0] != "ops" {
					t.Fatal("lost target")
				}
			}
			if _, err := client.Notifications.SendPushNotification(req); err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestInvalidPushLevel(t *testing.T) {
	for _, level := range []PushInterruptionLevel{"critical", "timeSensitive", "default"} {
		if _, err := normalizePushNotificationRequest(PushNotificationInput{Title: "Test", InterruptionLevel: level}); err == nil {
			t.Fatal("accepted invalid level")
		}
	}
}
