# AppIconBadgeCountUpdateError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Code** | **string** |  | 
**Message** | **string** |  | 
**Badge** | **int32** |  | 
**DevicesTargeted** | Pointer to **int32** |  | [optional] 
**DevicesUpdated** | **int32** |  | 
**UsersUpdated** | Pointer to **int32** |  | [optional] 
**DevicesNotified** | Pointer to **int32** | Deprecated compatibility alias for devices_updated. | [optional] 
**EffectiveChannelSlugs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAppIconBadgeCountUpdateError

`func NewAppIconBadgeCountUpdateError(error_ string, code string, message string, badge int32, devicesUpdated int32, ) *AppIconBadgeCountUpdateError`

NewAppIconBadgeCountUpdateError instantiates a new AppIconBadgeCountUpdateError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppIconBadgeCountUpdateErrorWithDefaults

`func NewAppIconBadgeCountUpdateErrorWithDefaults() *AppIconBadgeCountUpdateError`

NewAppIconBadgeCountUpdateErrorWithDefaults instantiates a new AppIconBadgeCountUpdateError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *AppIconBadgeCountUpdateError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AppIconBadgeCountUpdateError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AppIconBadgeCountUpdateError) SetError(v string)`

SetError sets Error field to given value.


### GetCode

`func (o *AppIconBadgeCountUpdateError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AppIconBadgeCountUpdateError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AppIconBadgeCountUpdateError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *AppIconBadgeCountUpdateError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AppIconBadgeCountUpdateError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AppIconBadgeCountUpdateError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetBadge

`func (o *AppIconBadgeCountUpdateError) GetBadge() int32`

GetBadge returns the Badge field if non-nil, zero value otherwise.

### GetBadgeOk

`func (o *AppIconBadgeCountUpdateError) GetBadgeOk() (*int32, bool)`

GetBadgeOk returns a tuple with the Badge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadge

`func (o *AppIconBadgeCountUpdateError) SetBadge(v int32)`

SetBadge sets Badge field to given value.


### GetDevicesTargeted

`func (o *AppIconBadgeCountUpdateError) GetDevicesTargeted() int32`

GetDevicesTargeted returns the DevicesTargeted field if non-nil, zero value otherwise.

### GetDevicesTargetedOk

`func (o *AppIconBadgeCountUpdateError) GetDevicesTargetedOk() (*int32, bool)`

GetDevicesTargetedOk returns a tuple with the DevicesTargeted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesTargeted

`func (o *AppIconBadgeCountUpdateError) SetDevicesTargeted(v int32)`

SetDevicesTargeted sets DevicesTargeted field to given value.

### HasDevicesTargeted

`func (o *AppIconBadgeCountUpdateError) HasDevicesTargeted() bool`

HasDevicesTargeted returns a boolean if a field has been set.

### GetDevicesUpdated

`func (o *AppIconBadgeCountUpdateError) GetDevicesUpdated() int32`

GetDevicesUpdated returns the DevicesUpdated field if non-nil, zero value otherwise.

### GetDevicesUpdatedOk

`func (o *AppIconBadgeCountUpdateError) GetDevicesUpdatedOk() (*int32, bool)`

GetDevicesUpdatedOk returns a tuple with the DevicesUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesUpdated

`func (o *AppIconBadgeCountUpdateError) SetDevicesUpdated(v int32)`

SetDevicesUpdated sets DevicesUpdated field to given value.


### GetUsersUpdated

`func (o *AppIconBadgeCountUpdateError) GetUsersUpdated() int32`

GetUsersUpdated returns the UsersUpdated field if non-nil, zero value otherwise.

### GetUsersUpdatedOk

`func (o *AppIconBadgeCountUpdateError) GetUsersUpdatedOk() (*int32, bool)`

GetUsersUpdatedOk returns a tuple with the UsersUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersUpdated

`func (o *AppIconBadgeCountUpdateError) SetUsersUpdated(v int32)`

SetUsersUpdated sets UsersUpdated field to given value.

### HasUsersUpdated

`func (o *AppIconBadgeCountUpdateError) HasUsersUpdated() bool`

HasUsersUpdated returns a boolean if a field has been set.

### GetDevicesNotified

`func (o *AppIconBadgeCountUpdateError) GetDevicesNotified() int32`

GetDevicesNotified returns the DevicesNotified field if non-nil, zero value otherwise.

### GetDevicesNotifiedOk

`func (o *AppIconBadgeCountUpdateError) GetDevicesNotifiedOk() (*int32, bool)`

GetDevicesNotifiedOk returns a tuple with the DevicesNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesNotified

`func (o *AppIconBadgeCountUpdateError) SetDevicesNotified(v int32)`

SetDevicesNotified sets DevicesNotified field to given value.

### HasDevicesNotified

`func (o *AppIconBadgeCountUpdateError) HasDevicesNotified() bool`

HasDevicesNotified returns a boolean if a field has been set.

### GetEffectiveChannelSlugs

`func (o *AppIconBadgeCountUpdateError) GetEffectiveChannelSlugs() []string`

GetEffectiveChannelSlugs returns the EffectiveChannelSlugs field if non-nil, zero value otherwise.

### GetEffectiveChannelSlugsOk

`func (o *AppIconBadgeCountUpdateError) GetEffectiveChannelSlugsOk() (*[]string, bool)`

GetEffectiveChannelSlugsOk returns a tuple with the EffectiveChannelSlugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveChannelSlugs

`func (o *AppIconBadgeCountUpdateError) SetEffectiveChannelSlugs(v []string)`

SetEffectiveChannelSlugs sets EffectiveChannelSlugs field to given value.

### HasEffectiveChannelSlugs

`func (o *AppIconBadgeCountUpdateError) HasEffectiveChannelSlugs() bool`

HasEffectiveChannelSlugs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


