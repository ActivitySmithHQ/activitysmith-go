# LiveActivityLimitError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Message** | **string** |  | 
**Limit** | **int32** |  | 
**BlockedDevices** | **int32** | Number of targeted devices that have reached the enforced iOS Live Activity concurrency threshold. | 

## Methods

### NewLiveActivityLimitError

`func NewLiveActivityLimitError(error_ string, message string, limit int32, blockedDevices int32, ) *LiveActivityLimitError`

NewLiveActivityLimitError instantiates a new LiveActivityLimitError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityLimitErrorWithDefaults

`func NewLiveActivityLimitErrorWithDefaults() *LiveActivityLimitError`

NewLiveActivityLimitErrorWithDefaults instantiates a new LiveActivityLimitError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *LiveActivityLimitError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LiveActivityLimitError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *LiveActivityLimitError) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *LiveActivityLimitError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LiveActivityLimitError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LiveActivityLimitError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetLimit

`func (o *LiveActivityLimitError) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LiveActivityLimitError) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LiveActivityLimitError) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetBlockedDevices

`func (o *LiveActivityLimitError) GetBlockedDevices() int32`

GetBlockedDevices returns the BlockedDevices field if non-nil, zero value otherwise.

### GetBlockedDevicesOk

`func (o *LiveActivityLimitError) GetBlockedDevicesOk() (*int32, bool)`

GetBlockedDevicesOk returns a tuple with the BlockedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDevices

`func (o *LiveActivityLimitError) SetBlockedDevices(v int32)`

SetBlockedDevices sets BlockedDevices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


