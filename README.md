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
  - [Start & Update Live Activity](#start--update-live-activity)
  - [End Live Activity](#end-live-activity)
  - [Live Activity Action](#live-activity-action)
  - [Icons and Badges](#icons-and-badges)
  - [Live Activity Colors](#live-activity-colors)
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

Push notification `Redirection` and `Actions` are optional. Use them to open HTTPS URLs, run a specific iPhone Shortcut with `shortcuts://run-shortcut?name=...`, or trigger backend webhook workflows.
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
			"Chat with Jarvis",
			"open_url",
			"shortcuts://run-shortcut?name=Jarvis",
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

There are six types of Live Activities:

- `stats`: best for showing business numbers side by side, such as revenue, sales, new users, conversion, refunds, or any other value you want visible at a glance
- `metrics`: best for live percentage values that change often, like server CPU, memory usage, disk usage, or error rate
- `segmented_progress`: best for anything that moves through clear stages, like deployments, onboarding flows, backups, ETL pipelines, migrations, and AI agent runs
- `progress`: best for tracking real-time progress with percentage, like tasks, backups, migrations, syncs, or uploads
- `alert`: best for status updates, such as feature adoption, reactivation, onboarding blockers, incidents, escalations, and other operational states
- `timer`: best for countdowns and elapsed runtime, like benchmark runs, uploads, backups, transcodes, and long-running jobs

### Start & Update Live Activity

Use a stable `streamKey` to identify the metric, job, deployment, or system you want to keep visible. The first `Stream(...)` call starts the Live Activity. Later calls with the same `streamKey` update it.

#### Stats

<p align="center">
  <img
    src="https://cdn.activitysmith.com/features/stats-live-activity.png"
    alt="Stats Live Activity stream example"
    width="680"
  />
</p>

```go
activitysmith.LiveActivities.Stream(
	"sales-hourly",
	activitysmithsdk.LiveActivityStreamInput{
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
)
```

#### Metrics

<p align="center">
  <img
    src="https://cdn.activitysmith.com/features/metrics-live-activity-start.png"
    alt="Metrics Live Activity stream example"
    width="680"
  />
</p>

```go
activitysmith.LiveActivities.Stream(
	"prod-web-1",
	activitysmithsdk.LiveActivityStreamInput{
		Title:    "Server Health",
		Subtitle: "prod-web-1",
		Type:     "metrics",
		Metrics: []activitysmithsdk.ActivityMetric{
			activitysmithsdk.Metric("CPU", 9, activitysmithsdk.MetricUnit("%")),
			activitysmithsdk.Metric("MEM", 45, activitysmithsdk.MetricUnit("%")),
		},
	},
)
```

#### Segmented Progress

<p align="center">
  <img
    src="https://cdn.activitysmith.com/features/update-live-activity.png"
    alt="Segmented Progress Live Activity stream example"
    width="680"
  />
</p>

```go
activitysmith.LiveActivities.Stream(
	"nightly-backup",
	activitysmithsdk.LiveActivityStreamInput{
		Title:         "Nightly Backup",
		Subtitle:      "upload archive",
		Type:          "segmented_progress",
		NumberOfSteps: 3,
		CurrentStep:   2,
	},
)
```

#### Progress

<p align="center">
  <img
    src="https://cdn.activitysmith.com/features/progress-live-activity.png"
    alt="Progress Live Activity stream example"
    width="680"
  />
</p>

```go
activitysmith.LiveActivities.Stream(
	"search-reindex",
	activitysmithsdk.LiveActivityStreamInput{
		Title:      "Search Reindex",
		Subtitle:   "catalog-v2",
		Type:       "progress",
		Percentage: 42,
	},
)
```

#### Alert

<p align="center">
  <img
    src="https://cdn.activitysmith.com/features/alert-live-activity.png"
    alt="Alert Live Activity stream example"
    width="680"
  />
</p>

```go
activitysmith.LiveActivities.Stream(
	"customer-ops",
	activitysmithsdk.LiveActivityStreamInput{
		Title:   "Reactivation",
		Message: "Lumen came back after 2 weeks",
		Type:    activitysmithsdk.LiveActivityTypeAlert,
		Icon:    activitysmithsdk.AlertIcon("cloud.sun", "yellow"),
		Badge:   activitysmithsdk.AlertBadge("Customer", "magenta"),
	},
)
```

#### Timer

<p align="center">
  <img
    src="https://cdn.activitysmith.com/features/timer-live-activity.png"
    alt="Timer Live Activity showing a benchmark run countdown"
    width="680"
  />
</p>

```go
activitysmith.LiveActivities.Stream(
	"benchmark-run",
	activitysmithsdk.LiveActivityStreamInput{
		Title:           "Benchmark Run",
		Subtitle:        "sampling",
		Type:            "timer",
		DurationSeconds: 300,
		Color:           "cyan",
	},
)
```

For a countdown, send `duration_seconds`. You can update `title`, `subtitle`, `color`, or any other visible field as the work changes. Leave `duration_seconds` out unless you want to change the timer.

To start at 00:00 and count up, set `counts_down: false` and leave out `duration_seconds`.

### End Live Activity

Call `EndStream(...)` with the same `streamKey` to dismiss the Live Activity. You can include final values before it is removed. By default, iOS removes the Live Activity after two minutes. Set `AutoDismissMinutes` to choose a different dismissal time, including `0` for immediate dismissal.

```go
activitysmith.LiveActivities.EndStream(
	"prod-web-1",
	activitysmithsdk.LiveActivityStreamEndInput{
		Title:    "Server Health",
		Subtitle: "prod-web-1",
		Type:     "metrics",
		Metrics: []activitysmithsdk.ActivityMetric{
			activitysmithsdk.Metric("CPU", 7, activitysmithsdk.MetricUnit("%")),
			activitysmithsdk.Metric("MEM", 38, activitysmithsdk.MetricUnit("%")),
		},
		AutoDismissMinutes: 2,
	},
)
```

### Live Activity Action

Live Activities can include one optional action button.

- `open_url`: open an HTTPS URL.
- `open_url` with a `shortcuts://run-shortcut?name=...` URL: run a specific iPhone Shortcut, for example to open an app.
- `webhook`: trigger a backend GET/POST workflow.

<p align="center">
  <img
    src="https://cdn.activitysmith.com/features/live-activity-with-action.png?v=20260319-1"
    alt="Live Activity with action button"
    width="680"
  />
</p>

#### Open URL action

```go
activitysmith.LiveActivities.Stream(
	"prod-web-1",
	activitysmithsdk.LiveActivityStreamInput{
		Title:    "Server Health",
		Subtitle: "prod-web-1",
		Type:     "metrics",
		Metrics: []activitysmithsdk.ActivityMetric{
			activitysmithsdk.Metric("CPU", 76, activitysmithsdk.MetricUnit("%")),
			activitysmithsdk.Metric("MEM", 52, activitysmithsdk.MetricUnit("%")),
		},
		Action: &activitysmithsdk.LiveActivityActionInput{
			Title: "Dashboard",
			Type:  "open_url",
			URL:   "https://ops.example.com/servers/prod-web-1",
		},
	},
)
```

#### Apple Shortcut action

```go
activitysmith.LiveActivities.Stream(
	"deploy-payments-api",
	activitysmithsdk.LiveActivityStreamInput{
		Title:         "Deploying payments-api",
		Subtitle:      "Running database migrations",
		Type:          "segmented_progress",
		NumberOfSteps: 5,
		CurrentStep:   3,
		Action: &activitysmithsdk.LiveActivityActionInput{
			Title: "Chat with Jarvis",
			Type:  "open_url",
			URL:   "shortcuts://run-shortcut?name=Jarvis",
		},
	},
)
```

#### Webhook action

```go
activitysmith.LiveActivities.Stream(
	"search-reindex",
	activitysmithsdk.LiveActivityStreamInput{
		Title:         "Reindexing product search",
		Subtitle:      "Shard 7 of 12",
		Type:          "segmented_progress",
		NumberOfSteps: 12,
		CurrentStep:   7,
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
	},
)
```

### Icons and Badges

Add more context to Live Activities with icons and badges.

#### Icon

Supported Live Activity types: `stats`, `metrics`, `progress`, `segmented_progress`, `alert`, and `timer`.

<p align="center">
  <img
    src="https://cdn.activitysmith.com/features/metrics-live-activity-with-icon.png"
    alt="Metrics Live Activity with an SF Symbol icon on the iPhone Lock Screen"
    width="680"
  />
</p>

```go
activitysmith.LiveActivities.Stream(
	"prod-web-1",
	activitysmithsdk.LiveActivityStreamInput{
		Title:    "Server Health",
		Subtitle: "prod-web-1",
		Type:     "metrics",
		Icon:     activitysmithsdk.AlertIcon("server.rack", "blue"),
		Metrics: []activitysmithsdk.ActivityMetric{
			activitysmithsdk.Metric("CPU", 18, activitysmithsdk.MetricUnit("%")),
			activitysmithsdk.Metric("MEM", 42, activitysmithsdk.MetricUnit("%")),
		},
	},
)
```

The `Icon` symbol value is an Apple SF Symbol name. Browse the catalog with one of these tools:

- [ActivitySmith app](https://apps.apple.com/us/app/activitysmith/id6752254835) - Open Settings -> SF Symbols to browse 45 hand-picked icons ready to use
- [SF Symbols](https://developer.apple.com/sf-symbols/) - Apple's official macOS app
- [Interactful](https://apps.apple.com/app/interactful/id1528095640) - free third-party iOS app listing all SF Symbols under Foundations -> Iconography

#### Badge

Badges are supported by `alert`, `progress`, and `segmented_progress` Live Activities.

<p align="center">
  <img
    src="https://cdn.activitysmith.com/features/progress-live-activity-with-badge.png"
    alt="Progress Live Activity with a badge on the iPhone Lock Screen"
    width="680"
  />
</p>

```go
activitysmith.LiveActivities.Stream(
	"nightly-database-backup",
	activitysmithsdk.LiveActivityStreamInput{
		Title:      "Nightly Database Backup",
		Subtitle:   "verify restore",
		Type:       "progress",
		Badge:      activitysmithsdk.AlertBadge("S3", "cyan"),
		Percentage: 62,
	},
)
```

### Live Activity Colors

Choose from these colors for the Live Activity accent, including progress bars and action buttons, or apply them to an individual icon or badge:

`lime`, `green`, `cyan`, `blue`, `purple`, `magenta`, `red`, `orange`, `yellow`, `gray`

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
