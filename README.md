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
	"github.com/ActivitySmithHQ/activitysmith-go/generated"
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

<p align="center">
  <img src="https://cdn.activitysmith.com/features/new-subscription-push-notification.png" alt="Push notification example" width="680" />
</p>

Use `activitysmith.Notifications.Send` with either `activitysmithsdk.PushNotificationInput` (basic) or `generated.PushNotificationRequest` (advanced fields like `redirection` and `actions`).

```go
input := activitysmithsdk.PushNotificationInput{
	Title:   "New subscription 💸",
	Message: "Customer upgraded to Pro plan",
}

response, err := activitysmith.Notifications.Send(input)
if err != nil {
	log.Fatal(err)
}

log.Println(response.GetSuccess())
log.Println(response.GetDevicesNotified())
```

## Live Activities

Live Activities come in two UI types, but the lifecycle stays the same:
start the activity, keep the returned `activityID`, update it as state changes,
then end it when the work is done.

- `segmented_progress`: best for jobs tracked in steps
- `progress`: best for jobs tracked as a percentage or numeric range

### Shared flow

1. Call `activitysmith.LiveActivities.Start(...)`.
2. Save the returned `activityID`.
3. Call `activitysmith.LiveActivities.Update(...)` as progress changes.
4. Call `activitysmith.LiveActivities.End(...)` when the work is finished.

### Segmented Progress Type

Use `segmented_progress` when progress is easier to follow as steps instead of a
raw percentage. It fits jobs like backups, deployments, ETL pipelines, and
checklists where "step 2 of 3" is more useful than "67%".
`NumberOfSteps` is dynamic, so you can increase or decrease it later if the
workflow changes.

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
	Channels:      []string{"devs", "ops"}, // Optional
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
	NumberOfSteps: 4,
	CurrentStep:   2,
}

update, err := activitysmith.LiveActivities.Update(updateInput)
if err != nil {
	log.Fatal(err)
}

log.Println(update.GetDevicesNotified())
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
	NumberOfSteps:      4,
	CurrentStep:        4,
	AutoDismissMinutes: 2,
}

end, err := activitysmith.LiveActivities.End(endInput)
if err != nil {
	log.Fatal(err)
}

log.Println(end.GetSuccess())
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
	Color:      "lime",
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
	ActivityID:  activityID,
	Title:       "EV Charging",
	Subtitle:    "Added 120 mi range",
	Percentage:  60,
}

update, err := activitysmith.LiveActivities.Update(updateInput)
if err != nil {
	log.Fatal(err)
}

log.Println(update.GetDevicesNotified())
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

end, err := activitysmith.LiveActivities.End(endInput)
if err != nil {
	log.Fatal(err)
}

log.Println(end.GetSuccess())
```

## Channels

Channels are used to target specific team members or devices. Can be used for both push notifications and live activities.

```go
request := generated.NewPushNotificationRequest("New subscription 💸")
request.SetMessage("Customer upgraded to Pro plan")
request.SetTarget(generated.ChannelTarget{Channels: []string{"sales", "customer-success"}}) // Optional

response, err := activitysmith.Notifications.Send(request)
if err != nil {
	log.Fatal(err)
}

log.Println(response.GetSuccess())
log.Println(response.GetDevicesNotified())
```

## Push Notification Redirection and Actions

Push notification redirection and actions are optional and can be used to redirect the user to a specific URL when they tap the notification or to trigger a specific action when they long-press the notification.
Webhooks are executed by ActivitySmith backend.

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

response, err := activitysmith.Notifications.Send(request)
if err != nil {
	log.Fatal(err)
}

log.Println(response.GetSuccess())
log.Println(response.GetDevicesNotified())
```

## Error Handling

SDK methods return `(response, error)`. Always check `error` on each call.

## Requirements

- Go 1.22+

## License

MIT
