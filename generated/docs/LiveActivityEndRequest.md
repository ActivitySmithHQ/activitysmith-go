# LiveActivityEndRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | Additional information shown in notification and Live Activity details in ActivitySmith. Not displayed in the Push Notification or Live Activity on the device. Values must be strings, finite numbers, or booleans. At most 50 entries and 16 KB of serialized UTF-8 JSON. Omit on updates to preserve existing Metadata; send {} to clear it. | [optional] 
**ActivityId** | **string** |  | 
**Tags** | Pointer to **[]string** | Tags for notification history. Omit to keep existing Tags, supply an array to replace them, or send an empty array to clear them. | [optional] 
**ContentState** | [**ContentStateEnd**](ContentStateEnd.md) |  | 
**Action** | Pointer to [**LiveActivityAction**](LiveActivityAction.md) |  | [optional] 
**SecondaryAction** | Pointer to [**LiveActivityAction**](LiveActivityAction.md) | Optional secondary action button. Supported for alert, progress, segmented_progress, and value Live Activities. Uses the same open_url, shortcuts://, and webhook shapes as action. | [optional] 
**Sound** | Pointer to **string** | Optional sound for the end. Ends are silent unless you set one. iOS plays the sound only while the screen is on. Sounds bundled with the ActivitySmith app: &#x60;arrive&#x60;, &#x60;bell&#x60;, &#x60;bells&#x60;, &#x60;bitflip&#x60;, &#x60;bloom&#x60;, &#x60;callout&#x60;, &#x60;chime&#x60;, &#x60;coin&#x60;, &#x60;corrupt&#x60;, &#x60;crest&#x60;, &#x60;cue&#x60;, &#x60;ding&#x60;, &#x60;drop&#x60;, &#x60;ember&#x60;, &#x60;emblem&#x60;, &#x60;flourish&#x60;, &#x60;glimmer&#x60;, &#x60;glint&#x60;, &#x60;glitch&#x60;, &#x60;lantern&#x60;, &#x60;nightfall&#x60;, &#x60;nudge&#x60;, &#x60;pixel&#x60;, &#x60;pop&#x60;, &#x60;signal&#x60;, &#x60;streak&#x60;, &#x60;tally&#x60;, &#x60;wand&#x60;, &#x60;whoosh&#x60;, &#x60;zip&#x60;. Use &#x60;default&#x60; for the sound picked on the device, or &#x60;none&#x60; for no sound. | [optional] 

## Methods

### NewLiveActivityEndRequest

`func NewLiveActivityEndRequest(activityId string, contentState ContentStateEnd, ) *LiveActivityEndRequest`

NewLiveActivityEndRequest instantiates a new LiveActivityEndRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityEndRequestWithDefaults

`func NewLiveActivityEndRequestWithDefaults() *LiveActivityEndRequest`

NewLiveActivityEndRequestWithDefaults instantiates a new LiveActivityEndRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *LiveActivityEndRequest) GetMetadata() map[string]MetadataValue`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LiveActivityEndRequest) GetMetadataOk() (*map[string]MetadataValue, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LiveActivityEndRequest) SetMetadata(v map[string]MetadataValue)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LiveActivityEndRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetActivityId

`func (o *LiveActivityEndRequest) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *LiveActivityEndRequest) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *LiveActivityEndRequest) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.


### GetTags

`func (o *LiveActivityEndRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LiveActivityEndRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LiveActivityEndRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LiveActivityEndRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetContentState

`func (o *LiveActivityEndRequest) GetContentState() ContentStateEnd`

GetContentState returns the ContentState field if non-nil, zero value otherwise.

### GetContentStateOk

`func (o *LiveActivityEndRequest) GetContentStateOk() (*ContentStateEnd, bool)`

GetContentStateOk returns a tuple with the ContentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentState

`func (o *LiveActivityEndRequest) SetContentState(v ContentStateEnd)`

SetContentState sets ContentState field to given value.


### GetAction

`func (o *LiveActivityEndRequest) GetAction() LiveActivityAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LiveActivityEndRequest) GetActionOk() (*LiveActivityAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LiveActivityEndRequest) SetAction(v LiveActivityAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *LiveActivityEndRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSecondaryAction

`func (o *LiveActivityEndRequest) GetSecondaryAction() LiveActivityAction`

GetSecondaryAction returns the SecondaryAction field if non-nil, zero value otherwise.

### GetSecondaryActionOk

`func (o *LiveActivityEndRequest) GetSecondaryActionOk() (*LiveActivityAction, bool)`

GetSecondaryActionOk returns a tuple with the SecondaryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAction

`func (o *LiveActivityEndRequest) SetSecondaryAction(v LiveActivityAction)`

SetSecondaryAction sets SecondaryAction field to given value.

### HasSecondaryAction

`func (o *LiveActivityEndRequest) HasSecondaryAction() bool`

HasSecondaryAction returns a boolean if a field has been set.

### GetSound

`func (o *LiveActivityEndRequest) GetSound() string`

GetSound returns the Sound field if non-nil, zero value otherwise.

### GetSoundOk

`func (o *LiveActivityEndRequest) GetSoundOk() (*string, bool)`

GetSoundOk returns a tuple with the Sound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSound

`func (o *LiveActivityEndRequest) SetSound(v string)`

SetSound sets Sound field to given value.

### HasSound

`func (o *LiveActivityEndRequest) HasSound() bool`

HasSound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


