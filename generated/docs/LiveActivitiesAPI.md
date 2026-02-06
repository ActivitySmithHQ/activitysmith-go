# \LiveActivitiesAPI

All URIs are relative to *https://activitysmith.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndLiveActivity**](LiveActivitiesAPI.md#EndLiveActivity) | **Post** /live-activity/end | End a Live Activity
[**StartLiveActivity**](LiveActivitiesAPI.md#StartLiveActivity) | **Post** /live-activity/start | Start a Live Activity
[**UpdateLiveActivity**](LiveActivitiesAPI.md#UpdateLiveActivity) | **Post** /live-activity/update | Update a Live Activity



## EndLiveActivity

> LiveActivityEndResponse EndLiveActivity(ctx).LiveActivityEndRequest(liveActivityEndRequest).Execute()

End a Live Activity



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
	liveActivityEndRequest := *openapiclient.NewLiveActivityEndRequest("ActivityId_example", *openapiclient.NewContentStateEnd("Title_example", int32(123))) // LiveActivityEndRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveActivitiesAPI.EndLiveActivity(context.Background()).LiveActivityEndRequest(liveActivityEndRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveActivitiesAPI.EndLiveActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndLiveActivity`: LiveActivityEndResponse
	fmt.Fprintf(os.Stdout, "Response from `LiveActivitiesAPI.EndLiveActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndLiveActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **liveActivityEndRequest** | [**LiveActivityEndRequest**](LiveActivityEndRequest.md) |  | 

### Return type

[**LiveActivityEndResponse**](LiveActivityEndResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartLiveActivity

> LiveActivityStartResponse StartLiveActivity(ctx).LiveActivityStartRequest(liveActivityStartRequest).Execute()

Start a Live Activity



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
	liveActivityStartRequest := *openapiclient.NewLiveActivityStartRequest(*openapiclient.NewContentStateStart("Title_example", int32(123), int32(123), "Type_example")) // LiveActivityStartRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveActivitiesAPI.StartLiveActivity(context.Background()).LiveActivityStartRequest(liveActivityStartRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveActivitiesAPI.StartLiveActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartLiveActivity`: LiveActivityStartResponse
	fmt.Fprintf(os.Stdout, "Response from `LiveActivitiesAPI.StartLiveActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartLiveActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **liveActivityStartRequest** | [**LiveActivityStartRequest**](LiveActivityStartRequest.md) |  | 

### Return type

[**LiveActivityStartResponse**](LiveActivityStartResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLiveActivity

> LiveActivityUpdateResponse UpdateLiveActivity(ctx).LiveActivityUpdateRequest(liveActivityUpdateRequest).Execute()

Update a Live Activity



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
	liveActivityUpdateRequest := *openapiclient.NewLiveActivityUpdateRequest("ActivityId_example", *openapiclient.NewContentStateUpdate("Title_example", int32(123))) // LiveActivityUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveActivitiesAPI.UpdateLiveActivity(context.Background()).LiveActivityUpdateRequest(liveActivityUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveActivitiesAPI.UpdateLiveActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLiveActivity`: LiveActivityUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `LiveActivitiesAPI.UpdateLiveActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLiveActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **liveActivityUpdateRequest** | [**LiveActivityUpdateRequest**](LiveActivityUpdateRequest.md) |  | 

### Return type

[**LiveActivityUpdateResponse**](LiveActivityUpdateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

