# PushNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**Subtitle** | Pointer to **string** |  | [optional] 
**Media** | Pointer to **string** | Optional HTTPS URL for an image, audio file, or video that users can preview or play when they expand the notification. If &#x60;redirection&#x60; is omitted, tapping the notification opens this URL. Cannot be combined with &#x60;actions&#x60;. | [optional] 
**Redirection** | Pointer to **string** | Optional HTTPS URL opened when user taps the notification body. Overrides the default tap target from &#x60;media&#x60; when both are provided. | [optional] 
**Actions** | Pointer to [**[]PushNotificationAction**](PushNotificationAction.md) | Optional interactive actions shown when users expand the notification. Cannot be combined with &#x60;media&#x60;. | [optional] 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 
**Badge** | Pointer to **int32** |  | [optional] 
**Sound** | Pointer to **string** |  | [optional] 
**Target** | Pointer to [**ChannelTarget**](ChannelTarget.md) |  | [optional] 

## Methods

### NewPushNotificationRequest

`func NewPushNotificationRequest(title string, ) *PushNotificationRequest`

NewPushNotificationRequest instantiates a new PushNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushNotificationRequestWithDefaults

`func NewPushNotificationRequestWithDefaults() *PushNotificationRequest`

NewPushNotificationRequestWithDefaults instantiates a new PushNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PushNotificationRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PushNotificationRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PushNotificationRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMessage

`func (o *PushNotificationRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PushNotificationRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PushNotificationRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PushNotificationRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSubtitle

`func (o *PushNotificationRequest) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *PushNotificationRequest) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *PushNotificationRequest) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *PushNotificationRequest) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetMedia

`func (o *PushNotificationRequest) GetMedia() string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *PushNotificationRequest) GetMediaOk() (*string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *PushNotificationRequest) SetMedia(v string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *PushNotificationRequest) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetRedirection

`func (o *PushNotificationRequest) GetRedirection() string`

GetRedirection returns the Redirection field if non-nil, zero value otherwise.

### GetRedirectionOk

`func (o *PushNotificationRequest) GetRedirectionOk() (*string, bool)`

GetRedirectionOk returns a tuple with the Redirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirection

`func (o *PushNotificationRequest) SetRedirection(v string)`

SetRedirection sets Redirection field to given value.

### HasRedirection

`func (o *PushNotificationRequest) HasRedirection() bool`

HasRedirection returns a boolean if a field has been set.

### GetActions

`func (o *PushNotificationRequest) GetActions() []PushNotificationAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PushNotificationRequest) GetActionsOk() (*[]PushNotificationAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PushNotificationRequest) SetActions(v []PushNotificationAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PushNotificationRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPayload

`func (o *PushNotificationRequest) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *PushNotificationRequest) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *PushNotificationRequest) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *PushNotificationRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetBadge

`func (o *PushNotificationRequest) GetBadge() int32`

GetBadge returns the Badge field if non-nil, zero value otherwise.

### GetBadgeOk

`func (o *PushNotificationRequest) GetBadgeOk() (*int32, bool)`

GetBadgeOk returns a tuple with the Badge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadge

`func (o *PushNotificationRequest) SetBadge(v int32)`

SetBadge sets Badge field to given value.

### HasBadge

`func (o *PushNotificationRequest) HasBadge() bool`

HasBadge returns a boolean if a field has been set.

### GetSound

`func (o *PushNotificationRequest) GetSound() string`

GetSound returns the Sound field if non-nil, zero value otherwise.

### GetSoundOk

`func (o *PushNotificationRequest) GetSoundOk() (*string, bool)`

GetSoundOk returns a tuple with the Sound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSound

`func (o *PushNotificationRequest) SetSound(v string)`

SetSound sets Sound field to given value.

### HasSound

`func (o *PushNotificationRequest) HasSound() bool`

HasSound returns a boolean if a field has been set.

### GetTarget

`func (o *PushNotificationRequest) GetTarget() ChannelTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PushNotificationRequest) GetTargetOk() (*ChannelTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PushNotificationRequest) SetTarget(v ChannelTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PushNotificationRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


