package activitysmith

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ActivitySmithHQ/activitysmith-go/generated"
)

type capturedRequest struct {
	Path          string
	Authorization string
	Body          string
}

func newAPITestServer(t *testing.T) (*httptest.Server, *[]capturedRequest) {
	t.Helper()

	requests := []capturedRequest{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}
		defer r.Body.Close()

		requests = append(requests, capturedRequest{
			Path:          r.URL.Path,
			Authorization: r.Header.Get("Authorization"),
			Body:          string(body),
		})

		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/push-notification":
			_, _ = w.Write([]byte(`{"success":true,"devices_notified":1,"timestamp":"2026-02-07T00:00:00Z"}`))
		case "/live-activity/start":
			_, _ = w.Write([]byte(`{"success":true,"activity_id":"act-1","devices_notified":1,"timestamp":"2026-02-07T00:00:00Z"}`))
		case "/live-activity/update":
			_, _ = w.Write([]byte(`{"success":true,"activity_id":"act-1","devices_notified":1,"timestamp":"2026-02-07T00:00:00Z"}`))
		case "/live-activity/end":
			_, _ = w.Write([]byte(`{"success":true,"activity_id":"act-1","devices_notified":1,"timestamp":"2026-02-07T00:00:00Z"}`))
		default:
			http.NotFound(w, r)
		}
	})

	server := httptest.NewServer(handler)
	return server, &requests
}

func TestNotificationsShortAndLegacyMethods(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()

	client, err := New("test-api-key", &Options{BaseURL: server.URL})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	payload := generated.PushNotificationRequest{Title: "Build Failed"}
	if _, err := client.Notifications.Send(payload); err != nil {
		t.Fatalf("Send returned error: %v", err)
	}
	if _, err := client.Notifications.SendPushNotification(payload); err != nil {
		t.Fatalf("SendPushNotification returned error: %v", err)
	}

	if len(*requests) != 2 {
		t.Fatalf("expected 2 requests, got %d", len(*requests))
	}

	for i, req := range *requests {
		if req.Path != "/push-notification" {
			t.Fatalf("request %d path mismatch: %s", i, req.Path)
		}
		if req.Authorization != "Bearer test-api-key" {
			t.Fatalf("request %d auth mismatch: %s", i, req.Authorization)
		}
		if !strings.Contains(req.Body, `"title":"Build Failed"`) {
			t.Fatalf("request %d body missing title: %s", i, req.Body)
		}
	}
}

func TestLiveActivitiesShortAndLegacyMethods(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()

	client, err := New("test-api-key", &Options{BaseURL: server.URL})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	startPayload := generated.LiveActivityStartRequest{
		ContentState: generated.ContentStateStart{
			Title:         "Deploy",
			NumberOfSteps: 4,
			CurrentStep:   1,
			Type:          "segmented_progress",
		},
	}
	updatePayload := generated.LiveActivityUpdateRequest{
		ActivityId: "act-1",
		ContentState: generated.ContentStateUpdate{
			Title:       "Deploy",
			CurrentStep: 2,
		},
	}
	endPayload := generated.LiveActivityEndRequest{
		ActivityId: "act-1",
		ContentState: generated.ContentStateEnd{
			Title:       "Deploy",
			CurrentStep: 4,
		},
	}

	if _, err := client.LiveActivities.Start(startPayload); err != nil {
		t.Fatalf("Start returned error: %v", err)
	}
	if _, err := client.LiveActivities.Update(updatePayload); err != nil {
		t.Fatalf("Update returned error: %v", err)
	}
	if _, err := client.LiveActivities.End(endPayload); err != nil {
		t.Fatalf("End returned error: %v", err)
	}
	if _, err := client.LiveActivities.StartLiveActivity(startPayload); err != nil {
		t.Fatalf("StartLiveActivity returned error: %v", err)
	}
	if _, err := client.LiveActivities.UpdateLiveActivity(updatePayload); err != nil {
		t.Fatalf("UpdateLiveActivity returned error: %v", err)
	}
	if _, err := client.LiveActivities.EndLiveActivity(endPayload); err != nil {
		t.Fatalf("EndLiveActivity returned error: %v", err)
	}

	if len(*requests) != 6 {
		t.Fatalf("expected 6 requests, got %d", len(*requests))
	}

	gotPaths := []string{}
	for _, req := range *requests {
		gotPaths = append(gotPaths, req.Path)
		if req.Authorization != "Bearer test-api-key" {
			t.Fatalf("auth mismatch: %s", req.Authorization)
		}
	}

	wantPaths := []string{
		"/live-activity/start",
		"/live-activity/update",
		"/live-activity/end",
		"/live-activity/start",
		"/live-activity/update",
		"/live-activity/end",
	}

	gotJSON, _ := json.Marshal(gotPaths)
	wantJSON, _ := json.Marshal(wantPaths)
	if string(gotJSON) != string(wantJSON) {
		t.Fatalf("path order mismatch: got=%s want=%s", gotJSON, wantJSON)
	}
}
