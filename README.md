# ActivitySmith Go SDK

The ActivitySmith Go SDK provides convenient access to the ActivitySmith API from Go applications.

## Documentation

See [API reference](https://activitysmith.com/docs/api-reference/introduction).

## Installation

```sh
go get github.com/ActivitySmithHQ/activitysmith-go
```

## Usage

```go
package main

import (
	"log"

	activitysmith "github.com/ActivitySmithHQ/activitysmith-go"
	"github.com/ActivitySmithHQ/activitysmith-go/generated"
)

func main() {
	client, err := activitysmith.New("YOUR_API_KEY", nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Notifications.
		SendPushNotification(
			generated.PushNotificationRequest{
				Title:   "Build Failed",
				Message: generated.PtrString("CI pipeline failed on main branch"),
			},
		)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Notified:", response.GetDevicesNotified())
}
```

### Send a Push Notification

Use `client.Notifications.SendPushNotification` with a `generated.PushNotificationRequest`. `Title` is required; `Message` and `Subtitle` are optional.

```go
response, err := client.Notifications.
	SendPushNotification(
		generated.PushNotificationRequest{
			Title:   "Build Failed",
			Message: generated.PtrString("CI pipeline failed on main branch"),
		},
	)
if err != nil {
	log.Fatal(err)
}

log.Println(response.GetSuccess())
log.Println(response.GetDevicesNotified())
```

### Start a Live Activity

Use `client.LiveActivities.StartLiveActivity` with a `generated.LiveActivityStartRequest`.

```go
start, err := client.LiveActivities.
	StartLiveActivity(
		generated.LiveActivityStartRequest{
			ContentState: generated.ContentStateStart{
				Title:         "ActivitySmith API Deployment",
				Subtitle:      generated.PtrString("start"),
				NumberOfSteps: 4,
				CurrentStep:   1,
				Type:          "segmented_progress",
				Color:         generated.PtrString("yellow"),
			},
		},
	)
if err != nil {
	log.Fatal(err)
}

activityID := start.GetActivityId()
```

### Update a Live Activity

Use `client.LiveActivities.UpdateLiveActivity` with the `activityID` from `StartLiveActivity`.

```go
update, err := client.LiveActivities.
	UpdateLiveActivity(
		generated.LiveActivityUpdateRequest{
			ActivityId: activityID,
			ContentState: generated.ContentStateUpdate{
				Title:       "ActivitySmith API Deployment",
				Subtitle:    generated.PtrString("npm i & pm2"),
				CurrentStep: 3,
			},
		},
	)
if err != nil {
	log.Fatal(err)
}

log.Println(update.GetDevicesNotified())
```

### End a Live Activity

Use `client.LiveActivities.EndLiveActivity` to end the activity and optionally control dismissal time.

```go
end, err := client.LiveActivities.
	EndLiveActivity(
		generated.LiveActivityEndRequest{
			ActivityId: activityID,
			ContentState: generated.ContentStateEnd{
				Title:              "ActivitySmith API Deployment",
				Subtitle:           generated.PtrString("done"),
				CurrentStep:        4,
				AutoDismissMinutes: generated.PtrInt32(3),
			},
		},
	)
if err != nil {
	log.Fatal(err)
}

log.Println(end.GetSuccess())
```

## Error Handling

SDK methods return `(response, error)`. Always check `error` on each call.

## API Surface

- `client.Notifications`
- `client.LiveActivities`

## Requirements

- Go 1.22+

## License

MIT
