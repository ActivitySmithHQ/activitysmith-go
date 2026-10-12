# PushNotificationReplaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | True when APNs accepts the replacement for at least one device. | 
**NotificationId** | **string** | ID of the replaced Push Notification. It stays the same, so you can replace it again. | 
**DevicesNotified** | Pointer to **int32** | Number of device tokens for which APNs accepted the notification, without confirming on-device delivery. | [optional] 
**UsersNotified** | Pointer to **int32** |  | [optional] 
**EffectiveChannelSlugs** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** | Optional tags to organize and filter notification history. | [optional] 
**Timestamp** | **time.Time** |  | 

## Methods

### NewPushNotificationReplaceResponse

`func NewPushNotificationReplaceResponse(success bool, notificationId string, timestamp time.Time, ) *PushNotificationReplaceResponse`

NewPushNotificationReplaceResponse instantiates a new PushNotificationReplaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushNotificationReplaceResponseWithDefaults

`func NewPushNotificationReplaceResponseWithDefaults() *PushNotificationReplaceResponse`

NewPushNotificationReplaceResponseWithDefaults instantiates a new PushNotificationReplaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *PushNotificationReplaceResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PushNotificationReplaceResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PushNotificationReplaceResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetNotificationId

`func (o *PushNotificationReplaceResponse) GetNotificationId() string`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *PushNotificationReplaceResponse) GetNotificationIdOk() (*string, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *PushNotificationReplaceResponse) SetNotificationId(v string)`

SetNotificationId sets NotificationId field to given value.


### GetDevicesNotified

`func (o *PushNotificationReplaceResponse) GetDevicesNotified() int32`

GetDevicesNotified returns the DevicesNotified field if non-nil, zero value otherwise.

### GetDevicesNotifiedOk

`func (o *PushNotificationReplaceResponse) GetDevicesNotifiedOk() (*int32, bool)`

GetDevicesNotifiedOk returns a tuple with the DevicesNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesNotified

`func (o *PushNotificationReplaceResponse) SetDevicesNotified(v int32)`

SetDevicesNotified sets DevicesNotified field to given value.

### HasDevicesNotified

`func (o *PushNotificationReplaceResponse) HasDevicesNotified() bool`

HasDevicesNotified returns a boolean if a field has been set.

### GetUsersNotified

`func (o *PushNotificationReplaceResponse) GetUsersNotified() int32`

GetUsersNotified returns the UsersNotified field if non-nil, zero value otherwise.

### GetUsersNotifiedOk

`func (o *PushNotificationReplaceResponse) GetUsersNotifiedOk() (*int32, bool)`

GetUsersNotifiedOk returns a tuple with the UsersNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersNotified

`func (o *PushNotificationReplaceResponse) SetUsersNotified(v int32)`

SetUsersNotified sets UsersNotified field to given value.

### HasUsersNotified

`func (o *PushNotificationReplaceResponse) HasUsersNotified() bool`

HasUsersNotified returns a boolean if a field has been set.

### GetEffectiveChannelSlugs

`func (o *PushNotificationReplaceResponse) GetEffectiveChannelSlugs() []string`

GetEffectiveChannelSlugs returns the EffectiveChannelSlugs field if non-nil, zero value otherwise.

### GetEffectiveChannelSlugsOk

`func (o *PushNotificationReplaceResponse) GetEffectiveChannelSlugsOk() (*[]string, bool)`

GetEffectiveChannelSlugsOk returns a tuple with the EffectiveChannelSlugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveChannelSlugs

`func (o *PushNotificationReplaceResponse) SetEffectiveChannelSlugs(v []string)`

SetEffectiveChannelSlugs sets EffectiveChannelSlugs field to given value.

### HasEffectiveChannelSlugs

`func (o *PushNotificationReplaceResponse) HasEffectiveChannelSlugs() bool`

HasEffectiveChannelSlugs returns a boolean if a field has been set.

### GetTags

`func (o *PushNotificationReplaceResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PushNotificationReplaceResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PushNotificationReplaceResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PushNotificationReplaceResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *PushNotificationReplaceResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PushNotificationReplaceResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PushNotificationReplaceResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


