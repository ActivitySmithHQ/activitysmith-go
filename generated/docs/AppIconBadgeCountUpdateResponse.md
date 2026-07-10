# AppIconBadgeCountUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Badge** | **int32** |  | 
**DevicesNotified** | **int32** |  | 
**UsersNotified** | **int32** |  | 
**EffectiveChannelSlugs** | **[]string** |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewAppIconBadgeCountUpdateResponse

`func NewAppIconBadgeCountUpdateResponse(success bool, badge int32, devicesNotified int32, usersNotified int32, effectiveChannelSlugs []string, timestamp time.Time, ) *AppIconBadgeCountUpdateResponse`

NewAppIconBadgeCountUpdateResponse instantiates a new AppIconBadgeCountUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppIconBadgeCountUpdateResponseWithDefaults

`func NewAppIconBadgeCountUpdateResponseWithDefaults() *AppIconBadgeCountUpdateResponse`

NewAppIconBadgeCountUpdateResponseWithDefaults instantiates a new AppIconBadgeCountUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AppIconBadgeCountUpdateResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AppIconBadgeCountUpdateResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AppIconBadgeCountUpdateResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetBadge

`func (o *AppIconBadgeCountUpdateResponse) GetBadge() int32`

GetBadge returns the Badge field if non-nil, zero value otherwise.

### GetBadgeOk

`func (o *AppIconBadgeCountUpdateResponse) GetBadgeOk() (*int32, bool)`

GetBadgeOk returns a tuple with the Badge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadge

`func (o *AppIconBadgeCountUpdateResponse) SetBadge(v int32)`

SetBadge sets Badge field to given value.


### GetDevicesNotified

`func (o *AppIconBadgeCountUpdateResponse) GetDevicesNotified() int32`

GetDevicesNotified returns the DevicesNotified field if non-nil, zero value otherwise.

### GetDevicesNotifiedOk

`func (o *AppIconBadgeCountUpdateResponse) GetDevicesNotifiedOk() (*int32, bool)`

GetDevicesNotifiedOk returns a tuple with the DevicesNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesNotified

`func (o *AppIconBadgeCountUpdateResponse) SetDevicesNotified(v int32)`

SetDevicesNotified sets DevicesNotified field to given value.


### GetUsersNotified

`func (o *AppIconBadgeCountUpdateResponse) GetUsersNotified() int32`

GetUsersNotified returns the UsersNotified field if non-nil, zero value otherwise.

### GetUsersNotifiedOk

`func (o *AppIconBadgeCountUpdateResponse) GetUsersNotifiedOk() (*int32, bool)`

GetUsersNotifiedOk returns a tuple with the UsersNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersNotified

`func (o *AppIconBadgeCountUpdateResponse) SetUsersNotified(v int32)`

SetUsersNotified sets UsersNotified field to given value.


### GetEffectiveChannelSlugs

`func (o *AppIconBadgeCountUpdateResponse) GetEffectiveChannelSlugs() []string`

GetEffectiveChannelSlugs returns the EffectiveChannelSlugs field if non-nil, zero value otherwise.

### GetEffectiveChannelSlugsOk

`func (o *AppIconBadgeCountUpdateResponse) GetEffectiveChannelSlugsOk() (*[]string, bool)`

GetEffectiveChannelSlugsOk returns a tuple with the EffectiveChannelSlugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveChannelSlugs

`func (o *AppIconBadgeCountUpdateResponse) SetEffectiveChannelSlugs(v []string)`

SetEffectiveChannelSlugs sets EffectiveChannelSlugs field to given value.


### GetTimestamp

`func (o *AppIconBadgeCountUpdateResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AppIconBadgeCountUpdateResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AppIconBadgeCountUpdateResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


