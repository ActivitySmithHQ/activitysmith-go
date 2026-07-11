# AppIconBadgeCountUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Badge** | **int32** | The count to show on the ActivitySmith app icon. Send 0 to clear it. | 
**Target** | Pointer to [**ChannelTarget**](ChannelTarget.md) |  | [optional] 

## Methods

### NewAppIconBadgeCountUpdateRequest

`func NewAppIconBadgeCountUpdateRequest(badge int32, ) *AppIconBadgeCountUpdateRequest`

NewAppIconBadgeCountUpdateRequest instantiates a new AppIconBadgeCountUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppIconBadgeCountUpdateRequestWithDefaults

`func NewAppIconBadgeCountUpdateRequestWithDefaults() *AppIconBadgeCountUpdateRequest`

NewAppIconBadgeCountUpdateRequestWithDefaults instantiates a new AppIconBadgeCountUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadge

`func (o *AppIconBadgeCountUpdateRequest) GetBadge() int32`

GetBadge returns the Badge field if non-nil, zero value otherwise.

### GetBadgeOk

`func (o *AppIconBadgeCountUpdateRequest) GetBadgeOk() (*int32, bool)`

GetBadgeOk returns a tuple with the Badge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadge

`func (o *AppIconBadgeCountUpdateRequest) SetBadge(v int32)`

SetBadge sets Badge field to given value.


### GetTarget

`func (o *AppIconBadgeCountUpdateRequest) GetTarget() ChannelTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AppIconBadgeCountUpdateRequest) GetTargetOk() (*ChannelTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AppIconBadgeCountUpdateRequest) SetTarget(v ChannelTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AppIconBadgeCountUpdateRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


