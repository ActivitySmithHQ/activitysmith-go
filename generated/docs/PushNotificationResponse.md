# PushNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | True when APNs accepts the notification for at least one targeted device. | 
**NotificationId** | Pointer to **string** | ID of this Push Notification. Pass it to &#x60;PUT /push-notification/{notification_id}&#x60; to replace the notification later. | [optional] 
**DevicesNotified** | Pointer to **int32** | Number of device tokens for which APNs accepted the notification, without confirming on-device delivery. | [optional] 
**UsersNotified** | Pointer to **int32** |  | [optional] 
**EffectiveChannelSlugs** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** | Optional tags to organize and filter notification history. | [optional] 
**Timestamp** | **time.Time** |  | 
**DevicesTargeted** | Pointer to **int32** | Number of device tokens targeted. | [optional] 
**DevicesDeactivated** | Pointer to **int32** | Number of device tokens deactivated because APNs reported them unregistered. | [optional] 
**DevicesUnregistered** | Pointer to **int32** | Number of targeted device tokens APNs reported as unregistered. | [optional] 
**CredentialTest** | Pointer to **bool** | True when the request was an integration connection test and no paired devices were available. | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewPushNotificationResponse

`func NewPushNotificationResponse(success bool, timestamp time.Time, ) *PushNotificationResponse`

NewPushNotificationResponse instantiates a new PushNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushNotificationResponseWithDefaults

`func NewPushNotificationResponseWithDefaults() *PushNotificationResponse`

NewPushNotificationResponseWithDefaults instantiates a new PushNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *PushNotificationResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PushNotificationResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PushNotificationResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetNotificationId

`func (o *PushNotificationResponse) GetNotificationId() string`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *PushNotificationResponse) GetNotificationIdOk() (*string, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *PushNotificationResponse) SetNotificationId(v string)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *PushNotificationResponse) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetDevicesNotified

`func (o *PushNotificationResponse) GetDevicesNotified() int32`

GetDevicesNotified returns the DevicesNotified field if non-nil, zero value otherwise.

### GetDevicesNotifiedOk

`func (o *PushNotificationResponse) GetDevicesNotifiedOk() (*int32, bool)`

GetDevicesNotifiedOk returns a tuple with the DevicesNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesNotified

`func (o *PushNotificationResponse) SetDevicesNotified(v int32)`

SetDevicesNotified sets DevicesNotified field to given value.

### HasDevicesNotified

`func (o *PushNotificationResponse) HasDevicesNotified() bool`

HasDevicesNotified returns a boolean if a field has been set.

### GetUsersNotified

`func (o *PushNotificationResponse) GetUsersNotified() int32`

GetUsersNotified returns the UsersNotified field if non-nil, zero value otherwise.

### GetUsersNotifiedOk

`func (o *PushNotificationResponse) GetUsersNotifiedOk() (*int32, bool)`

GetUsersNotifiedOk returns a tuple with the UsersNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersNotified

`func (o *PushNotificationResponse) SetUsersNotified(v int32)`

SetUsersNotified sets UsersNotified field to given value.

### HasUsersNotified

`func (o *PushNotificationResponse) HasUsersNotified() bool`

HasUsersNotified returns a boolean if a field has been set.

### GetEffectiveChannelSlugs

`func (o *PushNotificationResponse) GetEffectiveChannelSlugs() []string`

GetEffectiveChannelSlugs returns the EffectiveChannelSlugs field if non-nil, zero value otherwise.

### GetEffectiveChannelSlugsOk

`func (o *PushNotificationResponse) GetEffectiveChannelSlugsOk() (*[]string, bool)`

GetEffectiveChannelSlugsOk returns a tuple with the EffectiveChannelSlugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveChannelSlugs

`func (o *PushNotificationResponse) SetEffectiveChannelSlugs(v []string)`

SetEffectiveChannelSlugs sets EffectiveChannelSlugs field to given value.

### HasEffectiveChannelSlugs

`func (o *PushNotificationResponse) HasEffectiveChannelSlugs() bool`

HasEffectiveChannelSlugs returns a boolean if a field has been set.

### GetTags

`func (o *PushNotificationResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PushNotificationResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PushNotificationResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PushNotificationResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *PushNotificationResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PushNotificationResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PushNotificationResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetDevicesTargeted

`func (o *PushNotificationResponse) GetDevicesTargeted() int32`

GetDevicesTargeted returns the DevicesTargeted field if non-nil, zero value otherwise.

### GetDevicesTargetedOk

`func (o *PushNotificationResponse) GetDevicesTargetedOk() (*int32, bool)`

GetDevicesTargetedOk returns a tuple with the DevicesTargeted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesTargeted

`func (o *PushNotificationResponse) SetDevicesTargeted(v int32)`

SetDevicesTargeted sets DevicesTargeted field to given value.

### HasDevicesTargeted

`func (o *PushNotificationResponse) HasDevicesTargeted() bool`

HasDevicesTargeted returns a boolean if a field has been set.

### GetDevicesDeactivated

`func (o *PushNotificationResponse) GetDevicesDeactivated() int32`

GetDevicesDeactivated returns the DevicesDeactivated field if non-nil, zero value otherwise.

### GetDevicesDeactivatedOk

`func (o *PushNotificationResponse) GetDevicesDeactivatedOk() (*int32, bool)`

GetDevicesDeactivatedOk returns a tuple with the DevicesDeactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesDeactivated

`func (o *PushNotificationResponse) SetDevicesDeactivated(v int32)`

SetDevicesDeactivated sets DevicesDeactivated field to given value.

### HasDevicesDeactivated

`func (o *PushNotificationResponse) HasDevicesDeactivated() bool`

HasDevicesDeactivated returns a boolean if a field has been set.

### GetDevicesUnregistered

`func (o *PushNotificationResponse) GetDevicesUnregistered() int32`

GetDevicesUnregistered returns the DevicesUnregistered field if non-nil, zero value otherwise.

### GetDevicesUnregisteredOk

`func (o *PushNotificationResponse) GetDevicesUnregisteredOk() (*int32, bool)`

GetDevicesUnregisteredOk returns a tuple with the DevicesUnregistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesUnregistered

`func (o *PushNotificationResponse) SetDevicesUnregistered(v int32)`

SetDevicesUnregistered sets DevicesUnregistered field to given value.

### HasDevicesUnregistered

`func (o *PushNotificationResponse) HasDevicesUnregistered() bool`

HasDevicesUnregistered returns a boolean if a field has been set.

### GetCredentialTest

`func (o *PushNotificationResponse) GetCredentialTest() bool`

GetCredentialTest returns the CredentialTest field if non-nil, zero value otherwise.

### GetCredentialTestOk

`func (o *PushNotificationResponse) GetCredentialTestOk() (*bool, bool)`

GetCredentialTestOk returns a tuple with the CredentialTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialTest

`func (o *PushNotificationResponse) SetCredentialTest(v bool)`

SetCredentialTest sets CredentialTest field to given value.

### HasCredentialTest

`func (o *PushNotificationResponse) HasCredentialTest() bool`

HasCredentialTest returns a boolean if a field has been set.

### GetMessage

`func (o *PushNotificationResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PushNotificationResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PushNotificationResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PushNotificationResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


