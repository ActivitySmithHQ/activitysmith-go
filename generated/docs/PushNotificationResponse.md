# PushNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**DevicesNotified** | Pointer to **int32** |  | [optional] 
**UsersNotified** | Pointer to **int32** |  | [optional] 
**Timestamp** | **time.Time** |  | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


