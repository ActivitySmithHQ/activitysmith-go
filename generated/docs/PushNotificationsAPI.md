# \PushNotificationsAPI

All URIs are relative to *https://activitysmith.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendPushNotification**](PushNotificationsAPI.md#SendPushNotification) | **Post** /push-notification | Send a push notification



## SendPushNotification

> PushNotificationResponse SendPushNotification(ctx).PushNotificationRequest(pushNotificationRequest).Execute()

Send a push notification



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pushNotificationRequest := *openapiclient.NewPushNotificationRequest("Title_example") // PushNotificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PushNotificationsAPI.SendPushNotification(context.Background()).PushNotificationRequest(pushNotificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PushNotificationsAPI.SendPushNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendPushNotification`: PushNotificationResponse
	fmt.Fprintf(os.Stdout, "Response from `PushNotificationsAPI.SendPushNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendPushNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushNotificationRequest** | [**PushNotificationRequest**](PushNotificationRequest.md) |  | 

### Return type

[**PushNotificationResponse**](PushNotificationResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

