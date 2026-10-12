# PushNotificationReplaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | Additional information shown in notification and Live Activity details in ActivitySmith. Not displayed in the Push Notification or Live Activity on the device. Values must be strings, finite numbers, or booleans. At most 50 entries and 16 KB of serialized UTF-8 JSON. Omit on updates to preserve existing Metadata; send {} to clear it. | [optional] 
**Title** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**Subtitle** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** | Optional HTTPS image URL without credentials for a custom notification icon. If the image cannot be loaded, the app icon is used. iOS may omit subtitle when displaying a custom icon. | [optional] 
**Media** | Pointer to **string** | Optional HTTPS URL for an image, audio file, or video that users can preview or play when they expand the notification. If &#x60;redirection&#x60; is omitted, tapping the notification opens this URL. Cannot be combined with &#x60;actions&#x60;. | [optional] 
**Redirection** | Pointer to **string** | Optional HTTP, HTTPS, Shortcuts, or installed app URL opened when the user taps the notification body. Custom schemes such as spotify:// and spotify:track:123 require iOS 1.13.4 build 2 or later and an installed handler; no web fallback is provided. Internal and executable schemes are blocked. Overrides the default tap target from media. | [optional] 
**Actions** | Pointer to [**[]PushNotificationAction**](PushNotificationAction.md) | Optional interactive actions shown when users expand the notification. Cannot be combined with &#x60;media&#x60;. | [optional] 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 
**Badge** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** | Optional tags to organize and filter notification history. | [optional] 

## Methods

### NewPushNotificationReplaceRequest

`func NewPushNotificationReplaceRequest(title string, ) *PushNotificationReplaceRequest`

NewPushNotificationReplaceRequest instantiates a new PushNotificationReplaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushNotificationReplaceRequestWithDefaults

`func NewPushNotificationReplaceRequestWithDefaults() *PushNotificationReplaceRequest`

NewPushNotificationReplaceRequestWithDefaults instantiates a new PushNotificationReplaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *PushNotificationReplaceRequest) GetMetadata() map[string]MetadataValue`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PushNotificationReplaceRequest) GetMetadataOk() (*map[string]MetadataValue, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PushNotificationReplaceRequest) SetMetadata(v map[string]MetadataValue)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PushNotificationReplaceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTitle

`func (o *PushNotificationReplaceRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PushNotificationReplaceRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PushNotificationReplaceRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMessage

`func (o *PushNotificationReplaceRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PushNotificationReplaceRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PushNotificationReplaceRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PushNotificationReplaceRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSubtitle

`func (o *PushNotificationReplaceRequest) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *PushNotificationReplaceRequest) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *PushNotificationReplaceRequest) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *PushNotificationReplaceRequest) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetIcon

`func (o *PushNotificationReplaceRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PushNotificationReplaceRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PushNotificationReplaceRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PushNotificationReplaceRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetMedia

`func (o *PushNotificationReplaceRequest) GetMedia() string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *PushNotificationReplaceRequest) GetMediaOk() (*string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *PushNotificationReplaceRequest) SetMedia(v string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *PushNotificationReplaceRequest) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetRedirection

`func (o *PushNotificationReplaceRequest) GetRedirection() string`

GetRedirection returns the Redirection field if non-nil, zero value otherwise.

### GetRedirectionOk

`func (o *PushNotificationReplaceRequest) GetRedirectionOk() (*string, bool)`

GetRedirectionOk returns a tuple with the Redirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirection

`func (o *PushNotificationReplaceRequest) SetRedirection(v string)`

SetRedirection sets Redirection field to given value.

### HasRedirection

`func (o *PushNotificationReplaceRequest) HasRedirection() bool`

HasRedirection returns a boolean if a field has been set.

### GetActions

`func (o *PushNotificationReplaceRequest) GetActions() []PushNotificationAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PushNotificationReplaceRequest) GetActionsOk() (*[]PushNotificationAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PushNotificationReplaceRequest) SetActions(v []PushNotificationAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PushNotificationReplaceRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPayload

`func (o *PushNotificationReplaceRequest) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *PushNotificationReplaceRequest) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *PushNotificationReplaceRequest) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *PushNotificationReplaceRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetBadge

`func (o *PushNotificationReplaceRequest) GetBadge() int32`

GetBadge returns the Badge field if non-nil, zero value otherwise.

### GetBadgeOk

`func (o *PushNotificationReplaceRequest) GetBadgeOk() (*int32, bool)`

GetBadgeOk returns a tuple with the Badge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadge

`func (o *PushNotificationReplaceRequest) SetBadge(v int32)`

SetBadge sets Badge field to given value.

### HasBadge

`func (o *PushNotificationReplaceRequest) HasBadge() bool`

HasBadge returns a boolean if a field has been set.

### GetTags

`func (o *PushNotificationReplaceRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PushNotificationReplaceRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PushNotificationReplaceRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PushNotificationReplaceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


