# LiveActivityStreamPutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Operation** | **string** |  | 
**StreamKey** | **string** |  | 
**ActivityId** | Pointer to **NullableString** |  | [optional] 
**PreviousActivityId** | Pointer to **string** |  | [optional] 
**DevicesNotified** | Pointer to **int32** |  | [optional] 
**DevicesQueued** | Pointer to **int32** |  | [optional] 
**UsersNotified** | Pointer to **int32** |  | [optional] 
**EffectiveChannelSlugs** | Pointer to **[]string** |  | [optional] 
**Timestamp** | **time.Time** |  | 

## Methods

### NewLiveActivityStreamPutResponse

`func NewLiveActivityStreamPutResponse(success bool, operation string, streamKey string, timestamp time.Time, ) *LiveActivityStreamPutResponse`

NewLiveActivityStreamPutResponse instantiates a new LiveActivityStreamPutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityStreamPutResponseWithDefaults

`func NewLiveActivityStreamPutResponseWithDefaults() *LiveActivityStreamPutResponse`

NewLiveActivityStreamPutResponseWithDefaults instantiates a new LiveActivityStreamPutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *LiveActivityStreamPutResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *LiveActivityStreamPutResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *LiveActivityStreamPutResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetOperation

`func (o *LiveActivityStreamPutResponse) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *LiveActivityStreamPutResponse) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *LiveActivityStreamPutResponse) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetStreamKey

`func (o *LiveActivityStreamPutResponse) GetStreamKey() string`

GetStreamKey returns the StreamKey field if non-nil, zero value otherwise.

### GetStreamKeyOk

`func (o *LiveActivityStreamPutResponse) GetStreamKeyOk() (*string, bool)`

GetStreamKeyOk returns a tuple with the StreamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamKey

`func (o *LiveActivityStreamPutResponse) SetStreamKey(v string)`

SetStreamKey sets StreamKey field to given value.


### GetActivityId

`func (o *LiveActivityStreamPutResponse) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *LiveActivityStreamPutResponse) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *LiveActivityStreamPutResponse) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *LiveActivityStreamPutResponse) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *LiveActivityStreamPutResponse) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *LiveActivityStreamPutResponse) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetPreviousActivityId

`func (o *LiveActivityStreamPutResponse) GetPreviousActivityId() string`

GetPreviousActivityId returns the PreviousActivityId field if non-nil, zero value otherwise.

### GetPreviousActivityIdOk

`func (o *LiveActivityStreamPutResponse) GetPreviousActivityIdOk() (*string, bool)`

GetPreviousActivityIdOk returns a tuple with the PreviousActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousActivityId

`func (o *LiveActivityStreamPutResponse) SetPreviousActivityId(v string)`

SetPreviousActivityId sets PreviousActivityId field to given value.

### HasPreviousActivityId

`func (o *LiveActivityStreamPutResponse) HasPreviousActivityId() bool`

HasPreviousActivityId returns a boolean if a field has been set.

### GetDevicesNotified

`func (o *LiveActivityStreamPutResponse) GetDevicesNotified() int32`

GetDevicesNotified returns the DevicesNotified field if non-nil, zero value otherwise.

### GetDevicesNotifiedOk

`func (o *LiveActivityStreamPutResponse) GetDevicesNotifiedOk() (*int32, bool)`

GetDevicesNotifiedOk returns a tuple with the DevicesNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesNotified

`func (o *LiveActivityStreamPutResponse) SetDevicesNotified(v int32)`

SetDevicesNotified sets DevicesNotified field to given value.

### HasDevicesNotified

`func (o *LiveActivityStreamPutResponse) HasDevicesNotified() bool`

HasDevicesNotified returns a boolean if a field has been set.

### GetDevicesQueued

`func (o *LiveActivityStreamPutResponse) GetDevicesQueued() int32`

GetDevicesQueued returns the DevicesQueued field if non-nil, zero value otherwise.

### GetDevicesQueuedOk

`func (o *LiveActivityStreamPutResponse) GetDevicesQueuedOk() (*int32, bool)`

GetDevicesQueuedOk returns a tuple with the DevicesQueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesQueued

`func (o *LiveActivityStreamPutResponse) SetDevicesQueued(v int32)`

SetDevicesQueued sets DevicesQueued field to given value.

### HasDevicesQueued

`func (o *LiveActivityStreamPutResponse) HasDevicesQueued() bool`

HasDevicesQueued returns a boolean if a field has been set.

### GetUsersNotified

`func (o *LiveActivityStreamPutResponse) GetUsersNotified() int32`

GetUsersNotified returns the UsersNotified field if non-nil, zero value otherwise.

### GetUsersNotifiedOk

`func (o *LiveActivityStreamPutResponse) GetUsersNotifiedOk() (*int32, bool)`

GetUsersNotifiedOk returns a tuple with the UsersNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersNotified

`func (o *LiveActivityStreamPutResponse) SetUsersNotified(v int32)`

SetUsersNotified sets UsersNotified field to given value.

### HasUsersNotified

`func (o *LiveActivityStreamPutResponse) HasUsersNotified() bool`

HasUsersNotified returns a boolean if a field has been set.

### GetEffectiveChannelSlugs

`func (o *LiveActivityStreamPutResponse) GetEffectiveChannelSlugs() []string`

GetEffectiveChannelSlugs returns the EffectiveChannelSlugs field if non-nil, zero value otherwise.

### GetEffectiveChannelSlugsOk

`func (o *LiveActivityStreamPutResponse) GetEffectiveChannelSlugsOk() (*[]string, bool)`

GetEffectiveChannelSlugsOk returns a tuple with the EffectiveChannelSlugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveChannelSlugs

`func (o *LiveActivityStreamPutResponse) SetEffectiveChannelSlugs(v []string)`

SetEffectiveChannelSlugs sets EffectiveChannelSlugs field to given value.

### HasEffectiveChannelSlugs

`func (o *LiveActivityStreamPutResponse) HasEffectiveChannelSlugs() bool`

HasEffectiveChannelSlugs returns a boolean if a field has been set.

### GetTimestamp

`func (o *LiveActivityStreamPutResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LiveActivityStreamPutResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LiveActivityStreamPutResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


