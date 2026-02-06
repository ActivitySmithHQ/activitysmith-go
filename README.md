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

	_, err = client.Notifications.SendPushNotification(generated.PushNotificationRequest{})
	if err != nil {
		log.Fatal(err)
	}
}
```

## API Surface

- `client.Notifications`
- `client.LiveActivities`

## Requirements

- Go 1.22+

## License

MIT
