# LiveActivityStartResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**DevicesNotified** | Pointer to **int32** |  | [optional] 
**UsersNotified** | Pointer to **int32** |  | [optional] 
**ActivityId** | **string** |  | 
**EffectiveChannelSlugs** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** | Optional tags to organize and filter notification history. | [optional] 
**Timestamp** | **time.Time** |  | 
**EndedActivityId** | Pointer to **string** | Present when starting this Live Activity ended an older one to stay within the device limit. | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** | Per-device delivery errors when some devices could not be reached. | [optional] 

## Methods

### NewLiveActivityStartResponse

`func NewLiveActivityStartResponse(success bool, activityId string, timestamp time.Time, ) *LiveActivityStartResponse`

NewLiveActivityStartResponse instantiates a new LiveActivityStartResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityStartResponseWithDefaults

`func NewLiveActivityStartResponseWithDefaults() *LiveActivityStartResponse`

NewLiveActivityStartResponseWithDefaults instantiates a new LiveActivityStartResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *LiveActivityStartResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *LiveActivityStartResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *LiveActivityStartResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetDevicesNotified

`func (o *LiveActivityStartResponse) GetDevicesNotified() int32`

GetDevicesNotified returns the DevicesNotified field if non-nil, zero value otherwise.

### GetDevicesNotifiedOk

`func (o *LiveActivityStartResponse) GetDevicesNotifiedOk() (*int32, bool)`

GetDevicesNotifiedOk returns a tuple with the DevicesNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesNotified

`func (o *LiveActivityStartResponse) SetDevicesNotified(v int32)`

SetDevicesNotified sets DevicesNotified field to given value.

### HasDevicesNotified

`func (o *LiveActivityStartResponse) HasDevicesNotified() bool`

HasDevicesNotified returns a boolean if a field has been set.

### GetUsersNotified

`func (o *LiveActivityStartResponse) GetUsersNotified() int32`

GetUsersNotified returns the UsersNotified field if non-nil, zero value otherwise.

### GetUsersNotifiedOk

`func (o *LiveActivityStartResponse) GetUsersNotifiedOk() (*int32, bool)`

GetUsersNotifiedOk returns a tuple with the UsersNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersNotified

`func (o *LiveActivityStartResponse) SetUsersNotified(v int32)`

SetUsersNotified sets UsersNotified field to given value.

### HasUsersNotified

`func (o *LiveActivityStartResponse) HasUsersNotified() bool`

HasUsersNotified returns a boolean if a field has been set.

### GetActivityId

`func (o *LiveActivityStartResponse) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *LiveActivityStartResponse) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *LiveActivityStartResponse) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.


### GetEffectiveChannelSlugs

`func (o *LiveActivityStartResponse) GetEffectiveChannelSlugs() []string`

GetEffectiveChannelSlugs returns the EffectiveChannelSlugs field if non-nil, zero value otherwise.

### GetEffectiveChannelSlugsOk

`func (o *LiveActivityStartResponse) GetEffectiveChannelSlugsOk() (*[]string, bool)`

GetEffectiveChannelSlugsOk returns a tuple with the EffectiveChannelSlugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveChannelSlugs

`func (o *LiveActivityStartResponse) SetEffectiveChannelSlugs(v []string)`

SetEffectiveChannelSlugs sets EffectiveChannelSlugs field to given value.

### HasEffectiveChannelSlugs

`func (o *LiveActivityStartResponse) HasEffectiveChannelSlugs() bool`

HasEffectiveChannelSlugs returns a boolean if a field has been set.

### GetTags

`func (o *LiveActivityStartResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LiveActivityStartResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LiveActivityStartResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LiveActivityStartResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *LiveActivityStartResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LiveActivityStartResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LiveActivityStartResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetEndedActivityId

`func (o *LiveActivityStartResponse) GetEndedActivityId() string`

GetEndedActivityId returns the EndedActivityId field if non-nil, zero value otherwise.

### GetEndedActivityIdOk

`func (o *LiveActivityStartResponse) GetEndedActivityIdOk() (*string, bool)`

GetEndedActivityIdOk returns a tuple with the EndedActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedActivityId

`func (o *LiveActivityStartResponse) SetEndedActivityId(v string)`

SetEndedActivityId sets EndedActivityId field to given value.

### HasEndedActivityId

`func (o *LiveActivityStartResponse) HasEndedActivityId() bool`

HasEndedActivityId returns a boolean if a field has been set.

### GetMessage

`func (o *LiveActivityStartResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LiveActivityStartResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LiveActivityStartResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LiveActivityStartResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *LiveActivityStartResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LiveActivityStartResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LiveActivityStartResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *LiveActivityStartResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


