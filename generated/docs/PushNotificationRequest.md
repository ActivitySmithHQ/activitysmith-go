# PushNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**Subtitle** | Pointer to **string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


