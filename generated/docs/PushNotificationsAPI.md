# \PushNotificationsAPI

All URIs are relative to *https://activitysmith.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReplacePushNotification**](PushNotificationsAPI.md#ReplacePushNotification) | **Put** /push-notification/{notification_id} | Replace a push notification
[**SendPushNotification**](PushNotificationsAPI.md#SendPushNotification) | **Post** /push-notification | Send a push notification



## ReplacePushNotification

> PushNotificationReplaceResponse ReplacePushNotification(ctx, notificationId).PushNotificationReplaceRequest(pushNotificationReplaceRequest).Execute()

Replace a push notification



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
	notificationId := "notificationId_example" // string | The `notification_id` returned when the Push Notification was sent.
	pushNotificationReplaceRequest := *openapiclient.NewPushNotificationReplaceRequest("Title_example") // PushNotificationReplaceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PushNotificationsAPI.ReplacePushNotification(context.Background(), notificationId).PushNotificationReplaceRequest(pushNotificationReplaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PushNotificationsAPI.ReplacePushNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplacePushNotification`: PushNotificationReplaceResponse
	fmt.Fprintf(os.Stdout, "Response from `PushNotificationsAPI.ReplacePushNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | The &#x60;notification_id&#x60; returned when the Push Notification was sent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePushNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pushNotificationReplaceRequest** | [**PushNotificationReplaceRequest**](PushNotificationReplaceRequest.md) |  | 

### Return type

[**PushNotificationReplaceResponse**](PushNotificationReplaceResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

