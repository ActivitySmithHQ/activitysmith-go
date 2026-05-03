# WidgetMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicId** | **string** |  | 
**Key** | **string** |  | 
**Label** | **string** |  | 
**CurrencyCode** | **string** | Present when format is currency. | 
**Unit** | **string** | Present when format is unit. | 
**UnitSpacing** | [**MetricUnitSpacing**](MetricUnitSpacing.md) |  | 
**Format** | [**MetricFormat**](MetricFormat.md) |  | 
**LatestValue** | **float32** | Latest metric value. Numeric formats return a number. String metrics return text. | 
**LatestValueAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewWidgetMetric

`func NewWidgetMetric(publicId string, key string, label string, currencyCode string, unit string, unitSpacing MetricUnitSpacing, format MetricFormat, latestValue float32, latestValueAt time.Time, createdAt time.Time, updatedAt time.Time, ) *WidgetMetric`

NewWidgetMetric instantiates a new WidgetMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetMetricWithDefaults

`func NewWidgetMetricWithDefaults() *WidgetMetric`

NewWidgetMetricWithDefaults instantiates a new WidgetMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicId

`func (o *WidgetMetric) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *WidgetMetric) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *WidgetMetric) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetKey

`func (o *WidgetMetric) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WidgetMetric) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WidgetMetric) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *WidgetMetric) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WidgetMetric) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WidgetMetric) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCurrencyCode

`func (o *WidgetMetric) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *WidgetMetric) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *WidgetMetric) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetUnit

`func (o *WidgetMetric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *WidgetMetric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *WidgetMetric) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetUnitSpacing

`func (o *WidgetMetric) GetUnitSpacing() MetricUnitSpacing`

GetUnitSpacing returns the UnitSpacing field if non-nil, zero value otherwise.

### GetUnitSpacingOk

`func (o *WidgetMetric) GetUnitSpacingOk() (*MetricUnitSpacing, bool)`

GetUnitSpacingOk returns a tuple with the UnitSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitSpacing

`func (o *WidgetMetric) SetUnitSpacing(v MetricUnitSpacing)`

SetUnitSpacing sets UnitSpacing field to given value.


### GetFormat

`func (o *WidgetMetric) GetFormat() MetricFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WidgetMetric) GetFormatOk() (*MetricFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WidgetMetric) SetFormat(v MetricFormat)`

SetFormat sets Format field to given value.


### GetLatestValue

`func (o *WidgetMetric) GetLatestValue() float32`

GetLatestValue returns the LatestValue field if non-nil, zero value otherwise.

### GetLatestValueOk

`func (o *WidgetMetric) GetLatestValueOk() (*float32, bool)`

GetLatestValueOk returns a tuple with the LatestValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestValue

`func (o *WidgetMetric) SetLatestValue(v float32)`

SetLatestValue sets LatestValue field to given value.


### GetLatestValueAt

`func (o *WidgetMetric) GetLatestValueAt() time.Time`

GetLatestValueAt returns the LatestValueAt field if non-nil, zero value otherwise.

### GetLatestValueAtOk

`func (o *WidgetMetric) GetLatestValueAtOk() (*time.Time, bool)`

GetLatestValueAtOk returns a tuple with the LatestValueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestValueAt

`func (o *WidgetMetric) SetLatestValueAt(v time.Time)`

SetLatestValueAt sets LatestValueAt field to given value.


### GetCreatedAt

`func (o *WidgetMetric) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WidgetMetric) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WidgetMetric) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WidgetMetric) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WidgetMetric) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WidgetMetric) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


