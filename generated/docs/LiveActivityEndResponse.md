# LiveActivityEndResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**ActivityId** | **string** |  | 
**DevicesQueued** | Pointer to **int32** |  | [optional] 
**DevicesNotified** | Pointer to **int32** |  | [optional] 
**Timestamp** | **time.Time** |  | 

## Methods

### NewLiveActivityEndResponse

`func NewLiveActivityEndResponse(success bool, activityId string, timestamp time.Time, ) *LiveActivityEndResponse`

NewLiveActivityEndResponse instantiates a new LiveActivityEndResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityEndResponseWithDefaults

`func NewLiveActivityEndResponseWithDefaults() *LiveActivityEndResponse`

NewLiveActivityEndResponseWithDefaults instantiates a new LiveActivityEndResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *LiveActivityEndResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *LiveActivityEndResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *LiveActivityEndResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetActivityId

`func (o *LiveActivityEndResponse) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *LiveActivityEndResponse) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *LiveActivityEndResponse) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.


### GetDevicesQueued

`func (o *LiveActivityEndResponse) GetDevicesQueued() int32`

GetDevicesQueued returns the DevicesQueued field if non-nil, zero value otherwise.

### GetDevicesQueuedOk

`func (o *LiveActivityEndResponse) GetDevicesQueuedOk() (*int32, bool)`

GetDevicesQueuedOk returns a tuple with the DevicesQueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesQueued

`func (o *LiveActivityEndResponse) SetDevicesQueued(v int32)`

SetDevicesQueued sets DevicesQueued field to given value.

### HasDevicesQueued

`func (o *LiveActivityEndResponse) HasDevicesQueued() bool`

HasDevicesQueued returns a boolean if a field has been set.

### GetDevicesNotified

`func (o *LiveActivityEndResponse) GetDevicesNotified() int32`

GetDevicesNotified returns the DevicesNotified field if non-nil, zero value otherwise.

### GetDevicesNotifiedOk

`func (o *LiveActivityEndResponse) GetDevicesNotifiedOk() (*int32, bool)`

GetDevicesNotifiedOk returns a tuple with the DevicesNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesNotified

`func (o *LiveActivityEndResponse) SetDevicesNotified(v int32)`

SetDevicesNotified sets DevicesNotified field to given value.

### HasDevicesNotified

`func (o *LiveActivityEndResponse) HasDevicesNotified() bool`

HasDevicesNotified returns a boolean if a field has been set.

### GetTimestamp

`func (o *LiveActivityEndResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LiveActivityEndResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LiveActivityEndResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


