# LiveActivityStartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentState** | [**ContentStateStart**](ContentStateStart.md) |  | 
**Alert** | Pointer to [**AlertPayload**](AlertPayload.md) |  | [optional] 
**Target** | Pointer to [**ChannelTarget**](ChannelTarget.md) |  | [optional] 

## Methods

### NewLiveActivityStartRequest

`func NewLiveActivityStartRequest(contentState ContentStateStart, ) *LiveActivityStartRequest`

NewLiveActivityStartRequest instantiates a new LiveActivityStartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityStartRequestWithDefaults

`func NewLiveActivityStartRequestWithDefaults() *LiveActivityStartRequest`

NewLiveActivityStartRequestWithDefaults instantiates a new LiveActivityStartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentState

`func (o *LiveActivityStartRequest) GetContentState() ContentStateStart`

GetContentState returns the ContentState field if non-nil, zero value otherwise.

### GetContentStateOk

`func (o *LiveActivityStartRequest) GetContentStateOk() (*ContentStateStart, bool)`

GetContentStateOk returns a tuple with the ContentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentState

`func (o *LiveActivityStartRequest) SetContentState(v ContentStateStart)`

SetContentState sets ContentState field to given value.


### GetAlert

`func (o *LiveActivityStartRequest) GetAlert() AlertPayload`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *LiveActivityStartRequest) GetAlertOk() (*AlertPayload, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *LiveActivityStartRequest) SetAlert(v AlertPayload)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *LiveActivityStartRequest) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetTarget

`func (o *LiveActivityStartRequest) GetTarget() ChannelTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LiveActivityStartRequest) GetTargetOk() (*ChannelTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *LiveActivityStartRequest) SetTarget(v ChannelTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *LiveActivityStartRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


