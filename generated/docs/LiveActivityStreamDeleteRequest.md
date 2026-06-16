# LiveActivityStreamDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentState** | Pointer to [**StreamContentState**](StreamContentState.md) |  | [optional] 
**Action** | Pointer to [**LiveActivityAction**](LiveActivityAction.md) |  | [optional] 
**SecondaryAction** | Pointer to [**LiveActivityAction**](LiveActivityAction.md) | Optional secondary action button. Supported only for alert, progress, and segmented_progress Live Activities. Uses the same open_url, shortcuts://, and webhook shapes as action. | [optional] 
**Alert** | Pointer to [**AlertPayload**](AlertPayload.md) |  | [optional] 

## Methods

### NewLiveActivityStreamDeleteRequest

`func NewLiveActivityStreamDeleteRequest() *LiveActivityStreamDeleteRequest`

NewLiveActivityStreamDeleteRequest instantiates a new LiveActivityStreamDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityStreamDeleteRequestWithDefaults

`func NewLiveActivityStreamDeleteRequestWithDefaults() *LiveActivityStreamDeleteRequest`

NewLiveActivityStreamDeleteRequestWithDefaults instantiates a new LiveActivityStreamDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentState

`func (o *LiveActivityStreamDeleteRequest) GetContentState() StreamContentState`

GetContentState returns the ContentState field if non-nil, zero value otherwise.

### GetContentStateOk

`func (o *LiveActivityStreamDeleteRequest) GetContentStateOk() (*StreamContentState, bool)`

GetContentStateOk returns a tuple with the ContentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentState

`func (o *LiveActivityStreamDeleteRequest) SetContentState(v StreamContentState)`

SetContentState sets ContentState field to given value.

### HasContentState

`func (o *LiveActivityStreamDeleteRequest) HasContentState() bool`

HasContentState returns a boolean if a field has been set.

### GetAction

`func (o *LiveActivityStreamDeleteRequest) GetAction() LiveActivityAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LiveActivityStreamDeleteRequest) GetActionOk() (*LiveActivityAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LiveActivityStreamDeleteRequest) SetAction(v LiveActivityAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *LiveActivityStreamDeleteRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSecondaryAction

`func (o *LiveActivityStreamDeleteRequest) GetSecondaryAction() LiveActivityAction`

GetSecondaryAction returns the SecondaryAction field if non-nil, zero value otherwise.

### GetSecondaryActionOk

`func (o *LiveActivityStreamDeleteRequest) GetSecondaryActionOk() (*LiveActivityAction, bool)`

GetSecondaryActionOk returns a tuple with the SecondaryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAction

`func (o *LiveActivityStreamDeleteRequest) SetSecondaryAction(v LiveActivityAction)`

SetSecondaryAction sets SecondaryAction field to given value.

### HasSecondaryAction

`func (o *LiveActivityStreamDeleteRequest) HasSecondaryAction() bool`

HasSecondaryAction returns a boolean if a field has been set.

### GetAlert

`func (o *LiveActivityStreamDeleteRequest) GetAlert() AlertPayload`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *LiveActivityStreamDeleteRequest) GetAlertOk() (*AlertPayload, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *LiveActivityStreamDeleteRequest) SetAlert(v AlertPayload)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *LiveActivityStreamDeleteRequest) HasAlert() bool`

HasAlert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


