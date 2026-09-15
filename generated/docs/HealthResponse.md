# HealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | 
**Service** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Checks** | [**[]HealthCheck**](HealthCheck.md) |  | 

## Methods

### NewHealthResponse

`func NewHealthResponse(ok bool, service string, timestamp time.Time, checks []HealthCheck, ) *HealthResponse`

NewHealthResponse instantiates a new HealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthResponseWithDefaults

`func NewHealthResponseWithDefaults() *HealthResponse`

NewHealthResponseWithDefaults instantiates a new HealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *HealthResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *HealthResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *HealthResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetService

`func (o *HealthResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *HealthResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *HealthResponse) SetService(v string)`

SetService sets Service field to given value.


### GetTimestamp

`func (o *HealthResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HealthResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HealthResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetChecks

`func (o *HealthResponse) GetChecks() []HealthCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *HealthResponse) GetChecksOk() (*[]HealthCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *HealthResponse) SetChecks(v []HealthCheck)`

SetChecks sets Checks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


