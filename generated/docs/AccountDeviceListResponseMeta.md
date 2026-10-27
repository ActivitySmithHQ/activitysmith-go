# AccountDeviceListResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Devices matching the filter. | 
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**HasMore** | **bool** |  | 
**NextOffset** | **NullableInt32** | Offset for the next page, or null on the last page. | 
**IncludeInactive** | **bool** |  | 

## Methods

### NewAccountDeviceListResponseMeta

`func NewAccountDeviceListResponseMeta(total int32, limit int32, offset int32, hasMore bool, nextOffset NullableInt32, includeInactive bool, ) *AccountDeviceListResponseMeta`

NewAccountDeviceListResponseMeta instantiates a new AccountDeviceListResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDeviceListResponseMetaWithDefaults

`func NewAccountDeviceListResponseMetaWithDefaults() *AccountDeviceListResponseMeta`

NewAccountDeviceListResponseMetaWithDefaults instantiates a new AccountDeviceListResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AccountDeviceListResponseMeta) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccountDeviceListResponseMeta) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccountDeviceListResponseMeta) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetLimit

`func (o *AccountDeviceListResponseMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AccountDeviceListResponseMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AccountDeviceListResponseMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *AccountDeviceListResponseMeta) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AccountDeviceListResponseMeta) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AccountDeviceListResponseMeta) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetHasMore

`func (o *AccountDeviceListResponseMeta) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *AccountDeviceListResponseMeta) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *AccountDeviceListResponseMeta) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNextOffset

`func (o *AccountDeviceListResponseMeta) GetNextOffset() int32`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *AccountDeviceListResponseMeta) GetNextOffsetOk() (*int32, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *AccountDeviceListResponseMeta) SetNextOffset(v int32)`

SetNextOffset sets NextOffset field to given value.


### SetNextOffsetNil

`func (o *AccountDeviceListResponseMeta) SetNextOffsetNil(b bool)`

 SetNextOffsetNil sets the value for NextOffset to be an explicit nil

### UnsetNextOffset
`func (o *AccountDeviceListResponseMeta) UnsetNextOffset()`

UnsetNextOffset ensures that no value is present for NextOffset, not even an explicit nil
### GetIncludeInactive

`func (o *AccountDeviceListResponseMeta) GetIncludeInactive() bool`

GetIncludeInactive returns the IncludeInactive field if non-nil, zero value otherwise.

### GetIncludeInactiveOk

`func (o *AccountDeviceListResponseMeta) GetIncludeInactiveOk() (*bool, bool)`

GetIncludeInactiveOk returns a tuple with the IncludeInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInactive

`func (o *AccountDeviceListResponseMeta) SetIncludeInactive(v bool)`

SetIncludeInactive sets IncludeInactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


