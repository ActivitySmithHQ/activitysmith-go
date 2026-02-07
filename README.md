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
		Send(
			generated.PushNotificationRequest{
				Title:   "Build Failed",
			},
		)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Notified:", response.GetDevicesNotified())
}
```

### Send a Push Notification

Use `client.Notifications.Send` with a `generated.PushNotificationRequest`. This example uses only required fields.

```go
response, err := client.Notifications.
	Send(
		generated.PushNotificationRequest{
			Title:   "Build Failed",
		},
	)
if err != nil {
	log.Fatal(err)
}

log.Println(response.GetSuccess())
log.Println(response.GetDevicesNotified())
```

### Start a Live Activity

Use `client.LiveActivities.Start` with a `generated.LiveActivityStartRequest`.

```go
start, err := client.LiveActivities.
	Start(
		generated.LiveActivityStartRequest{
			ContentState: generated.ContentStateStart{
				Title:         "ActivitySmith API Deployment",
				NumberOfSteps: 4,
				CurrentStep:   1,
				Type:          "segmented_progress",
			},
		},
	)
if err != nil {
	log.Fatal(err)
}

activityID := start.GetActivityId()
```

### Update a Live Activity

Use `client.LiveActivities.Update` with the `activityID` from `Start`.

```go
update, err := client.LiveActivities.
	Update(
		generated.LiveActivityUpdateRequest{
			ActivityId: activityID,
			ContentState: generated.ContentStateUpdate{
				Title:       "ActivitySmith API Deployment",
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

Use `client.LiveActivities.End` to end the activity. This example uses only required fields.

```go
end, err := client.LiveActivities.
	End(
		generated.LiveActivityEndRequest{
			ActivityId: activityID,
			ContentState: generated.ContentStateEnd{
				Title:              "ActivitySmith API Deployment",
				CurrentStep:        4,
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
