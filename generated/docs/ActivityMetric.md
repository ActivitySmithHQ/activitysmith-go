# ActivityMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Value** | **float32** |  | 
**Unit** | Pointer to **string** |  | [optional] 

## Methods

### NewActivityMetric

`func NewActivityMetric(label string, value float32, ) *ActivityMetric`

NewActivityMetric instantiates a new ActivityMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityMetricWithDefaults

`func NewActivityMetricWithDefaults() *ActivityMetric`

NewActivityMetricWithDefaults instantiates a new ActivityMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ActivityMetric) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ActivityMetric) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ActivityMetric) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValue

`func (o *ActivityMetric) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ActivityMetric) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ActivityMetric) SetValue(v float32)`

SetValue sets Value field to given value.


### GetUnit

`func (o *ActivityMetric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ActivityMetric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ActivityMetric) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ActivityMetric) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


