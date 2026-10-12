# LiveActivityStreamDeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Operation** | **string** |  | 
**StreamKey** | **string** |  | 
**ActivityId** | Pointer to **NullableString** |  | [optional] 
**DevicesQueued** | Pointer to **int32** |  | [optional] 
**DevicesNotified** | Pointer to **int32** |  | [optional] 
**Timestamp** | **time.Time** |  | 
**Ignored** | Pointer to **bool** | True when the stream was already ended and the request had no effect. | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ForcedCompletion** | Pointer to **bool** |  | [optional] 
**Warning** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** | Per-device delivery errors when some devices could not be reached. | [optional] 

## Methods

### NewLiveActivityStreamDeleteResponse

`func NewLiveActivityStreamDeleteResponse(success bool, operation string, streamKey string, timestamp time.Time, ) *LiveActivityStreamDeleteResponse`

NewLiveActivityStreamDeleteResponse instantiates a new LiveActivityStreamDeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityStreamDeleteResponseWithDefaults

`func NewLiveActivityStreamDeleteResponseWithDefaults() *LiveActivityStreamDeleteResponse`

NewLiveActivityStreamDeleteResponseWithDefaults instantiates a new LiveActivityStreamDeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *LiveActivityStreamDeleteResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *LiveActivityStreamDeleteResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *LiveActivityStreamDeleteResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetOperation

`func (o *LiveActivityStreamDeleteResponse) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *LiveActivityStreamDeleteResponse) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *LiveActivityStreamDeleteResponse) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetStreamKey

`func (o *LiveActivityStreamDeleteResponse) GetStreamKey() string`

GetStreamKey returns the StreamKey field if non-nil, zero value otherwise.

### GetStreamKeyOk

`func (o *LiveActivityStreamDeleteResponse) GetStreamKeyOk() (*string, bool)`

GetStreamKeyOk returns a tuple with the StreamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamKey

`func (o *LiveActivityStreamDeleteResponse) SetStreamKey(v string)`

SetStreamKey sets StreamKey field to given value.


### GetActivityId

`func (o *LiveActivityStreamDeleteResponse) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *LiveActivityStreamDeleteResponse) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *LiveActivityStreamDeleteResponse) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *LiveActivityStreamDeleteResponse) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *LiveActivityStreamDeleteResponse) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *LiveActivityStreamDeleteResponse) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetDevicesQueued

`func (o *LiveActivityStreamDeleteResponse) GetDevicesQueued() int32`

GetDevicesQueued returns the DevicesQueued field if non-nil, zero value otherwise.

### GetDevicesQueuedOk

`func (o *LiveActivityStreamDeleteResponse) GetDevicesQueuedOk() (*int32, bool)`

GetDevicesQueuedOk returns a tuple with the DevicesQueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesQueued

`func (o *LiveActivityStreamDeleteResponse) SetDevicesQueued(v int32)`

SetDevicesQueued sets DevicesQueued field to given value.

### HasDevicesQueued

`func (o *LiveActivityStreamDeleteResponse) HasDevicesQueued() bool`

HasDevicesQueued returns a boolean if a field has been set.

### GetDevicesNotified

`func (o *LiveActivityStreamDeleteResponse) GetDevicesNotified() int32`

GetDevicesNotified returns the DevicesNotified field if non-nil, zero value otherwise.

### GetDevicesNotifiedOk

`func (o *LiveActivityStreamDeleteResponse) GetDevicesNotifiedOk() (*int32, bool)`

GetDevicesNotifiedOk returns a tuple with the DevicesNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesNotified

`func (o *LiveActivityStreamDeleteResponse) SetDevicesNotified(v int32)`

SetDevicesNotified sets DevicesNotified field to given value.

### HasDevicesNotified

`func (o *LiveActivityStreamDeleteResponse) HasDevicesNotified() bool`

HasDevicesNotified returns a boolean if a field has been set.

### GetTimestamp

`func (o *LiveActivityStreamDeleteResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LiveActivityStreamDeleteResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LiveActivityStreamDeleteResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetIgnored

`func (o *LiveActivityStreamDeleteResponse) GetIgnored() bool`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *LiveActivityStreamDeleteResponse) GetIgnoredOk() (*bool, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *LiveActivityStreamDeleteResponse) SetIgnored(v bool)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *LiveActivityStreamDeleteResponse) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### GetMessage

`func (o *LiveActivityStreamDeleteResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LiveActivityStreamDeleteResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LiveActivityStreamDeleteResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LiveActivityStreamDeleteResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetForcedCompletion

`func (o *LiveActivityStreamDeleteResponse) GetForcedCompletion() bool`

GetForcedCompletion returns the ForcedCompletion field if non-nil, zero value otherwise.

### GetForcedCompletionOk

`func (o *LiveActivityStreamDeleteResponse) GetForcedCompletionOk() (*bool, bool)`

GetForcedCompletionOk returns a tuple with the ForcedCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedCompletion

`func (o *LiveActivityStreamDeleteResponse) SetForcedCompletion(v bool)`

SetForcedCompletion sets ForcedCompletion field to given value.

### HasForcedCompletion

`func (o *LiveActivityStreamDeleteResponse) HasForcedCompletion() bool`

HasForcedCompletion returns a boolean if a field has been set.

### GetWarning

`func (o *LiveActivityStreamDeleteResponse) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *LiveActivityStreamDeleteResponse) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *LiveActivityStreamDeleteResponse) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *LiveActivityStreamDeleteResponse) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetErrors

`func (o *LiveActivityStreamDeleteResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LiveActivityStreamDeleteResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LiveActivityStreamDeleteResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *LiveActivityStreamDeleteResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


