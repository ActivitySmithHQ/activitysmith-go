# LiveActivityUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**ActivityId** | **string** |  | 
**DevicesQueued** | Pointer to **int32** |  | [optional] 
**DevicesNotified** | Pointer to **int32** |  | [optional] 
**Timestamp** | **time.Time** |  | 
**UsersNotified** | Pointer to **int32** |  | [optional] 
**EffectiveChannelSlugs** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** | Optional tags to organize and filter notification history. | [optional] 
**Operation** | Pointer to **string** | Present when the update started a new Live Activity for a managed stream. | [optional] 
**StreamKey** | Pointer to **string** |  | [optional] 
**PreviousActivityId** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** | Per-device delivery errors when some devices could not be reached. | [optional] 

## Methods

### NewLiveActivityUpdateResponse

`func NewLiveActivityUpdateResponse(success bool, activityId string, timestamp time.Time, ) *LiveActivityUpdateResponse`

NewLiveActivityUpdateResponse instantiates a new LiveActivityUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityUpdateResponseWithDefaults

`func NewLiveActivityUpdateResponseWithDefaults() *LiveActivityUpdateResponse`

NewLiveActivityUpdateResponseWithDefaults instantiates a new LiveActivityUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *LiveActivityUpdateResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *LiveActivityUpdateResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *LiveActivityUpdateResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetActivityId

`func (o *LiveActivityUpdateResponse) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *LiveActivityUpdateResponse) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *LiveActivityUpdateResponse) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.


### GetDevicesQueued

`func (o *LiveActivityUpdateResponse) GetDevicesQueued() int32`

GetDevicesQueued returns the DevicesQueued field if non-nil, zero value otherwise.

### GetDevicesQueuedOk

`func (o *LiveActivityUpdateResponse) GetDevicesQueuedOk() (*int32, bool)`

GetDevicesQueuedOk returns a tuple with the DevicesQueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesQueued

`func (o *LiveActivityUpdateResponse) SetDevicesQueued(v int32)`

SetDevicesQueued sets DevicesQueued field to given value.

### HasDevicesQueued

`func (o *LiveActivityUpdateResponse) HasDevicesQueued() bool`

HasDevicesQueued returns a boolean if a field has been set.

### GetDevicesNotified

`func (o *LiveActivityUpdateResponse) GetDevicesNotified() int32`

GetDevicesNotified returns the DevicesNotified field if non-nil, zero value otherwise.

### GetDevicesNotifiedOk

`func (o *LiveActivityUpdateResponse) GetDevicesNotifiedOk() (*int32, bool)`

GetDevicesNotifiedOk returns a tuple with the DevicesNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesNotified

`func (o *LiveActivityUpdateResponse) SetDevicesNotified(v int32)`

SetDevicesNotified sets DevicesNotified field to given value.

### HasDevicesNotified

`func (o *LiveActivityUpdateResponse) HasDevicesNotified() bool`

HasDevicesNotified returns a boolean if a field has been set.

### GetTimestamp

`func (o *LiveActivityUpdateResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LiveActivityUpdateResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LiveActivityUpdateResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetUsersNotified

`func (o *LiveActivityUpdateResponse) GetUsersNotified() int32`

GetUsersNotified returns the UsersNotified field if non-nil, zero value otherwise.

### GetUsersNotifiedOk

`func (o *LiveActivityUpdateResponse) GetUsersNotifiedOk() (*int32, bool)`

GetUsersNotifiedOk returns a tuple with the UsersNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersNotified

`func (o *LiveActivityUpdateResponse) SetUsersNotified(v int32)`

SetUsersNotified sets UsersNotified field to given value.

### HasUsersNotified

`func (o *LiveActivityUpdateResponse) HasUsersNotified() bool`

HasUsersNotified returns a boolean if a field has been set.

### GetEffectiveChannelSlugs

`func (o *LiveActivityUpdateResponse) GetEffectiveChannelSlugs() []string`

GetEffectiveChannelSlugs returns the EffectiveChannelSlugs field if non-nil, zero value otherwise.

### GetEffectiveChannelSlugsOk

`func (o *LiveActivityUpdateResponse) GetEffectiveChannelSlugsOk() (*[]string, bool)`

GetEffectiveChannelSlugsOk returns a tuple with the EffectiveChannelSlugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveChannelSlugs

`func (o *LiveActivityUpdateResponse) SetEffectiveChannelSlugs(v []string)`

SetEffectiveChannelSlugs sets EffectiveChannelSlugs field to given value.

### HasEffectiveChannelSlugs

`func (o *LiveActivityUpdateResponse) HasEffectiveChannelSlugs() bool`

HasEffectiveChannelSlugs returns a boolean if a field has been set.

### GetTags

`func (o *LiveActivityUpdateResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LiveActivityUpdateResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LiveActivityUpdateResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LiveActivityUpdateResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperation

`func (o *LiveActivityUpdateResponse) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *LiveActivityUpdateResponse) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *LiveActivityUpdateResponse) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *LiveActivityUpdateResponse) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetStreamKey

`func (o *LiveActivityUpdateResponse) GetStreamKey() string`

GetStreamKey returns the StreamKey field if non-nil, zero value otherwise.

### GetStreamKeyOk

`func (o *LiveActivityUpdateResponse) GetStreamKeyOk() (*string, bool)`

GetStreamKeyOk returns a tuple with the StreamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamKey

`func (o *LiveActivityUpdateResponse) SetStreamKey(v string)`

SetStreamKey sets StreamKey field to given value.

### HasStreamKey

`func (o *LiveActivityUpdateResponse) HasStreamKey() bool`

HasStreamKey returns a boolean if a field has been set.

### GetPreviousActivityId

`func (o *LiveActivityUpdateResponse) GetPreviousActivityId() string`

GetPreviousActivityId returns the PreviousActivityId field if non-nil, zero value otherwise.

### GetPreviousActivityIdOk

`func (o *LiveActivityUpdateResponse) GetPreviousActivityIdOk() (*string, bool)`

GetPreviousActivityIdOk returns a tuple with the PreviousActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousActivityId

`func (o *LiveActivityUpdateResponse) SetPreviousActivityId(v string)`

SetPreviousActivityId sets PreviousActivityId field to given value.

### HasPreviousActivityId

`func (o *LiveActivityUpdateResponse) HasPreviousActivityId() bool`

HasPreviousActivityId returns a boolean if a field has been set.

### GetErrors

`func (o *LiveActivityUpdateResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LiveActivityUpdateResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LiveActivityUpdateResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *LiveActivityUpdateResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


