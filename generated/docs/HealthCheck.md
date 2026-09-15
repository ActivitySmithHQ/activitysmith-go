# HealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Status** | **string** |  | 
**DurationMs** | **int32** |  | 

## Methods

### NewHealthCheck

`func NewHealthCheck(name string, status string, durationMs int32, ) *HealthCheck`

NewHealthCheck instantiates a new HealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckWithDefaults

`func NewHealthCheckWithDefaults() *HealthCheck`

NewHealthCheckWithDefaults instantiates a new HealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HealthCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthCheck) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *HealthCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheck) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDurationMs

`func (o *HealthCheck) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *HealthCheck) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *HealthCheck) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


