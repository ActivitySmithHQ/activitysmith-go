package activitysmith

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ActivitySmithHQ/activitysmith-go/generated"
)

type capturedRequest struct {
	Path          string
	Authorization string
	SDK           string
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
			SDK:           r.Header.Get("X-ActivitySmith-SDK"),
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
		case "/live-activity/stream/prod-web-1":
			if r.Method == http.MethodPut {
				_, _ = w.Write([]byte(`{"success":true,"operation":"started","stream_key":"prod-web-1","activity_id":"act-1","timestamp":"2026-02-07T00:00:00Z"}`))
			} else if r.Method == http.MethodDelete {
				_, _ = w.Write([]byte(`{"success":true,"operation":"ended","stream_key":"prod-web-1","activity_id":"act-1","timestamp":"2026-02-07T00:00:00Z"}`))
			} else {
				http.NotFound(w, r)
			}
		case "/metrics/deploy.success_rate/value":
			_, _ = w.Write([]byte(`{"success":true}`))
		case "/metrics/prod.status/value":
			_, _ = w.Write([]byte(`{"success":true}`))
		default:
			http.NotFound(w, r)
		}
	})

	server := httptest.NewServer(handler)
	return server, &requests
}

func activityMetricNumber(value float32) generated.ActivityMetricValue {
	return generated.Float32AsActivityMetricValue(&value)
}

func TestNotificationsShortAndLegacyMethods(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()

	client, err := New("test-api-key")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	overrideHostForTests(client, server.URL)

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
		if req.SDK != "go-v"+Version {
			t.Fatalf("request %d sdk header mismatch: %s", i, req.SDK)
		}
		if !strings.Contains(req.Body, `"title":"Build Failed"`) {
			t.Fatalf("request %d body missing title: %s", i, req.Body)
		}
	}
}

func TestLiveActivitiesShortAndLegacyMethods(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()

	client, err := New("test-api-key")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	overrideHostForTests(client, server.URL)

	startState := generated.NewContentStateStart("Deploy", "segmented_progress")
	startState.SetNumberOfSteps(4)
	startState.SetCurrentStep(1)
	startPayload := generated.LiveActivityStartRequest{
		ContentState: *startState,
	}
	updateState := generated.NewContentStateUpdate("Deploy")
	updateState.SetCurrentStep(2)
	updatePayload := generated.LiveActivityUpdateRequest{
		ActivityId:   "act-1",
		ContentState: *updateState,
	}
	endState := generated.NewContentStateEnd("Deploy")
	endState.SetCurrentStep(4)
	endPayload := generated.LiveActivityEndRequest{
		ActivityId:   "act-1",
		ContentState: *endState,
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
		if req.SDK != "go-v"+Version {
			t.Fatalf("sdk header mismatch: %s", req.SDK)
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

func TestLiveActivitiesStreamShortAndLegacyMethods(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()

	client, err := New("test-api-key")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	overrideHostForTests(client, server.URL)

	streamInput := LiveActivityStreamInput{
		Title:    "Server Health",
		Subtitle: "prod-web-1",
		Type:     "metrics",
		Metrics: []generated.ActivityMetric{
			{Label: "CPU", Value: activityMetricNumber(9), Unit: generated.PtrString("%")},
			{Label: "MEM", Value: activityMetricNumber(45), Unit: generated.PtrString("%")},
		},
		Channels: []string{"ops"},
	}
	endInput := LiveActivityStreamEndInput{
		Title:    "Server Health",
		Subtitle: "prod-web-1",
		Type:     "metrics",
		Metrics: []generated.ActivityMetric{
			{Label: "CPU", Value: activityMetricNumber(7), Unit: generated.PtrString("%")},
			{Label: "MEM", Value: activityMetricNumber(38), Unit: generated.PtrString("%")},
		},
	}

	if _, err := client.LiveActivities.Stream("prod-web-1", streamInput); err != nil {
		t.Fatalf("Stream returned error: %v", err)
	}
	if _, err := client.LiveActivities.ReconcileLiveActivityStream("prod-web-1", streamInput.toGenerated()); err != nil {
		t.Fatalf("ReconcileLiveActivityStream returned error: %v", err)
	}
	if _, err := client.LiveActivities.EndStream("prod-web-1", endInput); err != nil {
		t.Fatalf("EndStream returned error: %v", err)
	}
	if _, err := client.LiveActivities.EndLiveActivityStream("prod-web-1", endInput.toGenerated()); err != nil {
		t.Fatalf("EndLiveActivityStream returned error: %v", err)
	}

	if len(*requests) != 4 {
		t.Fatalf("expected 4 requests, got %d", len(*requests))
	}

	wantPaths := []string{
		"/live-activity/stream/prod-web-1",
		"/live-activity/stream/prod-web-1",
		"/live-activity/stream/prod-web-1",
		"/live-activity/stream/prod-web-1",
	}

	gotPaths := make([]string, 0, len(*requests))
	for _, req := range *requests {
		gotPaths = append(gotPaths, req.Path)
		if req.Authorization != "Bearer test-api-key" {
			t.Fatalf("auth mismatch: %s", req.Authorization)
		}
		if req.SDK != "go-v"+Version {
			t.Fatalf("sdk header mismatch: %s", req.SDK)
		}
	}

	gotJSON, _ := json.Marshal(gotPaths)
	wantJSON, _ := json.Marshal(wantPaths)
	if string(gotJSON) != string(wantJSON) {
		t.Fatalf("path order mismatch: got=%s want=%s", gotJSON, wantJSON)
	}
	if !strings.Contains((*requests)[0].Body, `"target":{"channels":["ops"]}`) {
		t.Fatalf("stream body missing target channels: %s", (*requests)[0].Body)
	}
}

func TestMetricsShortAndLegacyMethods(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()

	client, err := New("test-api-key")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	overrideHostForTests(client, server.URL)

	timestamp, err := time.Parse(time.RFC3339, "2026-05-03T12:30:00Z")
	if err != nil {
		t.Fatalf("failed to parse timestamp: %v", err)
	}

	if _, err := client.Metrics.Update("deploy.success_rate", MetricValueInput{
		Value:     99.9,
		Timestamp: &timestamp,
	}); err != nil {
		t.Fatalf("Metrics.Update returned error: %v", err)
	}

	if _, err := client.Metrics.Update("prod.status", "healthy"); err != nil {
		t.Fatalf("Metrics.Update string returned error: %v", err)
	}

	value := float32(42)
	request := *generated.NewMetricValueUpdateRequest(
		generated.Float32AsMetricValueUpdateRequestValue(&value),
	)
	if _, err := client.Metrics.UpdateMetricValue("deploy.success_rate", request); err != nil {
		t.Fatalf("UpdateMetricValue returned error: %v", err)
	}

	if len(*requests) != 3 {
		t.Fatalf("expected 3 requests, got %d", len(*requests))
	}

	wantPaths := []string{
		"/metrics/deploy.success_rate/value",
		"/metrics/prod.status/value",
		"/metrics/deploy.success_rate/value",
	}
	for i, req := range *requests {
		if req.Path != wantPaths[i] {
			t.Fatalf("request %d path mismatch: got=%s want=%s", i, req.Path, wantPaths[i])
		}
		if req.Authorization != "Bearer test-api-key" {
			t.Fatalf("request %d auth mismatch: %s", i, req.Authorization)
		}
		if req.SDK != "go-v"+Version {
			t.Fatalf("request %d sdk header mismatch: %s", i, req.SDK)
		}
	}

	if !strings.Contains((*requests)[0].Body, `"value":99.9`) ||
		!strings.Contains((*requests)[0].Body, `"timestamp":"2026-05-03T12:30:00Z"`) {
		t.Fatalf("metric numeric body mismatch: %s", (*requests)[0].Body)
	}
	if !strings.Contains((*requests)[1].Body, `"value":"healthy"`) {
		t.Fatalf("metric string body mismatch: %s", (*requests)[1].Body)
	}
}

func TestDXInputsIncludeOptionalFields(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()

	client, err := New("test-api-key")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	overrideHostForTests(client, server.URL)

	pushInput := PushNotificationInput{
		Title:       "Deploy",
		Message:     "Build complete",
		Subtitle:    "production",
		Media:       "https://cdn.example.com/output/homepage.png",
		Redirection: "https://github.com/acme/web/pull/482",
		Channels:    []string{"devs", "ops"},
	}
	if _, err := client.Notifications.Send(pushInput); err != nil {
		t.Fatalf("Send returned error: %v", err)
	}

	startInput := LiveActivityStartInput{
		Title:         "Deploy",
		Subtitle:      "start",
		NumberOfSteps: 4,
		CurrentStep:   1,
		Type:          "segmented_progress",
		Color:         "yellow",
		Action: &LiveActivityActionInput{
			Title: "Open Workflow",
			Type:  "open_url",
			URL:   "https://github.com/acme/payments-api/actions/runs/1234567890",
		},
		Channels: []string{"devs", "ops"},
	}
	if _, err := client.LiveActivities.Start(startInput); err != nil {
		t.Fatalf("Start returned error: %v", err)
	}

	updateInput := LiveActivityUpdateInput{
		ActivityID:    "act-1",
		Title:         "Deploy",
		Subtitle:      "testing",
		CurrentStep:   2,
		NumberOfSteps: 4,
		Action: &LiveActivityActionInput{
			Title:  "Pause Reindex",
			Type:   "webhook",
			URL:    "https://ops.example.com/hooks/search/reindex/pause",
			Method: "POST",
			Body: map[string]interface{}{
				"job_id": "reindex-2026-03-19",
			},
		},
	}
	if _, err := client.LiveActivities.Update(updateInput); err != nil {
		t.Fatalf("Update returned error: %v", err)
	}

	endInput := LiveActivityEndInput{
		ActivityID:         "act-1",
		Title:              "Deploy",
		Subtitle:           "done",
		CurrentStep:        4,
		AutoDismissMinutes: 3,
		Action: &LiveActivityActionInput{
			Title: "Open Workflow",
			Type:  "open_url",
			URL:   "https://github.com/acme/payments-api/actions/runs/1234567890",
		},
	}
	if _, err := client.LiveActivities.End(endInput); err != nil {
		t.Fatalf("End returned error: %v", err)
	}

	if len(*requests) != 4 {
		t.Fatalf("expected 4 requests, got %d", len(*requests))
	}

	bodies := make([]string, 0, len(*requests))
	for _, req := range *requests {
		bodies = append(bodies, req.Body)
	}

	if !strings.Contains(bodies[0], `"message":"Build complete"`) {
		t.Fatalf("push body missing message: %s", bodies[0])
	}
	if !strings.Contains(bodies[0], `"subtitle":"production"`) {
		t.Fatalf("push body missing subtitle: %s", bodies[0])
	}
	if !strings.Contains(bodies[0], `"media":"https://cdn.example.com/output/homepage.png"`) {
		t.Fatalf("push body missing media: %s", bodies[0])
	}
	if !strings.Contains(bodies[0], `"redirection":"https://github.com/acme/web/pull/482"`) {
		t.Fatalf("push body missing redirection: %s", bodies[0])
	}
	if !strings.Contains(bodies[0], `"channels":["devs","ops"]`) {
		t.Fatalf("push body missing target channels: %s", bodies[0])
	}

	if !strings.Contains(bodies[1], `"subtitle":"start"`) {
		t.Fatalf("start body missing subtitle: %s", bodies[1])
	}
	if !strings.Contains(bodies[1], `"color":"yellow"`) {
		t.Fatalf("start body missing color: %s", bodies[1])
	}
	if !strings.Contains(bodies[1], `"action":{"title":"Open Workflow","type":"open_url","url":"https://github.com/acme/payments-api/actions/runs/1234567890"}`) {
		t.Fatalf("start body missing action: %s", bodies[1])
	}
	if !strings.Contains(bodies[1], `"channels":["devs","ops"]`) {
		t.Fatalf("start body missing target channels: %s", bodies[1])
	}

	if !strings.Contains(bodies[2], `"subtitle":"testing"`) {
		t.Fatalf("update body missing subtitle: %s", bodies[2])
	}
	if !strings.Contains(bodies[2], `"number_of_steps":4`) {
		t.Fatalf("update body missing number_of_steps: %s", bodies[2])
	}
	if !strings.Contains(bodies[2], `"title":"Pause Reindex"`) ||
		!strings.Contains(bodies[2], `"type":"webhook"`) ||
		!strings.Contains(bodies[2], `"url":"https://ops.example.com/hooks/search/reindex/pause"`) ||
		!strings.Contains(bodies[2], `"method":"POST"`) ||
		!strings.Contains(bodies[2], `"job_id":"reindex-2026-03-19"`) {
		t.Fatalf("update body missing action: %s", bodies[2])
	}

	if !strings.Contains(bodies[3], `"subtitle":"done"`) {
		t.Fatalf("end body missing subtitle: %s", bodies[3])
	}
	if !strings.Contains(bodies[3], `"auto_dismiss_minutes":3`) {
		t.Fatalf("end body missing auto_dismiss_minutes: %s", bodies[3])
	}
	if !strings.Contains(bodies[3], `"action":{"title":"Open Workflow","type":"open_url","url":"https://github.com/acme/payments-api/actions/runs/1234567890"}`) {
		t.Fatalf("end body missing action: %s", bodies[3])
	}
}

func TestEndInputCanExplicitlySendZeroAutoDismissMinutes(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()

	client, err := New("test-api-key")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	overrideHostForTests(client, server.URL)

	endInput := LiveActivityEndInput{
		ActivityID:  "act-1",
		Title:       "Deploy",
		CurrentStep: 4,
	}.WithAutoDismissMinutes(0)

	if _, err := client.LiveActivities.End(endInput); err != nil {
		t.Fatalf("End returned error: %v", err)
	}

	if len(*requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(*requests))
	}

	body := (*requests)[0].Body
	if !strings.Contains(body, `"auto_dismiss_minutes":0`) {
		t.Fatalf("end body missing explicit zero auto_dismiss_minutes: %s", body)
	}
}

func TestProgressInputsSerializeProgressFieldsWithoutSegmentedFields(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()

	client, err := New("test-api-key")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	overrideHostForTests(client, server.URL)

	startInput := LiveActivityStartInput{
		Title:    "Render export",
		Subtitle: "encoding frames",
		Type:     "progress",
		Color:    "purple",
	}.WithPercentage(67)

	updateInput := LiveActivityUpdateInput{
		ActivityID: "act-1",
		Title:      "Render export",
		Type:       "progress",
	}.WithValue(241).WithUpperLimit(360)

	endInput := LiveActivityEndInput{
		ActivityID: "act-1",
		Title:      "Render export",
		Type:       "progress",
	}.WithPercentage(100).WithAutoDismissMinutes(0)

	if _, err := client.LiveActivities.Start(startInput); err != nil {
		t.Fatalf("Start returned error: %v", err)
	}
	if _, err := client.LiveActivities.Update(updateInput); err != nil {
		t.Fatalf("Update returned error: %v", err)
	}
	if _, err := client.LiveActivities.End(endInput); err != nil {
		t.Fatalf("End returned error: %v", err)
	}

	if len(*requests) != 3 {
		t.Fatalf("expected 3 requests, got %d", len(*requests))
	}

	startBody := (*requests)[0].Body
	if !strings.Contains(startBody, `"type":"progress"`) {
		t.Fatalf("start body missing progress type: %s", startBody)
	}
	if !strings.Contains(startBody, `"percentage":67`) {
		t.Fatalf("start body missing percentage: %s", startBody)
	}
	if strings.Contains(startBody, `"current_step"`) || strings.Contains(startBody, `"number_of_steps"`) {
		t.Fatalf("start body should not include segmented fields: %s", startBody)
	}

	updateBody := (*requests)[1].Body
	if !strings.Contains(updateBody, `"value":241`) || !strings.Contains(updateBody, `"upper_limit":360`) {
		t.Fatalf("update body missing progress value fields: %s", updateBody)
	}
	if strings.Contains(updateBody, `"current_step"`) || strings.Contains(updateBody, `"number_of_steps"`) {
		t.Fatalf("update body should not include segmented fields: %s", updateBody)
	}

	endBody := (*requests)[2].Body
	if !strings.Contains(endBody, `"percentage":100`) {
		t.Fatalf("end body missing percentage: %s", endBody)
	}
	if !strings.Contains(endBody, `"auto_dismiss_minutes":0`) {
		t.Fatalf("end body missing explicit zero auto_dismiss_minutes: %s", endBody)
	}
	if strings.Contains(endBody, `"current_step"`) || strings.Contains(endBody, `"number_of_steps"`) {
		t.Fatalf("end body should not include segmented fields: %s", endBody)
	}
}

func TestStatsInputsSerializeMetricValuesAndColors(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()

	client, err := New("test-api-key")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	overrideHostForTests(client, server.URL)

	revenue, err := NewActivityMetric("Revenue", "$2430", WithActivityMetricColor("blue"))
	if err != nil {
		t.Fatalf("NewActivityMetric returned error: %v", err)
	}
	orders, err := NewActivityMetric("Orders", 37, WithActivityMetricColor("green"))
	if err != nil {
		t.Fatalf("NewActivityMetric returned error: %v", err)
	}

	startInput := LiveActivityStartInput{
		Title:    "Sales",
		Subtitle: "last hour",
		Type:     LiveActivityTypeStats,
		Metrics:  []generated.ActivityMetric{revenue, orders},
	}
	updateInput := LiveActivityUpdateInput{
		ActivityID: "act-1",
		Title:      "Sales",
		Type:       LiveActivityTypeStats,
		Metrics:    []generated.ActivityMetric{revenue},
	}
	endInput := LiveActivityEndInput{
		ActivityID: "act-1",
		Title:      "Sales",
		Type:       LiveActivityTypeStats,
		Metrics:    []generated.ActivityMetric{orders},
	}

	if _, err := client.LiveActivities.Start(startInput); err != nil {
		t.Fatalf("Start returned error: %v", err)
	}
	if _, err := client.LiveActivities.Update(updateInput); err != nil {
		t.Fatalf("Update returned error: %v", err)
	}
	if _, err := client.LiveActivities.End(endInput); err != nil {
		t.Fatalf("End returned error: %v", err)
	}

	if len(*requests) != 3 {
		t.Fatalf("expected 3 requests, got %d", len(*requests))
	}

	startBody := (*requests)[0].Body
	if !strings.Contains(startBody, `"type":"stats"`) {
		t.Fatalf("start body missing stats type: %s", startBody)
	}
	if !strings.Contains(startBody, `"value":"$2430"`) {
		t.Fatalf("start body missing string metric value: %s", startBody)
	}
	if !strings.Contains(startBody, `"color":"blue"`) || !strings.Contains(startBody, `"color":"green"`) {
		t.Fatalf("start body missing metric colors: %s", startBody)
	}

	if !strings.Contains((*requests)[1].Body, `"metrics"`) {
		t.Fatalf("update body missing metrics: %s", (*requests)[1].Body)
	}
	if !strings.Contains((*requests)[2].Body, `"metrics"`) {
		t.Fatalf("end body missing metrics: %s", (*requests)[2].Body)
	}
}

func TestAlertInputsSerializeMessageIconAndBadge(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()

	client, err := New("test-api-key")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	overrideHostForTests(client, server.URL)

	icon := AlertIcon("sparkles", "yellow")
	badge := AlertBadge("Customer", "magenta")
	streamInput := LiveActivityStreamInput{
		Title:   "Reactivation",
		Type:    LiveActivityTypeAlert,
		Message: "Lumen came back after 2 weeks",
		Icon:    &icon,
		Badge:   &badge,
		Color:   "red",
	}

	if _, err := client.LiveActivities.Stream("prod-web-1", streamInput); err != nil {
		t.Fatalf("Stream returned error: %v", err)
	}

	if len(*requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(*requests))
	}

	body := (*requests)[0].Body
	if !strings.Contains(body, `"type":"alert"`) {
		t.Fatalf("alert body missing type: %s", body)
	}
	if !strings.Contains(body, `"message":"Lumen came back after 2 weeks"`) {
		t.Fatalf("alert body missing message: %s", body)
	}
	if !strings.Contains(body, `"icon":{"color":"yellow","symbol":"sparkles"}`) {
		t.Fatalf("alert body missing icon: %s", body)
	}
	if !strings.Contains(body, `"badge":{"color":"magenta","title":"Customer"}`) {
		t.Fatalf("alert body missing badge: %s", body)
	}
	if strings.Contains(body, `"color":"red"`) || strings.Contains(body, `"color":"blue"`) {
		t.Fatalf("alert body should not include root color: %s", body)
	}
}

func TestServicesRejectUnsupportedInputTypes(t *testing.T) {
	server, _ := newAPITestServer(t)
	defer server.Close()

	client, err := New("test-api-key")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	overrideHostForTests(client, server.URL)

	if _, err := client.Notifications.Send(123); err == nil {
		t.Fatal("expected error for unsupported notifications input type")
	}
	if _, err := client.LiveActivities.Start(123); err == nil {
		t.Fatal("expected error for unsupported start input type")
	}
	if _, err := client.LiveActivities.Update(123); err == nil {
		t.Fatal("expected error for unsupported update input type")
	}
	if _, err := client.LiveActivities.End(123); err == nil {
		t.Fatal("expected error for unsupported end input type")
	}
}

func TestNotificationsRejectMediaAndActionsCombination(t *testing.T) {
	server, requests := newAPITestServer(t)
	defer server.Close()

	client, err := New("test-api-key")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	overrideHostForTests(client, server.URL)

	action := PushAction("Open", "open_url", "https://example.com")
	if _, err := client.Notifications.Send(PushNotificationInput{
		Title:   "Build Failed",
		Media:   "https://cdn.example.com/output/homepage.png",
		Actions: []PushNotificationAction{action},
	}); err == nil || err.Error() != ErrPushNotificationMediaActionsConflict.Error() {
		t.Fatalf("expected ErrPushNotificationMediaActionsConflict, got %v", err)
	}

	request := generated.PushNotificationRequest{Title: "Build Failed"}
	request.SetMedia("https://cdn.example.com/output/homepage.png")
	request.SetActions([]generated.PushNotificationAction{action})

	if _, err := client.Notifications.Send(request); err == nil || err.Error() != ErrPushNotificationMediaActionsConflict.Error() {
		t.Fatalf("expected ErrPushNotificationMediaActionsConflict for generated request, got %v", err)
	}

	if len(*requests) != 0 {
		t.Fatalf("expected no API calls, got %d", len(*requests))
	}
}

func overrideHostForTests(client *Client, host string) {
	serverConfig := generated.ServerConfigurations{{URL: host}}
	cfg := client.APIClient().GetConfig()
	cfg.Servers = serverConfig
	for key := range cfg.OperationServers {
		cfg.OperationServers[key] = serverConfig
	}
}
