# \MetricsAPI

All URIs are relative to *https://activitysmith.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateMetricValue**](MetricsAPI.md#UpdateMetricValue) | **Post** /metrics/{key}/value | Update a widget metric value



## UpdateMetricValue

> MetricValueUpdateResponse UpdateMetricValue(ctx, key).MetricValueUpdateRequest(metricValueUpdateRequest).Execute()

Update a widget metric value



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
	key := "key_example" // string | Metric key configured in the web app. Lowercase letters, numbers, dots, underscores, and dashes are allowed.
	metricValueUpdateRequest := *openapiclient.NewMetricValueUpdateRequest(openapiclient.MetricValueUpdateRequest_value{Float32: new(float32)}) // MetricValueUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.UpdateMetricValue(context.Background(), key).MetricValueUpdateRequest(metricValueUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.UpdateMetricValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMetricValue`: MetricValueUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.UpdateMetricValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Metric key configured in the web app. Lowercase letters, numbers, dots, underscores, and dashes are allowed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetricValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricValueUpdateRequest** | [**MetricValueUpdateRequest**](MetricValueUpdateRequest.md) |  | 

### Return type

[**MetricValueUpdateResponse**](MetricValueUpdateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

