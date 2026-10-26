# Metric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicId** | **string** |  | 
**Key** | **string** |  | 
**Label** | **string** |  | 
**Format** | **string** |  | 
**CurrencyCode** | **NullableString** |  | 
**Unit** | **NullableString** |  | 
**UnitSpacing** | **string** |  | 
**LatestValue** | **interface{}** |  | 
**LatestValueAt** | **NullableTime** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewMetric

`func NewMetric(publicId string, key string, label string, format string, currencyCode NullableString, unit NullableString, unitSpacing string, latestValue interface{}, latestValueAt NullableTime, createdAt time.Time, updatedAt time.Time, ) *Metric`

NewMetric instantiates a new Metric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricWithDefaults

`func NewMetricWithDefaults() *Metric`

NewMetricWithDefaults instantiates a new Metric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicId

`func (o *Metric) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *Metric) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *Metric) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetKey

`func (o *Metric) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Metric) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Metric) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *Metric) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Metric) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Metric) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetFormat

`func (o *Metric) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Metric) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Metric) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetCurrencyCode

`func (o *Metric) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Metric) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Metric) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### SetCurrencyCodeNil

`func (o *Metric) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *Metric) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetUnit

`func (o *Metric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Metric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Metric) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *Metric) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *Metric) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetUnitSpacing

`func (o *Metric) GetUnitSpacing() string`

GetUnitSpacing returns the UnitSpacing field if non-nil, zero value otherwise.

### GetUnitSpacingOk

`func (o *Metric) GetUnitSpacingOk() (*string, bool)`

GetUnitSpacingOk returns a tuple with the UnitSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitSpacing

`func (o *Metric) SetUnitSpacing(v string)`

SetUnitSpacing sets UnitSpacing field to given value.


### GetLatestValue

`func (o *Metric) GetLatestValue() interface{}`

GetLatestValue returns the LatestValue field if non-nil, zero value otherwise.

### GetLatestValueOk

`func (o *Metric) GetLatestValueOk() (*interface{}, bool)`

GetLatestValueOk returns a tuple with the LatestValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestValue

`func (o *Metric) SetLatestValue(v interface{})`

SetLatestValue sets LatestValue field to given value.


### SetLatestValueNil

`func (o *Metric) SetLatestValueNil(b bool)`

 SetLatestValueNil sets the value for LatestValue to be an explicit nil

### UnsetLatestValue
`func (o *Metric) UnsetLatestValue()`

UnsetLatestValue ensures that no value is present for LatestValue, not even an explicit nil
### GetLatestValueAt

`func (o *Metric) GetLatestValueAt() time.Time`

GetLatestValueAt returns the LatestValueAt field if non-nil, zero value otherwise.

### GetLatestValueAtOk

`func (o *Metric) GetLatestValueAtOk() (*time.Time, bool)`

GetLatestValueAtOk returns a tuple with the LatestValueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestValueAt

`func (o *Metric) SetLatestValueAt(v time.Time)`

SetLatestValueAt sets LatestValueAt field to given value.


### SetLatestValueAtNil

`func (o *Metric) SetLatestValueAtNil(b bool)`

 SetLatestValueAtNil sets the value for LatestValueAt to be an explicit nil

### UnsetLatestValueAt
`func (o *Metric) UnsetLatestValueAt()`

UnsetLatestValueAt ensures that no value is present for LatestValueAt, not even an explicit nil
### GetCreatedAt

`func (o *Metric) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Metric) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Metric) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Metric) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Metric) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Metric) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


