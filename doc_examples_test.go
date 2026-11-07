package activitysmith_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	activitysmithsdk "github.com/ActivitySmithHQ/activitysmith-go"
	"github.com/ActivitySmithHQ/activitysmith-go/generated"
)

// TestDocExamples runs the Go SDK doc snippets as written against a mock API,
// so a doc example that stops compiling or decoding fails CI.
func TestDocExamples(t *testing.T) {
	var calls []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.ReadAll(r.Body)
		call := r.Method + " " + r.URL.Path
		if r.URL.RawQuery != "" {
			call += "?" + r.URL.RawQuery
		}
		calls = append(calls, call)

		w.Header().Set("Content-Type", "application/json")
		switch r.Method + " " + r.URL.Path {
		case "POST /push-notification":
			_, _ = w.Write([]byte(`{"success":true,"notification_id":"pn-1","devices_targeted":1,"devices_notified":1,"users_notified":1,"timestamp":"2026-02-07T00:00:00Z"}`))
		case "POST /push-notification/pn-1/dismiss":
			_, _ = w.Write([]byte(`{"success":true,"notification_id":"pn-1","dismissed_at":"2026-02-07T00:05:00Z"}`))
		case "GET /account/devices":
			_, _ = w.Write([]byte(`{"devices":[{"user_email":"owner@example.com","device_id":"dev-1","device_type":"ios","is_active":true,"token_status":{"push_notifications":true},"app_version":"1.17.0","device_name":"iPhone","system_name":"iOS","system_version":"27.0","created_at":"2026-01-01T00:00:00Z","updated_at":"2026-02-07T00:00:00Z"}],"meta":{"total":1,"limit":100,"offset":0,"has_more":false,"next_offset":null,"include_inactive":false}}`))
		case "PUT /metrics/mrr":
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(`{"metric":{"public_id":"m-1","key":"mrr","label":"MRR","format":"currency","currency_code":"USD","unit":null,"unit_spacing":"none","latest_value":null,"latest_value_at":null,"created_at":"2026-10-26T09:00:00Z","updated_at":"2026-10-26T09:00:00Z"}}`))
		case "POST /metrics/mrr/value":
			_, _ = w.Write([]byte(`{"success":true}`))
		default:
			t.Errorf("unexpected request: %s", call)
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	activitysmith, err := activitysmithsdk.New("test-api-key")
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}
	cfg := activitysmith.APIClient().GetConfig()
	cfg.Servers = generated.ServerConfigurations{{URL: server.URL}}
	for key := range cfg.OperationServers {
		cfg.OperationServers[key] = cfg.Servers
	}

	// Dismiss a Push Notification.
	sent, err := activitysmith.Notifications.Send(activitysmithsdk.PushNotificationInput{
		Title:   "🙋 Agent needs your input",
		Message: "Approve the database migration?",
	})
	if err != nil {
		t.Fatalf("Send returned error: %v", err)
	}

	_, err = activitysmith.Notifications.Dismiss(sent.GetNotificationId())
	if err != nil {
		t.Fatalf("Dismiss returned error: %v", err)
	}

	// List paired devices.
	result, err := activitysmith.Devices.List(activitysmithsdk.DeviceListInput{})
	if err != nil {
		t.Fatalf("Devices.List returned error: %v", err)
	}

	for _, device := range result.GetDevices() {
		fmt.Println(device.GetDeviceName(), device.GetSystemName(), device.GetIsActive())
	}

	result, err = activitysmith.Devices.List(activitysmithsdk.DeviceListInput{
		IncludeInactive: true,
	})
	if err != nil {
		t.Fatalf("Devices.List with IncludeInactive returned error: %v", err)
	}
	if len(result.GetDevices()) != 1 {
		t.Fatalf("expected 1 device, got %d", len(result.GetDevices()))
	}

	// Create a metric.
	activitysmith.Metrics.Create("mrr", activitysmithsdk.MetricInput{
		Label:        "MRR",
		Format:       "currency",
		CurrencyCode: "USD",
	})

	activitysmith.Metrics.Update("mrr", 12450)

	want := []string{
		"POST /push-notification",
		"POST /push-notification/pn-1/dismiss",
		"GET /account/devices",
		"GET /account/devices?include_inactive=true",
		"PUT /metrics/mrr",
		"POST /metrics/mrr/value",
	}
	if len(calls) != len(want) {
		t.Fatalf("unexpected calls: %v", calls)
	}
	for i, call := range want {
		if calls[i] != call {
			t.Fatalf("call %d: want %q, got %q", i, call, calls[i])
		}
	}
}
