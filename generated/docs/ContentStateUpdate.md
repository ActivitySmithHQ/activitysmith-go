# ContentStateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Subtitle** | Pointer to **string** |  | [optional] 
**NumberOfSteps** | Pointer to **int32** | Total number of steps. Use for type&#x3D;segmented_progress. Optional on update, and safe to change if the workflow gains or loses steps. | [optional] 
**CurrentStep** | Pointer to **int32** | Current step. Use for type&#x3D;segmented_progress. | [optional] 
**Percentage** | Pointer to **float32** | Progress percentage (0–100). Use for type&#x3D;progress. Takes precedence over value/upper_limit if both are provided. | [optional] 
**Value** | Pointer to **float32** | Current progress value. Use with upper_limit for type&#x3D;progress. | [optional] 
**UpperLimit** | Pointer to **float32** | Maximum progress value. Use with value for type&#x3D;progress. | [optional] 
**Metrics** | Pointer to [**[]ActivityMetric**](ActivityMetric.md) | Use for type&#x3D;metrics. | [optional] 
**Type** | Pointer to **string** | Optional. When omitted, the API uses the existing Live Activity type. | [optional] 
**Color** | Pointer to **string** | Optional. Accent color for the Live Activity. Defaults to blue. | [optional] [default to "blue"]
**StepColor** | Pointer to **string** | Optional. Overrides color for the current step. Only applies to type&#x3D;segmented_progress. | [optional] 
**StepColors** | Pointer to **[]string** | Optional. Colors for completed steps. When used with segmented_progress, the array length should match current_step. | [optional] 

## Methods

### NewContentStateUpdate

`func NewContentStateUpdate(title string, ) *ContentStateUpdate`

NewContentStateUpdate instantiates a new ContentStateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentStateUpdateWithDefaults

`func NewContentStateUpdateWithDefaults() *ContentStateUpdate`

NewContentStateUpdateWithDefaults instantiates a new ContentStateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ContentStateUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentStateUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentStateUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSubtitle

`func (o *ContentStateUpdate) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *ContentStateUpdate) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *ContentStateUpdate) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *ContentStateUpdate) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetNumberOfSteps

`func (o *ContentStateUpdate) GetNumberOfSteps() int32`

GetNumberOfSteps returns the NumberOfSteps field if non-nil, zero value otherwise.

### GetNumberOfStepsOk

`func (o *ContentStateUpdate) GetNumberOfStepsOk() (*int32, bool)`

GetNumberOfStepsOk returns a tuple with the NumberOfSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSteps

`func (o *ContentStateUpdate) SetNumberOfSteps(v int32)`

SetNumberOfSteps sets NumberOfSteps field to given value.

### HasNumberOfSteps

`func (o *ContentStateUpdate) HasNumberOfSteps() bool`

HasNumberOfSteps returns a boolean if a field has been set.

### GetCurrentStep

`func (o *ContentStateUpdate) GetCurrentStep() int32`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *ContentStateUpdate) GetCurrentStepOk() (*int32, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *ContentStateUpdate) SetCurrentStep(v int32)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *ContentStateUpdate) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### GetPercentage

`func (o *ContentStateUpdate) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ContentStateUpdate) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ContentStateUpdate) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *ContentStateUpdate) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetValue

`func (o *ContentStateUpdate) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContentStateUpdate) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContentStateUpdate) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ContentStateUpdate) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUpperLimit

`func (o *ContentStateUpdate) GetUpperLimit() float32`

GetUpperLimit returns the UpperLimit field if non-nil, zero value otherwise.

### GetUpperLimitOk

`func (o *ContentStateUpdate) GetUpperLimitOk() (*float32, bool)`

GetUpperLimitOk returns a tuple with the UpperLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperLimit

`func (o *ContentStateUpdate) SetUpperLimit(v float32)`

SetUpperLimit sets UpperLimit field to given value.

### HasUpperLimit

`func (o *ContentStateUpdate) HasUpperLimit() bool`

HasUpperLimit returns a boolean if a field has been set.

### GetMetrics

`func (o *ContentStateUpdate) GetMetrics() []ActivityMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ContentStateUpdate) GetMetricsOk() (*[]ActivityMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ContentStateUpdate) SetMetrics(v []ActivityMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ContentStateUpdate) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetType

`func (o *ContentStateUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentStateUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentStateUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContentStateUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetColor

`func (o *ContentStateUpdate) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ContentStateUpdate) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ContentStateUpdate) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ContentStateUpdate) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetStepColor

`func (o *ContentStateUpdate) GetStepColor() string`

GetStepColor returns the StepColor field if non-nil, zero value otherwise.

### GetStepColorOk

`func (o *ContentStateUpdate) GetStepColorOk() (*string, bool)`

GetStepColorOk returns a tuple with the StepColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepColor

`func (o *ContentStateUpdate) SetStepColor(v string)`

SetStepColor sets StepColor field to given value.

### HasStepColor

`func (o *ContentStateUpdate) HasStepColor() bool`

HasStepColor returns a boolean if a field has been set.

### GetStepColors

`func (o *ContentStateUpdate) GetStepColors() []string`

GetStepColors returns the StepColors field if non-nil, zero value otherwise.

### GetStepColorsOk

`func (o *ContentStateUpdate) GetStepColorsOk() (*[]string, bool)`

GetStepColorsOk returns a tuple with the StepColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepColors

`func (o *ContentStateUpdate) SetStepColors(v []string)`

SetStepColors sets StepColors field to given value.

### HasStepColors

`func (o *ContentStateUpdate) HasStepColors() bool`

HasStepColors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


