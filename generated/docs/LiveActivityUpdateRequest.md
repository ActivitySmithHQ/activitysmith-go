# LiveActivityUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | Additional information shown in notification and Live Activity details in ActivitySmith. Not displayed in the Push Notification or Live Activity on the device. Values must be strings, finite numbers, or booleans. At most 50 entries and 16 KB of serialized UTF-8 JSON. Omit on updates to preserve existing Metadata; send {} to clear it. | [optional] 
**ActivityId** | **string** |  | 
**Tags** | Pointer to **[]string** | Tags for notification history. Omit to keep existing Tags, supply an array to replace them, or send an empty array to clear them. | [optional] 
**ContentState** | [**ContentStateUpdate**](ContentStateUpdate.md) |  | 
**Action** | Pointer to [**LiveActivityAction**](LiveActivityAction.md) |  | [optional] 
**SecondaryAction** | Pointer to [**LiveActivityAction**](LiveActivityAction.md) | Optional secondary action button. Supported for alert, progress, segmented_progress, and value Live Activities. Uses the same open_url, shortcuts://, and webhook shapes as action. | [optional] 

## Methods

### NewLiveActivityUpdateRequest

`func NewLiveActivityUpdateRequest(activityId string, contentState ContentStateUpdate, ) *LiveActivityUpdateRequest`

NewLiveActivityUpdateRequest instantiates a new LiveActivityUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityUpdateRequestWithDefaults

`func NewLiveActivityUpdateRequestWithDefaults() *LiveActivityUpdateRequest`

NewLiveActivityUpdateRequestWithDefaults instantiates a new LiveActivityUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *LiveActivityUpdateRequest) GetMetadata() map[string]MetadataValue`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LiveActivityUpdateRequest) GetMetadataOk() (*map[string]MetadataValue, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LiveActivityUpdateRequest) SetMetadata(v map[string]MetadataValue)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LiveActivityUpdateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetActivityId

`func (o *LiveActivityUpdateRequest) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *LiveActivityUpdateRequest) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *LiveActivityUpdateRequest) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.


### GetTags

`func (o *LiveActivityUpdateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LiveActivityUpdateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LiveActivityUpdateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LiveActivityUpdateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetContentState

`func (o *LiveActivityUpdateRequest) GetContentState() ContentStateUpdate`

GetContentState returns the ContentState field if non-nil, zero value otherwise.

### GetContentStateOk

`func (o *LiveActivityUpdateRequest) GetContentStateOk() (*ContentStateUpdate, bool)`

GetContentStateOk returns a tuple with the ContentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentState

`func (o *LiveActivityUpdateRequest) SetContentState(v ContentStateUpdate)`

SetContentState sets ContentState field to given value.


### GetAction

`func (o *LiveActivityUpdateRequest) GetAction() LiveActivityAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LiveActivityUpdateRequest) GetActionOk() (*LiveActivityAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LiveActivityUpdateRequest) SetAction(v LiveActivityAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *LiveActivityUpdateRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSecondaryAction

`func (o *LiveActivityUpdateRequest) GetSecondaryAction() LiveActivityAction`

GetSecondaryAction returns the SecondaryAction field if non-nil, zero value otherwise.

### GetSecondaryActionOk

`func (o *LiveActivityUpdateRequest) GetSecondaryActionOk() (*LiveActivityAction, bool)`

GetSecondaryActionOk returns a tuple with the SecondaryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAction

`func (o *LiveActivityUpdateRequest) SetSecondaryAction(v LiveActivityAction)`

SetSecondaryAction sets SecondaryAction field to given value.

### HasSecondaryAction

`func (o *LiveActivityUpdateRequest) HasSecondaryAction() bool`

HasSecondaryAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


