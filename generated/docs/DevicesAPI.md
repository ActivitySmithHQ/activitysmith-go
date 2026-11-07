# \DevicesAPI

All URIs are relative to *https://activitysmith.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccountDevices**](DevicesAPI.md#ListAccountDevices) | **Get** /account/devices | List paired devices



## ListAccountDevices

> AccountDeviceListResponse ListAccountDevices(ctx).Limit(limit).Offset(offset).IncludeInactive(includeInactive).Execute()

List paired devices



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
	limit := int32(56) // int32 | Maximum devices to return, 1 to 100. Defaults to 100. (optional) (default to 100)
	offset := int32(56) // int32 | Devices to skip. Use meta.next_offset from the previous page. (optional) (default to 0)
	includeInactive := true // bool | Include devices that can no longer receive deliveries, such as unpaired phones. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.ListAccountDevices(context.Background()).Limit(limit).Offset(offset).IncludeInactive(includeInactive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.ListAccountDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccountDevices`: AccountDeviceListResponse
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.ListAccountDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum devices to return, 1 to 100. Defaults to 100. | [default to 100]
 **offset** | **int32** | Devices to skip. Use meta.next_offset from the previous page. | [default to 0]
 **includeInactive** | **bool** | Include devices that can no longer receive deliveries, such as unpaired phones. | [default to false]

### Return type

[**AccountDeviceListResponse**](AccountDeviceListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

