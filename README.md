# ActivitySmith Go SDK

The ActivitySmith Go SDK provides convenient access to the ActivitySmith API from Go applications.

## Documentation

See [API reference](https://activitysmith.com/docs/api-reference/introduction).

## Installation

```sh
go get github.com/ActivitySmithHQ/activitysmith-go
```

## Setup

```go
package main

import (
	"log"

	activitysmithsdk "github.com/ActivitySmithHQ/activitysmith-go"
)

func main() {
	activitysmith, err := activitysmithsdk.New("YOUR_API_KEY")
	if err != nil {
		log.Fatal(err)
	}

	_ = activitysmith
}
```

## Push Notifications

### Send a Push Notification

<p align="center">
  <img src="https://cdn.activitysmith.com/features/new-subscription-push-notification.png" alt="Push notification example" width="680" />
</p>

Use `activitysmith.Notifications.Send` with either `activitysmithsdk.PushNotificationInput` for common notification fields or `generated.PushNotificationRequest` if you want full control over the generated model.

```go
input := activitysmithsdk.PushNotificationInput{
	Title:   "New subscription 💸",
	Message: "Customer upgraded to Pro plan",
}

_, err := activitysmith.Notifications.Send(input)
if err != nil {
	log.Fatal(err)
}
```

### Rich Push Notifications with Media

<p align="center">
  <img src="https://cdn.activitysmith.com/features/rich-push-notification-with-image.png" alt="Rich push notification with image" width="680" />
</p>

```go
input := activitysmithsdk.PushNotificationInput{
	Title:       "Homepage ready",
	Message:     "Your agent finished the redesign.",
	Media:       "https://cdn.example.com/output/homepage-v2.png",
	Redirection: "https://github.com/acme/web/pull/482",
}

_, err := activitysmith.Notifications.Send(input)
if err != nil {
	log.Fatal(err)
}
```

Send images, videos, or audio with your push notifications, press and hold to preview media directly from the notification, then tap through to open the linked content.

<p align="center">
  <img src="https://cdn.activitysmith.com/features/rich-push-notification-with-audio.png" alt="Rich push notification with audio" width="680" />
</p>

What will work:

- direct image URL: `.jpg`, `.png`, `.gif`, etc.
- direct audio file URL: `.mp3`, `.m4a`, etc.
- direct video file URL: `.mp4`, `.mov`, etc.
- URL that responds with a proper media `Content-Type`, even if the path has no extension

### Actionable Push Notifications

<p align="center">
  <img src="https://cdn.activitysmith.com/features/actionable-push-notifications-2.png" alt="Actionable push notification example" width="680" />
</p>

Actionable push notifications can open a URL on tap or trigger actions when someone long-presses the notification.
Webhooks are executed by the ActivitySmith backend.

```go
request := generated.NewPushNotificationRequest("New subscription 💸")
request.SetMessage("Customer upgraded to Pro plan")
request.SetRedirection("https://crm.example.com/customers/cus_9f3a1d") // Optional

crmAction := generated.NewPushNotificationAction(
	"Open CRM Profile",
	generated.PUSHNOTIFICATIONACTIONTYPE_OPEN_URL,
	"https://crm.example.com/customers/cus_9f3a1d",
)

onboardingAction := generated.NewPushNotificationAction(
	"Start Onboarding Workflow",
	generated.PUSHNOTIFICATIONACTIONTYPE_WEBHOOK,
	"https://hooks.example.com/activitysmith/onboarding/start",
)
onboardingAction.SetMethod(generated.PUSHNOTIFICATIONACTIONMETHOD_POST)
onboardingAction.SetBody(map[string]interface{}{
	"customer_id": "cus_9f3a1d",
	"plan": "pro",
})

request.SetActions([]generated.PushNotificationAction{
	*crmAction,
	*onboardingAction,
}) // Optional (max 4)

_, err := activitysmith.Notifications.Send(request)
if err != nil {
	log.Fatal(err)
}
```

## Live Activities

<p align="center">
  <img src="https://cdn.activitysmith.com/features/metrics-live-activity-action.png" alt="Live Activities example" width="680" />
</p>

ActivitySmith supports two ways to drive Live Activities:

- Recommended: stream updates with `activitysmith.LiveActivities.Stream(...)`
- Advanced: manual lifecycle control with `Start`, `Update`, and `End`

Use stream updates when you want the easiest, stateless flow. You don't need to
store `activityID` or manage lifecycle state yourself. Send the latest state
for a stable `streamKey` and ActivitySmith will start or update the Live
Activity for you. When the tracked process is over, call `EndStream(...)`.

Use the manual lifecycle methods when you need direct control over a specific
Live Activity instance.

Live Activity UI types:

- `metrics`: best for live operational stats like server CPU and memory, queue depth, or replica lag
- `segmented_progress`: best for step-based workflows like deployments, backups, and ETL pipelines
- `progress`: best for continuous jobs like uploads, reindexes, and long-running migrations tracked as a percentage

### Recommended: Stream updates

Use a stable `streamKey` to identify the system or workflow you are tracking,
such as a server, deployment, build pipeline, cron job, or charging session.
This is especially useful for cron jobs and other scheduled tasks where you do
not want to store `activityID` between runs.

#### Metrics

<p align="center">
  <img src="https://cdn.activitysmith.com/features/metrics-live-activity-start.png" alt="Metrics stream example" width="680" />
</p>

```go
streamInput := activitysmithsdk.LiveActivityStreamInput{
	Title:    "Server Health",
	Subtitle: "prod-web-1",
	Type:     "metrics",
	Metrics: []generated.ActivityMetric{
		{Label: "CPU", Value: 9, Unit: generated.PtrString("%")},
		{Label: "MEM", Value: 45, Unit: generated.PtrString("%")},
	},
}

status, err := activitysmith.LiveActivities.Stream("prod-web-1", streamInput)
if err != nil {
	log.Fatal(err)
}
```

#### Segmented progress

<p align="center">
  <img src="https://cdn.activitysmith.com/features/update-live-activity.png" alt="Segmented progress stream example" width="680" />
</p>

```go
streamInput := activitysmithsdk.LiveActivityStreamInput{
	Title:         "Nightly Backup",
	Subtitle:      "upload archive",
	Type:          "segmented_progress",
	NumberOfSteps: 3,
	CurrentStep:   2,
}

_, err := activitysmith.LiveActivities.Stream("nightly-backup", streamInput)
if err != nil {
	log.Fatal(err)
}
```

#### Progress

<p align="center">
  <img src="https://cdn.activitysmith.com/features/progress-live-activity.png" alt="Progress stream example" width="680" />
</p>

```go
streamInput := activitysmithsdk.LiveActivityStreamInput{
	Title:      "Search Reindex",
	Subtitle:   "catalog-v2",
	Type:       "progress",
	Percentage: 42,
}

_, err := activitysmith.LiveActivities.Stream("search-reindex", streamInput)
if err != nil {
	log.Fatal(err)
}
```

Call `Stream(...)` again with the same `streamKey` whenever the state changes.

#### End a stream

Use this when the tracked process is finished and you no longer want the Live
Activity on devices. `content_state` is optional here; include it if you want
to end the stream with a final state.

```go
endInput := activitysmithsdk.LiveActivityStreamEndInput{
	Title:    "Server Health",
	Subtitle: "prod-web-1",
	Type:     "metrics",
	Metrics: []generated.ActivityMetric{
		{Label: "CPU", Value: 7, Unit: generated.PtrString("%")},
		{Label: "MEM", Value: 38, Unit: generated.PtrString("%")},
	},
}

_, err := activitysmith.LiveActivities.EndStream("prod-web-1", endInput)
if err != nil {
	log.Fatal(err)
}
```

If you later send another `Stream(...)` request with the same `streamKey`,
ActivitySmith starts a new Live Activity for that stream again.

Stream responses include an `Operation` field:

- `started`: ActivitySmith started a new Live Activity for this `streamKey`
- `updated`: ActivitySmith updated the current Live Activity
- `rotated`: ActivitySmith ended the previous Live Activity and started a new one
- `noop`: the incoming state matched the current state, so no update was sent
- `paused`: the stream is paused, so no Live Activity was started or updated
- `ended`: returned by `EndStream(...)` after the stream is ended

### Advanced: Manual lifecycle control

Use these methods when you want to manage the Live Activity lifecycle yourself.

#### Shared flow

1. Call `activitysmith.LiveActivities.Start(...)`.
2. Save the returned `activityID`.
3. Call `activitysmith.LiveActivities.Update(...)` as progress changes.
4. Call `activitysmith.LiveActivities.End(...)` when the work is finished.

### Metrics Type

Use `metrics` when you want to keep a small set of live stats visible, such as
server health, queue pressure, or database load.

#### Start

<p align="center">
  <img src="https://cdn.activitysmith.com/features/metrics-live-activity-start.png" alt="Metrics start example" width="680" />
</p>

```go
startInput := activitysmithsdk.LiveActivityStartInput{
	Title:    "Server Health",
	Subtitle: "prod-web-1",
	Type:     "metrics",
	Metrics: []generated.ActivityMetric{
		{Label: "CPU", Value: 9, Unit: generated.PtrString("%")},
		{Label: "MEM", Value: 45, Unit: generated.PtrString("%")},
	},
}

start, err := activitysmith.LiveActivities.Start(startInput)
if err != nil {
	log.Fatal(err)
}

activityID := start.GetActivityId()
```

#### Update

<p align="center">
  <img src="https://cdn.activitysmith.com/features/metrics-live-activity-update.png" alt="Metrics update example" width="680" />
</p>

```go
updateInput := activitysmithsdk.LiveActivityUpdateInput{
	ActivityID: activityID,
	Title:      "Server Health",
	Subtitle:   "prod-web-1",
	Type:       "metrics",
	Metrics: []generated.ActivityMetric{
		{Label: "CPU", Value: 76, Unit: generated.PtrString("%")},
		{Label: "MEM", Value: 52, Unit: generated.PtrString("%")},
	},
}

_, err := activitysmith.LiveActivities.Update(updateInput)
if err != nil {
	log.Fatal(err)
}
```

#### End

<p align="center">
  <img src="https://cdn.activitysmith.com/features/metrics-live-activity-end.png" alt="Metrics end example" width="680" />
</p>

```go
endInput := activitysmithsdk.LiveActivityEndInput{
	ActivityID: activityID,
	Title:      "Server Health",
	Subtitle:   "prod-web-1",
	Type:       "metrics",
	Metrics: []generated.ActivityMetric{
		{Label: "CPU", Value: 7, Unit: generated.PtrString("%")},
		{Label: "MEM", Value: 38, Unit: generated.PtrString("%")},
	},
	AutoDismissMinutes: 2,
}

_, err := activitysmith.LiveActivities.End(endInput)
if err != nil {
	log.Fatal(err)
}
```

### Segmented Progress Type

Use `segmented_progress` when progress is easier to follow as steps instead of a
raw percentage. It fits jobs like backups, deployments, ETL pipelines, and
checklists where "step 2 of 3" is more useful than "67%". `NumberOfSteps` is
dynamic, so you can increase or decrease it later if the workflow changes.

#### Start

<p align="center">
  <img src="https://cdn.activitysmith.com/features/start-live-activity.png" alt="Segmented progress start example" width="680" />
</p>

```go
startInput := activitysmithsdk.LiveActivityStartInput{
	Title:         "Nightly database backup",
	Subtitle:      "create snapshot",
	NumberOfSteps: 3,
	CurrentStep:   1,
	Type:          "segmented_progress",
	Color:         "yellow",
}

start, err := activitysmith.LiveActivities.Start(startInput)
if err != nil {
	log.Fatal(err)
}

activityID := start.GetActivityId()
```

#### Update

<p align="center">
  <img src="https://cdn.activitysmith.com/features/update-live-activity.png" alt="Segmented progress update example" width="680" />
</p>

```go
updateInput := activitysmithsdk.LiveActivityUpdateInput{
	ActivityID:    activityID,
	Title:         "Nightly database backup",
	Subtitle:      "upload archive",
	NumberOfSteps: 3,
	CurrentStep:   2,
}

_, err := activitysmith.LiveActivities.Update(updateInput)
if err != nil {
	log.Fatal(err)
}
```

#### End

<p align="center">
  <img src="https://cdn.activitysmith.com/features/end-live-activity.png" alt="Segmented progress end example" width="680" />
</p>

```go
endInput := activitysmithsdk.LiveActivityEndInput{
	ActivityID:         activityID,
	Title:              "Nightly database backup",
	Subtitle:           "verify restore",
	NumberOfSteps:      3,
	CurrentStep:        3,
	AutoDismissMinutes: 2,
}

_, err := activitysmith.LiveActivities.End(endInput)
if err != nil {
	log.Fatal(err)
}
```

### Progress Type

Use `progress` when the state is naturally continuous. It fits charging,
downloads, sync jobs, uploads, timers, and any flow where a percentage or
numeric range is the clearest signal.

#### Start

<p align="center">
  <img src="https://cdn.activitysmith.com/features/progress-live-activity-start.png" alt="Progress start example" width="680" />
</p>

```go
startInput := activitysmithsdk.LiveActivityStartInput{
	Title:      "EV Charging",
	Subtitle:   "Added 30 mi range",
	Type:       "progress",
	Percentage: 15,
}

start, err := activitysmith.LiveActivities.Start(startInput)
if err != nil {
	log.Fatal(err)
}

activityID := start.GetActivityId()
```

#### Update

<p align="center">
  <img src="https://cdn.activitysmith.com/features/progress-live-activity-update.png" alt="Progress update example" width="680" />
</p>

```go
updateInput := activitysmithsdk.LiveActivityUpdateInput{
	ActivityID: activityID,
	Title:      "EV Charging",
	Subtitle:   "Added 120 mi range",
	Percentage: 60,
}

_, err := activitysmith.LiveActivities.Update(updateInput)
if err != nil {
	log.Fatal(err)
}
```

#### End

<p align="center">
  <img src="https://cdn.activitysmith.com/features/progress-live-activity-end.png" alt="Progress end example" width="680" />
</p>

```go
endInput := activitysmithsdk.LiveActivityEndInput{
	ActivityID:         activityID,
	Title:              "EV Charging",
	Subtitle:           "Added 200 mi range",
	Percentage:         100,
	AutoDismissMinutes: 2,
}

_, err := activitysmith.LiveActivities.End(endInput)
if err != nil {
	log.Fatal(err)
}
```

### Live Activity Action

Just like Actionable Push Notifications, Live Activities can have a button that opens a URL in a browser or triggers a webhook. Webhooks are executed by the ActivitySmith backend.

<p align="center">
  <img src="https://cdn.activitysmith.com/features/metrics-live-activity-action.png" alt="Metrics Live Activity with action" width="680" />
</p>

#### Open URL action

```go
startInput := activitysmithsdk.LiveActivityStartInput{
	Title:    "Server Health",
	Subtitle: "prod-web-1",
	Type:     "metrics",
	Metrics: []generated.ActivityMetric{
		{Label: "CPU", Value: 76, Unit: generated.PtrString("%")},
		{Label: "MEM", Value: 52, Unit: generated.PtrString("%")},
	},
	Action: &activitysmithsdk.LiveActivityActionInput{
		Title: "Open Dashboard",
		Type:  "open_url",
		URL:   "https://ops.example.com/servers/prod-web-1",
	},
}

start, err := activitysmith.LiveActivities.Start(startInput)
if err != nil {
	log.Fatal(err)
}

activityID := start.GetActivityId()
```

#### Webhook action

```go
updateInput := activitysmithsdk.LiveActivityUpdateInput{
	ActivityID: activityID,
	Title:      "Server Health",
	Subtitle:   "prod-web-1",
	Type:       "metrics",
	Metrics: []generated.ActivityMetric{
		{Label: "CPU", Value: 91, Unit: generated.PtrString("%")},
		{Label: "MEM", Value: 57, Unit: generated.PtrString("%")},
	},
	Action: &activitysmithsdk.LiveActivityActionInput{
		Title:  "Restart Service",
		Type:   "webhook",
		URL:    "https://ops.example.com/hooks/servers/prod-web-1/restart",
		Method: "POST",
		Body: map[string]interface{}{
			"server_id":    "prod-web-1",
			"requested_by": "activitysmith-go",
		},
	},
}

_, err = activitysmith.LiveActivities.Update(updateInput)
if err != nil {
	log.Fatal(err)
}
```

## Channels

Channels are used to target specific team members or devices. Can be used for both push notifications and live activities.

```go
request := generated.NewPushNotificationRequest("New subscription 💸")
request.SetMessage("Customer upgraded to Pro plan")
request.SetTarget(generated.ChannelTarget{Channels: []string{"sales", "customer-success"}}) // Optional

_, err := activitysmith.Notifications.Send(request)
if err != nil {
	log.Fatal(err)
}
```

## Error Handling

SDK methods return `(response, error)`. Always check `error` on each call.

## Requirements

- Go 1.22+

## License

MIT
