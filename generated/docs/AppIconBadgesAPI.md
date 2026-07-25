# \AppIconBadgesAPI

All URIs are relative to *https://activitysmith.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateAppIconBadgeCount**](AppIconBadgesAPI.md#UpdateAppIconBadgeCount) | **Post** /badge | Set App Icon Badge Count



## UpdateAppIconBadgeCount

> AppIconBadgeCountUpdateResponse UpdateAppIconBadgeCount(ctx).AppIconBadgeCountUpdateRequest(appIconBadgeCountUpdateRequest).Execute()

Set App Icon Badge Count



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
	appIconBadgeCountUpdateRequest := *openapiclient.NewAppIconBadgeCountUpdateRequest(int32(123)) // AppIconBadgeCountUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppIconBadgesAPI.UpdateAppIconBadgeCount(context.Background()).AppIconBadgeCountUpdateRequest(appIconBadgeCountUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppIconBadgesAPI.UpdateAppIconBadgeCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppIconBadgeCount`: AppIconBadgeCountUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `AppIconBadgesAPI.UpdateAppIconBadgeCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppIconBadgeCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appIconBadgeCountUpdateRequest** | [**AppIconBadgeCountUpdateRequest**](AppIconBadgeCountUpdateRequest.md) |  | 

### Return type

[**AppIconBadgeCountUpdateResponse**](AppIconBadgeCountUpdateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

