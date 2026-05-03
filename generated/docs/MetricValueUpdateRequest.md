# MetricValueUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**MetricValueUpdateRequestValue**](MetricValueUpdateRequestValue.md) |  | 
**Timestamp** | Pointer to **time.Time** | Optional ISO timestamp for when the metric value was measured. Defaults to the server receive time. | [optional] 

## Methods

### NewMetricValueUpdateRequest

`func NewMetricValueUpdateRequest(value MetricValueUpdateRequestValue, ) *MetricValueUpdateRequest`

NewMetricValueUpdateRequest instantiates a new MetricValueUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricValueUpdateRequestWithDefaults

`func NewMetricValueUpdateRequestWithDefaults() *MetricValueUpdateRequest`

NewMetricValueUpdateRequestWithDefaults instantiates a new MetricValueUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *MetricValueUpdateRequest) GetValue() MetricValueUpdateRequestValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricValueUpdateRequest) GetValueOk() (*MetricValueUpdateRequestValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricValueUpdateRequest) SetValue(v MetricValueUpdateRequestValue)`

SetValue sets Value field to given value.


### GetTimestamp

`func (o *MetricValueUpdateRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MetricValueUpdateRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MetricValueUpdateRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MetricValueUpdateRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


