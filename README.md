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

	activitysmithsdk "github.com/ActivitySmithHQ/activitysmith-go"
)

func main() {
	activitysmith, err := activitysmithsdk.New("YOUR_API_KEY")
	if err != nil {
		log.Fatal(err)
	}

	input := activitysmithsdk.PushNotificationInput{
		Title:   "New subscription 💸",
		Message: "Customer upgraded to Pro plan",
	}

	response, err := activitysmith.Notifications.
		Send(input)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Notified:", response.GetDevicesNotified())
}
```

### Send a Push Notification

Use `activitysmith.Notifications.Send` with an `activitysmithsdk.PushNotificationInput`.

```go
input := activitysmithsdk.PushNotificationInput{
	Title:   "New subscription 💸",
	Message: "Customer upgraded to Pro plan",
	Channels: []string{"devs", "ops"}, // Optional
}

response, err := activitysmith.Notifications.
	Send(input)
if err != nil {
	log.Fatal(err)
}

log.Println(response.GetSuccess())
log.Println(response.GetDevicesNotified())
```

### Start a Live Activity

Use `activitysmith.LiveActivities.Start` with an `activitysmithsdk.LiveActivityStartInput`.

```go
startInput := activitysmithsdk.LiveActivityStartInput{
	Title:         "Nightly database backup",
	Subtitle:      "create snapshot",
	NumberOfSteps: 3,
	CurrentStep:   1,
	Type:          "segmented_progress",
	Color:         "yellow",
	Channels:      []string{"devs", "ops"}, // Optional
}

start, err := activitysmith.LiveActivities.
	Start(startInput)
if err != nil {
	log.Fatal(err)
}

activityID := start.GetActivityId()
```

### Update a Live Activity

Use `activitysmith.LiveActivities.Update` with the `activityID` from `Start`.

```go
updateInput := activitysmithsdk.LiveActivityUpdateInput{
	ActivityID:  activityID,
	Title:       "Nightly database backup",
	Subtitle:    "upload archive",
	CurrentStep: 2,
}

update, err := activitysmith.LiveActivities.
	Update(updateInput)
if err != nil {
	log.Fatal(err)
}

log.Println(update.GetDevicesNotified())
```

### End a Live Activity

Use `activitysmith.LiveActivities.End` to end the activity.
If `AutoDismissMinutes` is omitted, backend default `3` is used.

```go
endInput := activitysmithsdk.LiveActivityEndInput{
	ActivityID:         activityID,
	Title:              "Nightly database backup",
	Subtitle:           "verify restore",
	CurrentStep: 3,
	AutoDismissMinutes: 2,
}

end, err := activitysmith.LiveActivities.
	End(endInput)
if err != nil {
	log.Fatal(err)
}

log.Println(end.GetSuccess())
```

## Error Handling

SDK methods return `(response, error)`. Always check `error` on each call.

## API Surface

- `activitysmith.Notifications`
- `activitysmith.LiveActivities`

## Requirements

- Go 1.22+

## License

MIT
