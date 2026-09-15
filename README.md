# ActivitySmith Go SDK

[Documentation](https://activitysmith.com/docs/sdks/go)

## Installation

Install the ActivitySmith Go SDK with `go get`:

```bash
go get github.com/ActivitySmithHQ/activitysmith-go
```

## Quickstart

1. [Create an API key](https://activitysmith.com/app/keys)
2. Pass the API key into `activitysmithsdk.New`.

```go
package main

import (
	"os"

	activitysmithsdk "github.com/ActivitySmithHQ/activitysmith-go"
)

func main() {
	activitysmith, _ := activitysmithsdk.New(os.Getenv("ACTIVITYSMITH_API_KEY"))
}
```

## Push Notifications

### Send a Push Notification

Send an immediate notification for a completed task or event.

![Push Notification example for a new subscription event](https://cdn.activitysmith.com/features/new-subscription-push-notification.png)

```go
input := activitysmithsdk.PushNotificationInput{
	Title:   "New subscription 💸",
	Message: "Customer upgraded to Pro plan",
}

activitysmith.Notifications.Send(input)
```

### Rich Push Notifications with Media

![Rich Push Notification with image](https://cdn.activitysmith.com/features/rich-push-notification-with-image.png)

```go
input := activitysmithsdk.PushNotificationInput{
	Title:   "Homepage ready",
	Message: "Your agent finished the redesign.",
	Media:   "https://cdn.example.com/output/homepage-v2.png",
}

activitysmith.Notifications.Send(input)
```

Attach images, videos, or audio to your Push Notifications. Press and hold the notification to preview the media.

![Rich Push Notification with audio](https://cdn.activitysmith.com/features/rich-push-notification-with-audio.png)

What will work:

- direct image URL: `.jpg`, `.png`, `.gif`, etc.
- direct audio file URL: `.mp3`, `.m4a`, etc.
- direct video file URL: `.mp4`, `.mov`, etc.
- URL that responds with a proper media `Content-Type`, even if the path has no extension

`Media` cannot be combined with `Actions`.

### Push Notifications with Redirection

Open a web page, run an iOS Shortcut, or open an app when someone taps the notification. `Redirection` supports:

- **HTTP/HTTPS:** Web pages, e.g. `https://example.com`
- **Shortcuts:** Run Jarvis with `shortcuts://run-shortcut?name=Jarvis` <!-- full-width -->
- **App deep links:** Installed apps or specific content within them
  - **Spotify:** A track, e.g. `spotify:track:6rqhFgbbKwnb9MLmUQDhG6`
  - **Termius:** `termius://` to open the app
  - **Claude:** `claude://code` to open the Code tab
  - **ChatGPT:** `chatgpt://` to open the app <!-- Verify ChatGPT URL scheme on iOS before publishing -->

```go
input := activitysmithsdk.PushNotificationInput{
	Title:       "Homepage ready",
	Message:     "Your agent finished the redesign.",
	Redirection: "https://github.com/acme/web/pull/482",
}

activitysmith.Notifications.Send(input)
```

### Actionable Push Notifications

![Actionable Push Notification with redirection and actions](https://cdn.activitysmith.com/features/actionable-push-notifications-2.png)

`open_url` actions open a web page, run an iOS Shortcut, or open an app when someone taps the button. Supported links:

- **HTTP/HTTPS:** Web pages, e.g. `https://example.com`
- **Shortcuts:** Run Jarvis with `shortcuts://run-shortcut?name=Jarvis` <!-- full-width -->
- **App deep links:** Installed apps or specific content within them
  - **Spotify:** A track, e.g. `spotify:track:6rqhFgbbKwnb9MLmUQDhG6`
  - **Termius:** `termius://` to open the app
  - **Claude:** `claude://code` to open the Code tab
  - **ChatGPT:** `chatgpt://` to open the app <!-- Verify ChatGPT URL scheme on iOS before publishing -->

Webhooks are executed by the ActivitySmith backend and must use HTTPS.

```go
input := activitysmithsdk.PushNotificationInput{
	Title:       "New subscription 💸",
	Message:     "Customer upgraded to Pro plan",
	Actions: []activitysmithsdk.PushNotificationAction{
		activitysmithsdk.PushAction(
			"Open CRM",
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
}

activitysmith.Notifications.Send(input)
```

## Live Activities

Choose the Live Activity type that matches what you want to show:

![Stats Live Activity with six labeled sales metrics](https://cdn.activitysmith.com/features/stats-live-activity.png)

**Stats**: Show up to 8 labeled values on your Lock Screen, from revenue and orders to uptime and conversion.

![Metrics Live Activity with CPU and memory values](https://cdn.activitysmith.com/features/metrics-live-activity-start.png)

**Metrics**: Track two related values with segmented bars, such as CPU and memory.

![Segmented Progress Live Activity showing a workflow step](https://cdn.activitysmith.com/features/update-live-activity.png)

**Segmented Progress**: Show progress through a known set of steps, like build, test, deploy, and verify.

![Progress Live Activity showing percentage completion](https://cdn.activitysmith.com/features/progress-live-activity.png)

**Progress**: Show percentage progress for jobs that move continuously toward completion.

![Alert Live Activity showing a customer reactivation update](https://cdn.activitysmith.com/features/alert-live-activity.png)

**Alert**: Show status updates with a clear message, badge, and icon. When you add an action button, `color` controls the button tint.

![Timer Live Activity showing a benchmark run countdown](https://cdn.activitysmith.com/features/timer-live-activity.png)

**Timer**: Count down from a duration, or count up from 00:00 while a job runs.

### Start & Update Live Activity

Use a stable `streamKey` to identify the metric, job, deployment, or system you want to keep visible. The first `Stream(...)` call starts the Live Activity. Later calls with the same `streamKey` update it.

#### Stats

![Stats Live Activity stream example](https://cdn.activitysmith.com/features/stats-live-activity.png)

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

![Metrics Live Activity stream example](https://cdn.activitysmith.com/features/metrics-live-activity-start.png)

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

![Segmented Progress Live Activity stream example](https://cdn.activitysmith.com/features/update-live-activity.png)

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

![Progress Live Activity stream example](https://cdn.activitysmith.com/features/progress-live-activity.png)

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

![Alert Live Activity stream example](https://cdn.activitysmith.com/features/alert-live-activity.png)

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

![Timer Live Activity stream example](https://cdn.activitysmith.com/features/timer-live-activity.png)

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

For a countdown, send `DurationSeconds`. You can update `Title`, `Subtitle`, `Color`, or any other visible field as the work changes. Leave `DurationSeconds` out unless you want to change the timer.

To start at 00:00 and count up, set `CountsDown` to `false` and leave out `DurationSeconds`.

### End Live Activity

Call `EndStream(...)` with the same `streamKey` to dismiss the Live Activity. You can include final values before it is removed. Set `AutoDismissSeconds` or `AutoDismissMinutes` on `ContentState` to delay dismissal. Seconds take precedence when both are supplied. Use `WithAutoDismissSeconds(0)` or `WithAutoDismissMinutes(0)` for immediate dismissal.

```go
activitysmith.LiveActivities.EndStream(
    "prod-web-1",
    activitysmithsdk.LiveActivityStreamEndInput{
        ContentState: activitysmithsdk.LiveActivityContentStateInput{
            Title: "Server Health",
            Subtitle: "prod-web-1",
            Type: "metrics",
            Metrics: []activitysmithsdk.ActivityMetric{
                activitysmithsdk.Metric("CPU", 7, activitysmithsdk.MetricUnit("%")),
                activitysmithsdk.Metric("MEM", 38, activitysmithsdk.MetricUnit("%")),
            },
            AutoDismissSeconds: 30,
        },
    },
)
```

### Icons and Badges

Add more context to Live Activities with icons and badges.

#### Icon

Supported Live Activity types: `stats`, `metrics`, `progress`, `segmented_progress`, `alert`, and `timer`.

![Metrics Live Activity with an SF Symbol icon on the iPhone Lock Screen](https://cdn.activitysmith.com/features/metrics-live-activity-with-icon.png)

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

![Progress Live Activity with a badge on the iPhone Lock Screen](https://cdn.activitysmith.com/features/progress-live-activity-with-badge.png)

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

### Live Activity Action

![Metrics Live Activity with action](https://cdn.activitysmith.com/features/metrics-live-activity-action.png)

Live Activities can include an action button.

- `open_url`: Open a web page or run an iOS Shortcut
- `webhook`: Trigger a backend GET/POST workflow

#### Open URL action

Open a web page or run an iOS Shortcut when someone taps the button. Supported links:

- **HTTP/HTTPS:** Web pages, e.g. `https://example.com`
- **Shortcuts:** Run Jarvis with `shortcuts://run-shortcut?name=Jarvis` <!-- full-width -->

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
			URL:   "https://status.example.com/servers/prod-web-1",
		},
	},
)
```

#### Apple Shortcut action

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

#### Secondary action

![Alert Live Activity with primary and secondary action buttons](https://cdn.activitysmith.com/features/live-activity-secondary-action.png)

Use `secondary_action` when you want a second button beside the primary `action`.

The secondary action button is supported for `alert`, `progress`, and `segmented_progress` Live Activities. Both buttons use the same `open_url`, `webhook`, and Apple Shortcut payload shapes.

```go
activitysmith.LiveActivities.Stream(
	"agent-approval",
	activitysmithsdk.LiveActivityStreamInput{
		Title:   "Approval Needed",
		Message: "Should I send the follow-up email to Brightlane?",
		Type:    "alert",
		Color:   "green",
		Icon:    activitysmithsdk.AlertIcon("sparkles", "green"),
		Badge:   activitysmithsdk.AlertBadge("Agent", "green"),
		Action: &activitysmithsdk.LiveActivityActionInput{
			Title:  "Send",
			Type:   "webhook",
			URL:    "https://agent.example.com/live-activity/approve",
			Method: "POST",
			Body: map[string]interface{}{
				"approval_id": "approval_01JY3J7Q9S0P8M1V5PZK7DR4M2",
				"decision":    "send",
			},
		},
		SecondaryAction: &activitysmithsdk.LiveActivityActionInput{
			Title:  "Deny",
			Type:   "webhook",
			URL:    "https://agent.example.com/live-activity/deny",
			Method: "POST",
			Body: map[string]interface{}{
				"approval_id": "approval_01JY3J7Q9S0P8M1V5PZK7DR4M2",
				"decision":    "deny",
			},
		},
	},
)
```

## Lock Screen Widgets

![Lock screen widgets](https://cdn.activitysmith.com/features/lock-screen-widgets.png)

ActivitySmith lets you display any value on your Lock Screen with widgets - SaaS metrics, revenue, signups, uptime, habits, or anything else you want to track. Create a metric in the [web app](https://activitysmith.com/app/widgets), then update the metric value using our API, add a widget to your lock screen and it will fetch the latest update automatically.

![Create widget metric](https://cdn.activitysmith.com/features/create-widget-metric.png)

Use the metric key to update its value.

```go
activitysmith.Metrics.Update("deploy.success_rate", 99.9)
```

String metric values work too.

```go
activitysmith.Metrics.Update("prod.status", "healthy")
```

## App Icon Badge Count

![ActivitySmith app icon with an App Icon Badge Count](https://cdn.activitysmith.com/features/badge-count.png)

Show the number you care about on your ActivitySmith app icon. Track MRR, a customer count, a stock price, or any other value you want to keep in view.

### Set or update the badge value

```go
activitysmith.BadgeCount(8333)
```

### Clear the badge

Pass `0` to clear the badge.

```go
activitysmith.BadgeCount(0)
```

## Metadata

Metadata adds information to Push Notification and Live Activity details in ActivitySmith. It does not appear in the notification or Live Activity on your device.

```go
activitysmith.Notifications.Send(activitysmithsdk.PushNotificationInput{
    Title: "New subscription 💸",
    Metadata: map[string]any{
        "customer_id": "382", "plan": "Pro", "amount": 29, "trial": false,
    },
})

activitysmith.LiveActivities.Stream("customer-import", activitysmithsdk.LiveActivityStreamInput{
    ContentState: activitysmithsdk.LiveActivityContentStateInput{
        Title: "Customer Import", Type: "progress", Percentage: 60,
    },
    Metadata: map[string]any{"job_id": "import-382", "records": 1200},
})
```

Supported on Push Notifications, Live Activity streams, and legacy `Start`, `Update`, and `End` calls. On updates or end calls, leave `Metadata` nil to keep it, supply an object to replace it, or send `Metadata: map[string]any{}` to clear it.

Values can be strings, numbers, or booleans. Metadata supports up to 50 entries and 16 KB of JSON, with keys up to 100 characters and strings up to 4,000 characters. Nested objects, arrays, and null values are not supported.

## Tags

Use Tags to organize and filter Push Notification and Live Activity history. Tags are created automatically when you first use them. Sending Tags requires SDK version 1.10.0 or later.

```go
activitysmith.Notifications.Send(activitysmithsdk.PushNotificationInput{
    Title: "New subscription 💸",
    Message: "Customer upgraded to Pro plan",
    Tags: []string{"user:382", "billing"},
})
```

On Live Activity stream updates and legacy `Update` or `End` calls, leave `Tags` nil to keep existing Tags, supply a slice to replace them, or pass `Tags: []string{}` to clear them.

```go
activitysmith.LiveActivities.Update(activitysmithsdk.LiveActivityUpdateInput{
    ActivityID: "YOUR_ACTIVITY_ID",
    ContentState: activitysmithsdk.LiveActivityContentStateInput{
        Title: "Customer Import",
        Percentage: 60,
    },
    Tags: []string{},
})
```

`EndStream` also accepts final Tags and Metadata. Omit them to preserve existing values, or supply empty collections to clear them.

## Channels

Use `Channels` to target specific team members or devices when sending Push Notifications, Live Activities, or App Icon Badge Count updates. Omit it for account-wide delivery.

```go
activitysmith.Notifications.Send(activitysmithsdk.PushNotificationInput{
	Title:    "New subscription 💸",
	Message:  "Customer upgraded to Pro plan",
	Channels: []string{"sales", "customer-success"},
})

activitysmith.LiveActivities.Stream(
	"nightly-backup",
	activitysmithsdk.LiveActivityStreamInput{
		Title:         "Nightly database backup",
		NumberOfSteps: 3,
		CurrentStep:   1,
		Type:          "segmented_progress",
		Channels:      []string{"ios-builds"},
	},
)

activitysmith.BadgeCount(3, "sales", "customer-success")
```

## Error Handling

SDK calls return `response, err`, so check `err` after every call.
