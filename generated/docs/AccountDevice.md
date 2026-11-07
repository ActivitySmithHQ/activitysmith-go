# AccountDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserEmail** | **string** | Email of the account member who paired the device. | 
**DeviceId** | **string** | Device ID. Use it to assign the device to Channels. | 
**DeviceType** | **NullableString** | Device platform, such as ios. iPhones and iPads both report ios; use system_name to tell them apart. | 
**IsActive** | **bool** | True when the device can receive Push Notifications or Live Activities. | 
**TokenStatus** | **map[string]bool** | Whether each delivery token is active, keyed by token type (push, start_live_activity). | 
**AppVersion** | **NullableString** | ActivitySmith app version on the device. | 
**DeviceName** | **NullableString** | Device name set on the device. | 
**SystemName** | **NullableString** | Operating system name, such as iOS or iPadOS. | 
**SystemVersion** | **NullableString** | Operating system version. | 
**CreatedAt** | **NullableTime** | When the device was first paired. | 
**UpdatedAt** | **NullableTime** | When the device last updated its pairing. | 

## Methods

### NewAccountDevice

`func NewAccountDevice(userEmail string, deviceId string, deviceType NullableString, isActive bool, tokenStatus map[string]bool, appVersion NullableString, deviceName NullableString, systemName NullableString, systemVersion NullableString, createdAt NullableTime, updatedAt NullableTime, ) *AccountDevice`

NewAccountDevice instantiates a new AccountDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDeviceWithDefaults

`func NewAccountDeviceWithDefaults() *AccountDevice`

NewAccountDeviceWithDefaults instantiates a new AccountDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserEmail

`func (o *AccountDevice) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *AccountDevice) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *AccountDevice) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetDeviceId

`func (o *AccountDevice) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AccountDevice) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AccountDevice) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetDeviceType

`func (o *AccountDevice) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *AccountDevice) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *AccountDevice) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### SetDeviceTypeNil

`func (o *AccountDevice) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *AccountDevice) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetIsActive

`func (o *AccountDevice) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AccountDevice) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AccountDevice) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetTokenStatus

`func (o *AccountDevice) GetTokenStatus() map[string]bool`

GetTokenStatus returns the TokenStatus field if non-nil, zero value otherwise.

### GetTokenStatusOk

`func (o *AccountDevice) GetTokenStatusOk() (*map[string]bool, bool)`

GetTokenStatusOk returns a tuple with the TokenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenStatus

`func (o *AccountDevice) SetTokenStatus(v map[string]bool)`

SetTokenStatus sets TokenStatus field to given value.


### GetAppVersion

`func (o *AccountDevice) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *AccountDevice) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *AccountDevice) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.


### SetAppVersionNil

`func (o *AccountDevice) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *AccountDevice) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetDeviceName

`func (o *AccountDevice) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *AccountDevice) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *AccountDevice) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### SetDeviceNameNil

`func (o *AccountDevice) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *AccountDevice) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetSystemName

`func (o *AccountDevice) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *AccountDevice) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *AccountDevice) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.


### SetSystemNameNil

`func (o *AccountDevice) SetSystemNameNil(b bool)`

 SetSystemNameNil sets the value for SystemName to be an explicit nil

### UnsetSystemName
`func (o *AccountDevice) UnsetSystemName()`

UnsetSystemName ensures that no value is present for SystemName, not even an explicit nil
### GetSystemVersion

`func (o *AccountDevice) GetSystemVersion() string`

GetSystemVersion returns the SystemVersion field if non-nil, zero value otherwise.

### GetSystemVersionOk

`func (o *AccountDevice) GetSystemVersionOk() (*string, bool)`

GetSystemVersionOk returns a tuple with the SystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemVersion

`func (o *AccountDevice) SetSystemVersion(v string)`

SetSystemVersion sets SystemVersion field to given value.


### SetSystemVersionNil

`func (o *AccountDevice) SetSystemVersionNil(b bool)`

 SetSystemVersionNil sets the value for SystemVersion to be an explicit nil

### UnsetSystemVersion
`func (o *AccountDevice) UnsetSystemVersion()`

UnsetSystemVersion ensures that no value is present for SystemVersion, not even an explicit nil
### GetCreatedAt

`func (o *AccountDevice) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountDevice) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountDevice) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *AccountDevice) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AccountDevice) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *AccountDevice) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountDevice) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountDevice) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *AccountDevice) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *AccountDevice) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


