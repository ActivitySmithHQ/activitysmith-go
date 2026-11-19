package activitysmith

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTimerDateSerialization(t *testing.T) {
	startAt := time.Date(2026, 11, 19, 9, 30, 0, 0, time.UTC)
	endAt := time.Date(2026, 11, 19, 10, 0, 0, 0, time.FixedZone("ICT", 7*60*60))

	server, requests := newAPITestServer(t)
	defer server.Close()
	client, err := New("test")
	if err != nil {
		t.Fatal(err)
	}
	overrideHostForTests(client, server.URL)

	timerState := LiveActivityContentStateInput{Title: "Benchmark Run", Type: LiveActivityTypeTimer}.
		WithTimerStartAt(startAt).
		WithTimerEndAt(endAt)
	plainState := LiveActivityContentStateInput{Title: "Benchmark Run", Type: LiveActivityTypeTimer}

	calls := []struct {
		name  string
		dates bool
		send  func() error
	}{
		{"start content state", true, func() error {
			_, err := client.LiveActivities.Start(LiveActivityStartInput{ContentState: timerState})
			return err
		}},
		{"start top level", true, func() error {
			_, err := client.LiveActivities.Start(LiveActivityStartInput{
				Title:        "Benchmark Run",
				Type:         LiveActivityTypeTimer,
				TimerStartAt: startAt,
				TimerEndAt:   endAt,
			})
			return err
		}},
		{"update content state", true, func() error {
			_, err := client.LiveActivities.Update(LiveActivityUpdateInput{ActivityID: "activity-1", ContentState: timerState})
			return err
		}},
		{"stream content state", true, func() error {
			_, err := client.LiveActivities.Stream("prod-web-1", LiveActivityStreamInput{ContentState: timerState})
			return err
		}},
		{"stream top level", true, func() error {
			_, err := client.LiveActivities.Stream("prod-web-1", LiveActivityStreamInput{
				Title:        "Benchmark Run",
				Type:         LiveActivityTypeTimer,
				TimerStartAt: startAt,
				TimerEndAt:   endAt,
			})
			return err
		}},
		{"end stream", true, func() error {
			_, err := client.LiveActivities.EndStream("prod-web-1", LiveActivityStreamEndInput{
				Title:        "Benchmark Run",
				TimerStartAt: startAt,
				TimerEndAt:   endAt,
			})
			return err
		}},
		{"start unset", false, func() error {
			_, err := client.LiveActivities.Start(LiveActivityStartInput{ContentState: plainState.WithDurationSeconds(300)})
			return err
		}},
		{"stream unset", false, func() error {
			_, err := client.LiveActivities.Stream("prod-web-1", LiveActivityStreamInput{ContentState: plainState})
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
		name := calls[index].name
		var body struct {
			ContentState map[string]any `json:"content_state"`
		}
		if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
			t.Fatal(err)
		}
		if body.ContentState == nil {
			t.Fatalf("%s: missing content_state: %s", name, request.Body)
		}
		startValue, hasStart := body.ContentState["timer_start_at"]
		endValue, hasEnd := body.ContentState["timer_end_at"]
		if !calls[index].dates {
			if hasStart || hasEnd {
				t.Fatalf("%s: unset timer dates serialized: %s", name, request.Body)
			}
			continue
		}
		if startValue != "2026-11-19T09:30:00Z" {
			t.Fatalf("%s: timer_start_at = %v: %s", name, startValue, request.Body)
		}
		if endValue != "2026-11-19T10:00:00+07:00" {
			t.Fatalf("%s: timer_end_at = %v: %s", name, endValue, request.Body)
		}
		for _, value := range []any{startValue, endValue} {
			if _, err := time.Parse(time.RFC3339, value.(string)); err != nil {
				t.Fatalf("%s: %v is not RFC 3339: %v", name, value, err)
			}
		}
	}
}
