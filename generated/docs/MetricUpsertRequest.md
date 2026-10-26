# MetricUpsertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Name shown on the widget. | [optional] 
**Format** | Pointer to **string** | How widgets display the value: number (1,234), currency (requires currency_code), percent (send 99.9 to show 99.9%), unit (a number with a unit suffix, such as 120 ms), or string (text up to 64 characters). Defaults to number for a new metric. | [optional] 
**CurrencyCode** | Pointer to **string** | 3-letter ISO 4217 currency code, such as USD. Required for the currency format. | [optional] 
**Unit** | Pointer to **string** | Unit suffix for the unit format, such as ms or GB. | [optional] 
**UnitSpacing** | Pointer to **string** | For the unit format: none shows 120ms, space shows 120 ms. | [optional] 

## Methods

### NewMetricUpsertRequest

`func NewMetricUpsertRequest() *MetricUpsertRequest`

NewMetricUpsertRequest instantiates a new MetricUpsertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricUpsertRequestWithDefaults

`func NewMetricUpsertRequestWithDefaults() *MetricUpsertRequest`

NewMetricUpsertRequestWithDefaults instantiates a new MetricUpsertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *MetricUpsertRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MetricUpsertRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MetricUpsertRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MetricUpsertRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetFormat

`func (o *MetricUpsertRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MetricUpsertRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MetricUpsertRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MetricUpsertRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *MetricUpsertRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *MetricUpsertRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *MetricUpsertRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *MetricUpsertRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetUnit

`func (o *MetricUpsertRequest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricUpsertRequest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricUpsertRequest) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricUpsertRequest) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitSpacing

`func (o *MetricUpsertRequest) GetUnitSpacing() string`

GetUnitSpacing returns the UnitSpacing field if non-nil, zero value otherwise.

### GetUnitSpacingOk

`func (o *MetricUpsertRequest) GetUnitSpacingOk() (*string, bool)`

GetUnitSpacingOk returns a tuple with the UnitSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitSpacing

`func (o *MetricUpsertRequest) SetUnitSpacing(v string)`

SetUnitSpacing sets UnitSpacing field to given value.

### HasUnitSpacing

`func (o *MetricUpsertRequest) HasUnitSpacing() bool`

HasUnitSpacing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


