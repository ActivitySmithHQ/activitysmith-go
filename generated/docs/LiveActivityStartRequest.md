# LiveActivityStartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | Additional information shown in notification and Live Activity details in ActivitySmith. Not displayed in the Push Notification or Live Activity on the device. Values must be strings, finite numbers, or booleans. At most 50 entries and 16 KB of serialized UTF-8 JSON. Omit on updates to preserve existing Metadata; send {} to clear it. | [optional] 
**ContentState** | [**ContentStateStart**](ContentStateStart.md) |  | 
**Action** | Pointer to [**LiveActivityAction**](LiveActivityAction.md) |  | [optional] 
**SecondaryAction** | Pointer to [**LiveActivityAction**](LiveActivityAction.md) | Optional secondary action button. Supported for alert, progress, segmented_progress, and value Live Activities. Uses the same open_url, shortcuts://, and webhook shapes as action. | [optional] 
**Alert** | Pointer to [**AlertPayload**](AlertPayload.md) |  | [optional] 
**Target** | Pointer to [**ChannelTarget**](ChannelTarget.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Optional tags to organize and filter notification history. | [optional] 

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

### GetMetadata

`func (o *LiveActivityStartRequest) GetMetadata() map[string]MetadataValue`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LiveActivityStartRequest) GetMetadataOk() (*map[string]MetadataValue, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LiveActivityStartRequest) SetMetadata(v map[string]MetadataValue)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LiveActivityStartRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

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


### GetAction

`func (o *LiveActivityStartRequest) GetAction() LiveActivityAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LiveActivityStartRequest) GetActionOk() (*LiveActivityAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LiveActivityStartRequest) SetAction(v LiveActivityAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *LiveActivityStartRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSecondaryAction

`func (o *LiveActivityStartRequest) GetSecondaryAction() LiveActivityAction`

GetSecondaryAction returns the SecondaryAction field if non-nil, zero value otherwise.

### GetSecondaryActionOk

`func (o *LiveActivityStartRequest) GetSecondaryActionOk() (*LiveActivityAction, bool)`

GetSecondaryActionOk returns a tuple with the SecondaryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAction

`func (o *LiveActivityStartRequest) SetSecondaryAction(v LiveActivityAction)`

SetSecondaryAction sets SecondaryAction field to given value.

### HasSecondaryAction

`func (o *LiveActivityStartRequest) HasSecondaryAction() bool`

HasSecondaryAction returns a boolean if a field has been set.

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

### GetTags

`func (o *LiveActivityStartRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LiveActivityStartRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LiveActivityStartRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LiveActivityStartRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


