# LiveActivityLimitError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Message** | **string** |  | 
**Limit** | **int32** |  | 
**Active** | **int32** | Current number of active Live Activities. | 

## Methods

### NewLiveActivityLimitError

`func NewLiveActivityLimitError(error_ string, message string, limit int32, active int32, ) *LiveActivityLimitError`

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


### GetActive

`func (o *LiveActivityLimitError) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *LiveActivityLimitError) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *LiveActivityLimitError) SetActive(v int32)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


