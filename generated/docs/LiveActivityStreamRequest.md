# LiveActivityStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentState** | [**StreamContentState**](StreamContentState.md) |  | 
**Action** | Pointer to [**LiveActivityAction**](LiveActivityAction.md) |  | [optional] 
**Alert** | Pointer to [**AlertPayload**](AlertPayload.md) |  | [optional] 
**Channels** | Pointer to **[]string** | Channel slugs. When omitted, API key scope determines recipients. | [optional] 
**Target** | Pointer to [**ChannelTarget**](ChannelTarget.md) |  | [optional] 

## Methods

### NewLiveActivityStreamRequest

`func NewLiveActivityStreamRequest(contentState StreamContentState, ) *LiveActivityStreamRequest`

NewLiveActivityStreamRequest instantiates a new LiveActivityStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityStreamRequestWithDefaults

`func NewLiveActivityStreamRequestWithDefaults() *LiveActivityStreamRequest`

NewLiveActivityStreamRequestWithDefaults instantiates a new LiveActivityStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentState

`func (o *LiveActivityStreamRequest) GetContentState() StreamContentState`

GetContentState returns the ContentState field if non-nil, zero value otherwise.

### GetContentStateOk

`func (o *LiveActivityStreamRequest) GetContentStateOk() (*StreamContentState, bool)`

GetContentStateOk returns a tuple with the ContentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentState

`func (o *LiveActivityStreamRequest) SetContentState(v StreamContentState)`

SetContentState sets ContentState field to given value.


### GetAction

`func (o *LiveActivityStreamRequest) GetAction() LiveActivityAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LiveActivityStreamRequest) GetActionOk() (*LiveActivityAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LiveActivityStreamRequest) SetAction(v LiveActivityAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *LiveActivityStreamRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAlert

`func (o *LiveActivityStreamRequest) GetAlert() AlertPayload`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *LiveActivityStreamRequest) GetAlertOk() (*AlertPayload, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *LiveActivityStreamRequest) SetAlert(v AlertPayload)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *LiveActivityStreamRequest) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetChannels

`func (o *LiveActivityStreamRequest) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *LiveActivityStreamRequest) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *LiveActivityStreamRequest) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *LiveActivityStreamRequest) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetTarget

`func (o *LiveActivityStreamRequest) GetTarget() ChannelTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LiveActivityStreamRequest) GetTargetOk() (*ChannelTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *LiveActivityStreamRequest) SetTarget(v ChannelTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *LiveActivityStreamRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


