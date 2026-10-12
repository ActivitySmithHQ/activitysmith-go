# LiveActivityEndResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**ActivityId** | **string** |  | 
**DevicesQueued** | Pointer to **int32** |  | [optional] 
**DevicesNotified** | Pointer to **int32** |  | [optional] 
**Timestamp** | **time.Time** |  | 
**Ignored** | Pointer to **bool** | True when the Live Activity had already ended and the request had no effect. | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**EndReason** | Pointer to **NullableString** |  | [optional] 
**ForcedCompletion** | Pointer to **bool** | True when the Live Activity was marked complete without a confirmed APNs end delivery. | [optional] 
**Warning** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** | Per-device delivery errors when some devices could not be reached. | [optional] 

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


### GetIgnored

`func (o *LiveActivityEndResponse) GetIgnored() bool`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *LiveActivityEndResponse) GetIgnoredOk() (*bool, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *LiveActivityEndResponse) SetIgnored(v bool)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *LiveActivityEndResponse) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### GetMessage

`func (o *LiveActivityEndResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LiveActivityEndResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LiveActivityEndResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LiveActivityEndResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEndReason

`func (o *LiveActivityEndResponse) GetEndReason() string`

GetEndReason returns the EndReason field if non-nil, zero value otherwise.

### GetEndReasonOk

`func (o *LiveActivityEndResponse) GetEndReasonOk() (*string, bool)`

GetEndReasonOk returns a tuple with the EndReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndReason

`func (o *LiveActivityEndResponse) SetEndReason(v string)`

SetEndReason sets EndReason field to given value.

### HasEndReason

`func (o *LiveActivityEndResponse) HasEndReason() bool`

HasEndReason returns a boolean if a field has been set.

### SetEndReasonNil

`func (o *LiveActivityEndResponse) SetEndReasonNil(b bool)`

 SetEndReasonNil sets the value for EndReason to be an explicit nil

### UnsetEndReason
`func (o *LiveActivityEndResponse) UnsetEndReason()`

UnsetEndReason ensures that no value is present for EndReason, not even an explicit nil
### GetForcedCompletion

`func (o *LiveActivityEndResponse) GetForcedCompletion() bool`

GetForcedCompletion returns the ForcedCompletion field if non-nil, zero value otherwise.

### GetForcedCompletionOk

`func (o *LiveActivityEndResponse) GetForcedCompletionOk() (*bool, bool)`

GetForcedCompletionOk returns a tuple with the ForcedCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedCompletion

`func (o *LiveActivityEndResponse) SetForcedCompletion(v bool)`

SetForcedCompletion sets ForcedCompletion field to given value.

### HasForcedCompletion

`func (o *LiveActivityEndResponse) HasForcedCompletion() bool`

HasForcedCompletion returns a boolean if a field has been set.

### GetWarning

`func (o *LiveActivityEndResponse) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *LiveActivityEndResponse) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *LiveActivityEndResponse) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *LiveActivityEndResponse) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetErrors

`func (o *LiveActivityEndResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LiveActivityEndResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LiveActivityEndResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *LiveActivityEndResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


