# LiveActivityStreamDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | Additional information shown in notification and Live Activity details in ActivitySmith. Not displayed in the Push Notification or Live Activity on the device. Values must be strings, finite numbers, or booleans. At most 50 entries and 16 KB of serialized UTF-8 JSON. Omit on updates to preserve existing Metadata; send {} to clear it. | [optional] 
**Tags** | Pointer to **[]string** | Optional tags to organize and filter notification history. | [optional] 
**ContentState** | Pointer to [**StreamContentState**](StreamContentState.md) |  | [optional] 
**Action** | Pointer to [**LiveActivityAction**](LiveActivityAction.md) |  | [optional] 
**SecondaryAction** | Pointer to [**LiveActivityAction**](LiveActivityAction.md) | Optional secondary action button. Supported for alert, progress, segmented_progress, and value Live Activities. Uses the same open_url, shortcuts://, and webhook shapes as action. | [optional] 
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

### GetMetadata

`func (o *LiveActivityStreamDeleteRequest) GetMetadata() map[string]MetadataValue`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LiveActivityStreamDeleteRequest) GetMetadataOk() (*map[string]MetadataValue, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LiveActivityStreamDeleteRequest) SetMetadata(v map[string]MetadataValue)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LiveActivityStreamDeleteRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTags

`func (o *LiveActivityStreamDeleteRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LiveActivityStreamDeleteRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LiveActivityStreamDeleteRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LiveActivityStreamDeleteRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

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


