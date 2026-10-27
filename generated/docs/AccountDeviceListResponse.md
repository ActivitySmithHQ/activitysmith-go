# AccountDeviceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | [**[]AccountDevice**](AccountDevice.md) |  | 
**Meta** | [**AccountDeviceListResponseMeta**](AccountDeviceListResponseMeta.md) |  | 

## Methods

### NewAccountDeviceListResponse

`func NewAccountDeviceListResponse(devices []AccountDevice, meta AccountDeviceListResponseMeta, ) *AccountDeviceListResponse`

NewAccountDeviceListResponse instantiates a new AccountDeviceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDeviceListResponseWithDefaults

`func NewAccountDeviceListResponseWithDefaults() *AccountDeviceListResponse`

NewAccountDeviceListResponseWithDefaults instantiates a new AccountDeviceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *AccountDeviceListResponse) GetDevices() []AccountDevice`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *AccountDeviceListResponse) GetDevicesOk() (*[]AccountDevice, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *AccountDeviceListResponse) SetDevices(v []AccountDevice)`

SetDevices sets Devices field to given value.


### GetMeta

`func (o *AccountDeviceListResponse) GetMeta() AccountDeviceListResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AccountDeviceListResponse) GetMetaOk() (*AccountDeviceListResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AccountDeviceListResponse) SetMeta(v AccountDeviceListResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


