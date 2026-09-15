# UpdateAppIconBadgeCount422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Message** | **string** |  | 
**EffectiveChannelSlugs** | Pointer to **[]string** |  | [optional] 
**Code** | **string** |  | 
**Badge** | **int32** |  | 
**DevicesTargeted** | Pointer to **int32** |  | [optional] 
**DevicesUpdated** | **int32** |  | 
**UsersUpdated** | Pointer to **int32** |  | [optional] 
**DevicesNotified** | Pointer to **int32** | Deprecated compatibility alias for devices_updated. | [optional] 

## Methods

### NewUpdateAppIconBadgeCount422Response

`func NewUpdateAppIconBadgeCount422Response(error_ string, message string, code string, badge int32, devicesUpdated int32, ) *UpdateAppIconBadgeCount422Response`

NewUpdateAppIconBadgeCount422Response instantiates a new UpdateAppIconBadgeCount422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppIconBadgeCount422ResponseWithDefaults

`func NewUpdateAppIconBadgeCount422ResponseWithDefaults() *UpdateAppIconBadgeCount422Response`

NewUpdateAppIconBadgeCount422ResponseWithDefaults instantiates a new UpdateAppIconBadgeCount422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *UpdateAppIconBadgeCount422Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UpdateAppIconBadgeCount422Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UpdateAppIconBadgeCount422Response) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *UpdateAppIconBadgeCount422Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateAppIconBadgeCount422Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateAppIconBadgeCount422Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetEffectiveChannelSlugs

`func (o *UpdateAppIconBadgeCount422Response) GetEffectiveChannelSlugs() []string`

GetEffectiveChannelSlugs returns the EffectiveChannelSlugs field if non-nil, zero value otherwise.

### GetEffectiveChannelSlugsOk

`func (o *UpdateAppIconBadgeCount422Response) GetEffectiveChannelSlugsOk() (*[]string, bool)`

GetEffectiveChannelSlugsOk returns a tuple with the EffectiveChannelSlugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveChannelSlugs

`func (o *UpdateAppIconBadgeCount422Response) SetEffectiveChannelSlugs(v []string)`

SetEffectiveChannelSlugs sets EffectiveChannelSlugs field to given value.

### HasEffectiveChannelSlugs

`func (o *UpdateAppIconBadgeCount422Response) HasEffectiveChannelSlugs() bool`

HasEffectiveChannelSlugs returns a boolean if a field has been set.

### GetCode

`func (o *UpdateAppIconBadgeCount422Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateAppIconBadgeCount422Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateAppIconBadgeCount422Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetBadge

`func (o *UpdateAppIconBadgeCount422Response) GetBadge() int32`

GetBadge returns the Badge field if non-nil, zero value otherwise.

### GetBadgeOk

`func (o *UpdateAppIconBadgeCount422Response) GetBadgeOk() (*int32, bool)`

GetBadgeOk returns a tuple with the Badge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadge

`func (o *UpdateAppIconBadgeCount422Response) SetBadge(v int32)`

SetBadge sets Badge field to given value.


### GetDevicesTargeted

`func (o *UpdateAppIconBadgeCount422Response) GetDevicesTargeted() int32`

GetDevicesTargeted returns the DevicesTargeted field if non-nil, zero value otherwise.

### GetDevicesTargetedOk

`func (o *UpdateAppIconBadgeCount422Response) GetDevicesTargetedOk() (*int32, bool)`

GetDevicesTargetedOk returns a tuple with the DevicesTargeted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesTargeted

`func (o *UpdateAppIconBadgeCount422Response) SetDevicesTargeted(v int32)`

SetDevicesTargeted sets DevicesTargeted field to given value.

### HasDevicesTargeted

`func (o *UpdateAppIconBadgeCount422Response) HasDevicesTargeted() bool`

HasDevicesTargeted returns a boolean if a field has been set.

### GetDevicesUpdated

`func (o *UpdateAppIconBadgeCount422Response) GetDevicesUpdated() int32`

GetDevicesUpdated returns the DevicesUpdated field if non-nil, zero value otherwise.

### GetDevicesUpdatedOk

`func (o *UpdateAppIconBadgeCount422Response) GetDevicesUpdatedOk() (*int32, bool)`

GetDevicesUpdatedOk returns a tuple with the DevicesUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesUpdated

`func (o *UpdateAppIconBadgeCount422Response) SetDevicesUpdated(v int32)`

SetDevicesUpdated sets DevicesUpdated field to given value.


### GetUsersUpdated

`func (o *UpdateAppIconBadgeCount422Response) GetUsersUpdated() int32`

GetUsersUpdated returns the UsersUpdated field if non-nil, zero value otherwise.

### GetUsersUpdatedOk

`func (o *UpdateAppIconBadgeCount422Response) GetUsersUpdatedOk() (*int32, bool)`

GetUsersUpdatedOk returns a tuple with the UsersUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersUpdated

`func (o *UpdateAppIconBadgeCount422Response) SetUsersUpdated(v int32)`

SetUsersUpdated sets UsersUpdated field to given value.

### HasUsersUpdated

`func (o *UpdateAppIconBadgeCount422Response) HasUsersUpdated() bool`

HasUsersUpdated returns a boolean if a field has been set.

### GetDevicesNotified

`func (o *UpdateAppIconBadgeCount422Response) GetDevicesNotified() int32`

GetDevicesNotified returns the DevicesNotified field if non-nil, zero value otherwise.

### GetDevicesNotifiedOk

`func (o *UpdateAppIconBadgeCount422Response) GetDevicesNotifiedOk() (*int32, bool)`

GetDevicesNotifiedOk returns a tuple with the DevicesNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesNotified

`func (o *UpdateAppIconBadgeCount422Response) SetDevicesNotified(v int32)`

SetDevicesNotified sets DevicesNotified field to given value.

### HasDevicesNotified

`func (o *UpdateAppIconBadgeCount422Response) HasDevicesNotified() bool`

HasDevicesNotified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


