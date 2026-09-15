# \PublicAPI

All URIs are relative to *https://activitysmith.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiHealth**](PublicAPI.md#GetApiHealth) | **Get** /health | Get API Health
[**ListPublicChangelogEntries**](PublicAPI.md#ListPublicChangelogEntries) | **Get** /changelog | List Public Changelog Entries



## GetApiHealth

> HealthResponse GetApiHealth(ctx).Execute()

Get API Health



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetApiHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetApiHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiHealth`: HealthResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetApiHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiHealthRequest struct via the builder pattern


### Return type

[**HealthResponse**](HealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPublicChangelogEntries

> ChangelogListResponse ListPublicChangelogEntries(ctx).Platform(platform).Limit(limit).Execute()

List Public Changelog Entries



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
	platform := "platform_example" // string | Platform whose published changelog entries should be returned. (optional) (default to "ios")
	limit := int32(56) // int32 | Maximum number of entries to return. (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.ListPublicChangelogEntries(context.Background()).Platform(platform).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.ListPublicChangelogEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPublicChangelogEntries`: ChangelogListResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.ListPublicChangelogEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPublicChangelogEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platform** | **string** | Platform whose published changelog entries should be returned. | [default to &quot;ios&quot;]
 **limit** | **int32** | Maximum number of entries to return. | [default to 10]

### Return type

[**ChangelogListResponse**](ChangelogListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

