# ActivitySmith Go SDK

The ActivitySmith Go SDK provides convenient access to the ActivitySmith API from Go applications.

## Documentation

See [API reference](https://activitysmith.com/docs/api-reference/introduction).

## Table of Contents

- [Installation](#installation)
- [Setup](#setup)
- [Push Notifications](#push-notifications)
  - [Send a Push Notification](#send-a-push-notification)
  - [Rich Push Notifications with Media](#rich-push-notifications-with-media)
  - [Actionable Push Notifications](#actionable-push-notifications)
- [Live Activities](#live-activities)
  - [Simple: Let ActivitySmith manage the Live Activity for you](#simple-let-activitysmith-manage-the-live-activity-for-you)
  - [Advanced: Full lifecycle control](#advanced-full-lifecycle-control)
  - [Stats Type](#stats-type)
  - [Metrics Type](#metrics-type)
  - [Segmented Progress Type](#segmented-progress-type)
  - [Progress Type](#progress-type)
  - [Live Activity Action](#live-activity-action)
- [Channels](#channels)
- [Widgets](#widgets)

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
_, err := activitysmith.Notifications.Send(activitysmithsdk.PushNotificationInput{
	Title:       "New subscription 💸",
	Message:     "Customer upgraded to Pro plan",
	Redirection: "https://crm.example.com/customers/cus_9f3a1d",
	Actions: []activitysmithsdk.PushNotificationAction{
		activitysmithsdk.PushAction(
			"Open CRM Profile",
			"open_url",
			"https://crm.example.com/customers/cus_9f3a1d",
		),
		activitysmithsdk.PushAction(
			"Start Onboarding Workflow",
			"webhook",
			"https://hooks.example.com/activitysmith/onboarding/start",
			activitysmithsdk.PushActionMethod("POST"),
			activitysmithsdk.PushActionBody(map[string]interface{}{
				"customer_id": "cus_9f3a1d",
				"plan":        "pro",
			}),
		),
	},
})
if err != nil {
	log.Fatal(err)
}
```

## Live Activities

<p align="center">
  <img src="https://cdn.activitysmith.com/features/metrics-live-activity-action.png" alt="Metrics Live Activity screenshot" width="680" />
</p>

There are four types of Live Activities:

- `stats`: best for compact business or product stats like revenue, orders, conversion, and average order value
- `metrics`: best for live operational stats like server CPU and memory, queue depth, or replica lag
- `segmented_progress`: best for step-based workflows like deployments, backups, and ETL pipelines
- `progress`: best for continuous jobs like uploads, reindexes, and long-running migrations tracked as a percentage

When working with Live Activities via our API, you have two approaches tailored
to different needs. First, the stateless mode is the simplest path - one API
call can initiate or update an activity, and another ends it - no state
tracking on your side.

This is ideal if you want minimal complexity, perfect for automated workflows
like cron jobs.

In contrast, if you need precise lifecycle control, the classic approach offers
distinct calls for start, updates, and end, giving you full control over the
activity's state.

In the following sections, we'll break down how to implement each method so you
can choose what fits your use case best.

### Simple: Let ActivitySmith manage the Live Activity for you

Use a stable `streamKey` to identify the system or workflow you are tracking,
such as a server, deployment, build pipeline, cron job, or charging session.
This is especially useful for cron jobs and other scheduled tasks where you do
not want to store `activityID` between runs.

#### Stats

<p align="center">
  <img src="https://cdn.activitysmith.com/features/stats-live-activity.png" alt="Stats stream example" width="680" />
</p>

```go
streamInput := activitysmithsdk.LiveActivityStreamInput{
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:    "Sales",
		Subtitle: "last hour",
		Type:     "stats",
		Metrics: []activitysmithsdk.ActivityMetric{
			activitysmithsdk.Metric("Revenue", "$2430", activitysmithsdk.MetricColor("blue")),
			activitysmithsdk.Metric("Orders", "37", activitysmithsdk.MetricColor("green")),
			activitysmithsdk.Metric("Conversion", "4.8%", activitysmithsdk.MetricColor("magenta")),
			activitysmithsdk.Metric("Avg Order", "$65.68", activitysmithsdk.MetricColor("yellow")),
			activitysmithsdk.Metric("Refunds", "$84", activitysmithsdk.MetricColor("red")),
			activitysmithsdk.Metric("New Buyers", "18", activitysmithsdk.MetricColor("cyan")),
		},
	},
}

status, err := activitysmith.LiveActivities.Stream("sales-hourly", streamInput)
if err != nil {
	log.Fatal(err)
}
```

#### Metrics

<p align="center">
  <img src="https://cdn.activitysmith.com/features/metrics-live-activity-start.png" alt="Metrics stream example" width="680" />
</p>

```go
streamInput := activitysmithsdk.LiveActivityStreamInput{
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:    "Server Health",
		Subtitle: "prod-web-1",
		Type:     "metrics",
		Metrics: []activitysmithsdk.ActivityMetric{
			activitysmithsdk.Metric("CPU", 9, activitysmithsdk.MetricUnit("%")),
			activitysmithsdk.Metric("MEM", 45, activitysmithsdk.MetricUnit("%")),
		},
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
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:         "Nightly Backup",
		Subtitle:      "upload archive",
		Type:          "segmented_progress",
		NumberOfSteps: 3,
		CurrentStep:   2,
	},
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
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:      "Search Reindex",
		Subtitle:   "catalog-v2",
		Type:       "progress",
		Percentage: 42,
	},
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
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:    "Server Health",
		Subtitle: "prod-web-1",
		Type:     "metrics",
		Metrics: []activitysmithsdk.ActivityMetric{
			activitysmithsdk.Metric("CPU", 7, activitysmithsdk.MetricUnit("%")),
			activitysmithsdk.Metric("MEM", 38, activitysmithsdk.MetricUnit("%")),
		},
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

### Advanced: Full lifecycle control

Use these methods when you want to manage the Live Activity lifecycle yourself:

1. Call `activitysmith.LiveActivities.Start(...)`.
2. Save the returned `activityID`.
3. Call `activitysmith.LiveActivities.Update(...)` as progress changes.
4. Call `activitysmith.LiveActivities.End(...)` when the work is finished.

### Stats Type

Keep your key numbers on your Lock Screen. `stats` fits up to 8 labeled values,
such as revenue, orders, conversion, uptime, or any other business metric you
want visible at a glance. Each metric can use a formatted string or number as
its `Value`. Add `Color` to a metric to show an accent dot next to its label;
omit `Color` to show the label without a dot.

#### Start

<p align="center">
  <img src="https://cdn.activitysmith.com/features/stats-live-activity.png" alt="Stats Live Activity with sales revenue, orders, conversion, and average order value" width="680" />
</p>

```go
startInput := activitysmithsdk.LiveActivityStartInput{
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:    "Sales",
		Subtitle: "last hour",
		Type:     "stats",
		Metrics: []activitysmithsdk.ActivityMetric{
			activitysmithsdk.Metric("Revenue", "$2430", activitysmithsdk.MetricColor("blue")),
			activitysmithsdk.Metric("Orders", "37", activitysmithsdk.MetricColor("green")),
			activitysmithsdk.Metric("Conversion", "4.8%", activitysmithsdk.MetricColor("magenta")),
			activitysmithsdk.Metric("Avg Order", "$65.68", activitysmithsdk.MetricColor("yellow")),
			activitysmithsdk.Metric("Refunds", "$84", activitysmithsdk.MetricColor("red")),
			activitysmithsdk.Metric("New Buyers", "18", activitysmithsdk.MetricColor("cyan")),
		},
	},
}

start, err := activitysmith.LiveActivities.Start(startInput)
if err != nil {
	log.Fatal(err)
}

activityID := start.GetActivityId()
```

#### Update

```go
updateInput := activitysmithsdk.LiveActivityUpdateInput{
	ActivityID: activityID,
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:    "Sales",
		Subtitle: "last hour",
		Type:     "stats",
		Metrics: []activitysmithsdk.ActivityMetric{
			activitysmithsdk.Metric("Revenue", "$3180", activitysmithsdk.MetricColor("blue")),
			activitysmithsdk.Metric("Orders", "51", activitysmithsdk.MetricColor("green")),
			activitysmithsdk.Metric("Conversion", "5.2%", activitysmithsdk.MetricColor("magenta")),
			activitysmithsdk.Metric("Avg Order", "$62.35", activitysmithsdk.MetricColor("yellow")),
			activitysmithsdk.Metric("Refunds", "$126", activitysmithsdk.MetricColor("red")),
			activitysmithsdk.Metric("New Buyers", "24", activitysmithsdk.MetricColor("cyan")),
		},
	},
}

_, err := activitysmith.LiveActivities.Update(updateInput)
if err != nil {
	log.Fatal(err)
}
```

#### End

```go
endInput := activitysmithsdk.LiveActivityEndInput{
	ActivityID: activityID,
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:    "Sales",
		Subtitle: "last hour",
		Type:     "stats",
		Metrics: []activitysmithsdk.ActivityMetric{
			activitysmithsdk.Metric("Revenue", "$3460", activitysmithsdk.MetricColor("blue")),
			activitysmithsdk.Metric("Orders", "58", activitysmithsdk.MetricColor("green")),
			activitysmithsdk.Metric("Conversion", "5.4%", activitysmithsdk.MetricColor("magenta")),
			activitysmithsdk.Metric("Avg Order", "$59.66", activitysmithsdk.MetricColor("yellow")),
			activitysmithsdk.Metric("Refunds", "$92", activitysmithsdk.MetricColor("red")),
			activitysmithsdk.Metric("New Buyers", "31", activitysmithsdk.MetricColor("cyan")),
		},
		AutoDismissMinutes: 2,
	},
}

_, err := activitysmith.LiveActivities.End(endInput)
if err != nil {
	log.Fatal(err)
}
```

### Metrics Type

Use `metrics` when you want to keep a small set of live stats visible, such as
server health, queue pressure, or database load.

#### Start

<p align="center">
  <img src="https://cdn.activitysmith.com/features/metrics-live-activity-start.png" alt="Metrics start example" width="680" />
</p>

```go
startInput := activitysmithsdk.LiveActivityStartInput{
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:    "Server Health",
		Subtitle: "prod-web-1",
		Type:     "metrics",
		Metrics: []activitysmithsdk.ActivityMetric{
			activitysmithsdk.Metric("CPU", 9, activitysmithsdk.MetricUnit("%")),
			activitysmithsdk.Metric("MEM", 45, activitysmithsdk.MetricUnit("%")),
		},
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
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:    "Server Health",
		Subtitle: "prod-web-1",
		Type:     "metrics",
		Metrics: []activitysmithsdk.ActivityMetric{
			activitysmithsdk.Metric("CPU", 76, activitysmithsdk.MetricUnit("%")),
			activitysmithsdk.Metric("MEM", 52, activitysmithsdk.MetricUnit("%")),
		},
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
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:    "Server Health",
		Subtitle: "prod-web-1",
		Type:     "metrics",
		Metrics: []activitysmithsdk.ActivityMetric{
			activitysmithsdk.Metric("CPU", 7, activitysmithsdk.MetricUnit("%")),
			activitysmithsdk.Metric("MEM", 38, activitysmithsdk.MetricUnit("%")),
		},
		AutoDismissMinutes: 2,
	},
}

_, err := activitysmith.LiveActivities.End(endInput)
if err != nil {
	log.Fatal(err)
}
```

### Segmented Progress Type

Use `segmented_progress` for jobs and workflows that move through clear steps or
phases. It fits jobs like backups, deployments, ETL pipelines, and checklists.
`NumberOfSteps` is dynamic, so you can increase or decrease it later if the
workflow changes.

#### Start

<p align="center">
  <img src="https://cdn.activitysmith.com/features/start-live-activity.png" alt="Segmented progress start example" width="680" />
</p>

```go
startInput := activitysmithsdk.LiveActivityStartInput{
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:         "Nightly database backup",
		Subtitle:      "create snapshot",
		Type:          "segmented_progress",
		NumberOfSteps: 3,
		CurrentStep:   1,
		Color:         "yellow",
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
  <img src="https://cdn.activitysmith.com/features/update-live-activity.png" alt="Segmented progress update example" width="680" />
</p>

```go
updateInput := activitysmithsdk.LiveActivityUpdateInput{
	ActivityID: activityID,
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:         "Nightly database backup",
		Subtitle:      "upload archive",
		NumberOfSteps: 3,
		CurrentStep:   2,
	},
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
	ActivityID: activityID,
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:              "Nightly database backup",
		Subtitle:           "verify restore",
		NumberOfSteps:      3,
		CurrentStep:        3,
		AutoDismissMinutes: 2,
	},
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
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:      "EV Charging",
		Subtitle:   "Added 30 mi range",
		Type:       "progress",
		Percentage: 15,
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
  <img src="https://cdn.activitysmith.com/features/progress-live-activity-update.png" alt="Progress update example" width="680" />
</p>

```go
updateInput := activitysmithsdk.LiveActivityUpdateInput{
	ActivityID: activityID,
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:      "EV Charging",
		Subtitle:   "Added 120 mi range",
		Percentage: 60,
	},
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
	ActivityID: activityID,
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:              "EV Charging",
		Subtitle:           "Added 200 mi range",
		Percentage:         100,
		AutoDismissMinutes: 2,
	},
}

_, err := activitysmith.LiveActivities.End(endInput)
if err != nil {
	log.Fatal(err)
}
```

### Live Activity Action

Just like Actionable Push Notifications, Live Activities can have a button that opens provided URL in a browser or triggers a webhook. Webhooks are executed by the ActivitySmith backend.

<p align="center">
  <img src="https://cdn.activitysmith.com/features/metrics-live-activity-action.png" alt="Metrics Live Activity with action" width="680" />
</p>

#### Open URL action

```go
startInput := activitysmithsdk.LiveActivityStartInput{
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:    "Server Health",
		Subtitle: "prod-web-1",
		Type:     "metrics",
		Metrics: []activitysmithsdk.ActivityMetric{
			activitysmithsdk.Metric("CPU", 76, activitysmithsdk.MetricUnit("%")),
			activitysmithsdk.Metric("MEM", 52, activitysmithsdk.MetricUnit("%")),
		},
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

<p align="center">
  <img src="https://cdn.activitysmith.com/features/live-activity-with-action.png?v=20260319-1" alt="Live Activity with action" width="680" />
</p>

```go
updateInput := activitysmithsdk.LiveActivityUpdateInput{
	ActivityID: activityID,
	ContentState: activitysmithsdk.LiveActivityContentStateInput{
		Title:         "Reindexing product search",
		Subtitle:      "Shard 7 of 12",
		NumberOfSteps: 12,
		CurrentStep:   7,
	},
	Action: &activitysmithsdk.LiveActivityActionInput{
		Title:  "Pause Reindex",
		Type:   "webhook",
		URL:    "https://ops.example.com/hooks/search/reindex/pause",
		Method: "POST",
		Body: map[string]interface{}{
			"job_id":       "reindex-2026-03-19",
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

## Widgets

<p align="center">
  <img src="https://cdn.activitysmith.com/features/lock-screen-widgets.png" alt="Lock screen widgets" width="680" />
</p>

ActivitySmith lets you display any value on your Lock Screen with widgets - SaaS metrics, revenue, signups, uptime, habits, or anything else you want to track. Create a metric in the <a href="https://activitysmith.com/app/widgets" target="_blank" rel="noopener noreferrer">web app</a>, then update the metric value using our API, add a widget to your lock screen and it will fetch the latest update automatically.

<p align="center">
  <img src="https://cdn.activitysmith.com/features/create-widget-metric.png" alt="Create widget metric" width="680" />
</p>

```go
_, err := activitysmith.Metrics.Update("deploy.success_rate", 99.9)
if err != nil {
	log.Fatal(err)
}
```

String metric values work too.

```go
_, err = activitysmith.Metrics.Update("prod.status", "healthy")
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
