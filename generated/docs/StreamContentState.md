# StreamContentState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Subtitle** | Pointer to **string** |  | [optional] 
**NumberOfSteps** | Pointer to **int32** | Use for segmented_progress. | [optional] 
**CurrentStep** | Pointer to **int32** | Use for segmented_progress. | [optional] 
**Percentage** | Pointer to **float32** | Use for progress. Takes precedence over value/upper_limit if both are provided. | [optional] 
**Value** | Pointer to **float32** | Current progress value. Use with upper_limit for progress. | [optional] 
**UpperLimit** | Pointer to **float32** | Maximum progress value. Use with value for progress. | [optional] 
**Type** | Pointer to **string** | Required on the first PUT or whenever the stream cannot infer the current activity type. | [optional] 
**Color** | Pointer to **string** | Optional. Accent color for the Live Activity. Defaults to blue. | [optional] [default to "blue"]
**StepColor** | Pointer to **string** | Optional. Overrides color for the current step. Only applies to segmented_progress. | [optional] 
**StepColors** | Pointer to **[]string** | Optional. Colors for completed steps. When used with segmented_progress, the array length should match current_step. | [optional] 
**Metrics** | Pointer to [**[]ActivityMetric**](ActivityMetric.md) | Use for metrics activities. | [optional] 
**AutoDismissSeconds** | Pointer to **int32** | Optional. Seconds before the ended Live Activity is dismissed. | [optional] 
**AutoDismissMinutes** | Pointer to **int32** | Optional. Minutes before the ended Live Activity is dismissed. | [optional] 

## Methods

### NewStreamContentState

`func NewStreamContentState(title string, ) *StreamContentState`

NewStreamContentState instantiates a new StreamContentState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamContentStateWithDefaults

`func NewStreamContentStateWithDefaults() *StreamContentState`

NewStreamContentStateWithDefaults instantiates a new StreamContentState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *StreamContentState) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StreamContentState) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StreamContentState) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSubtitle

`func (o *StreamContentState) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *StreamContentState) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *StreamContentState) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *StreamContentState) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetNumberOfSteps

`func (o *StreamContentState) GetNumberOfSteps() int32`

GetNumberOfSteps returns the NumberOfSteps field if non-nil, zero value otherwise.

### GetNumberOfStepsOk

`func (o *StreamContentState) GetNumberOfStepsOk() (*int32, bool)`

GetNumberOfStepsOk returns a tuple with the NumberOfSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSteps

`func (o *StreamContentState) SetNumberOfSteps(v int32)`

SetNumberOfSteps sets NumberOfSteps field to given value.

### HasNumberOfSteps

`func (o *StreamContentState) HasNumberOfSteps() bool`

HasNumberOfSteps returns a boolean if a field has been set.

### GetCurrentStep

`func (o *StreamContentState) GetCurrentStep() int32`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *StreamContentState) GetCurrentStepOk() (*int32, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *StreamContentState) SetCurrentStep(v int32)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *StreamContentState) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### GetPercentage

`func (o *StreamContentState) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *StreamContentState) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *StreamContentState) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *StreamContentState) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetValue

`func (o *StreamContentState) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StreamContentState) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StreamContentState) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *StreamContentState) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUpperLimit

`func (o *StreamContentState) GetUpperLimit() float32`

GetUpperLimit returns the UpperLimit field if non-nil, zero value otherwise.

### GetUpperLimitOk

`func (o *StreamContentState) GetUpperLimitOk() (*float32, bool)`

GetUpperLimitOk returns a tuple with the UpperLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperLimit

`func (o *StreamContentState) SetUpperLimit(v float32)`

SetUpperLimit sets UpperLimit field to given value.

### HasUpperLimit

`func (o *StreamContentState) HasUpperLimit() bool`

HasUpperLimit returns a boolean if a field has been set.

### GetType

`func (o *StreamContentState) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamContentState) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamContentState) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StreamContentState) HasType() bool`

HasType returns a boolean if a field has been set.

### GetColor

`func (o *StreamContentState) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *StreamContentState) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *StreamContentState) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *StreamContentState) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetStepColor

`func (o *StreamContentState) GetStepColor() string`

GetStepColor returns the StepColor field if non-nil, zero value otherwise.

### GetStepColorOk

`func (o *StreamContentState) GetStepColorOk() (*string, bool)`

GetStepColorOk returns a tuple with the StepColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepColor

`func (o *StreamContentState) SetStepColor(v string)`

SetStepColor sets StepColor field to given value.

### HasStepColor

`func (o *StreamContentState) HasStepColor() bool`

HasStepColor returns a boolean if a field has been set.

### GetStepColors

`func (o *StreamContentState) GetStepColors() []string`

GetStepColors returns the StepColors field if non-nil, zero value otherwise.

### GetStepColorsOk

`func (o *StreamContentState) GetStepColorsOk() (*[]string, bool)`

GetStepColorsOk returns a tuple with the StepColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepColors

`func (o *StreamContentState) SetStepColors(v []string)`

SetStepColors sets StepColors field to given value.

### HasStepColors

`func (o *StreamContentState) HasStepColors() bool`

HasStepColors returns a boolean if a field has been set.

### GetMetrics

`func (o *StreamContentState) GetMetrics() []ActivityMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *StreamContentState) GetMetricsOk() (*[]ActivityMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *StreamContentState) SetMetrics(v []ActivityMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *StreamContentState) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetAutoDismissSeconds

`func (o *StreamContentState) GetAutoDismissSeconds() int32`

GetAutoDismissSeconds returns the AutoDismissSeconds field if non-nil, zero value otherwise.

### GetAutoDismissSecondsOk

`func (o *StreamContentState) GetAutoDismissSecondsOk() (*int32, bool)`

GetAutoDismissSecondsOk returns a tuple with the AutoDismissSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDismissSeconds

`func (o *StreamContentState) SetAutoDismissSeconds(v int32)`

SetAutoDismissSeconds sets AutoDismissSeconds field to given value.

### HasAutoDismissSeconds

`func (o *StreamContentState) HasAutoDismissSeconds() bool`

HasAutoDismissSeconds returns a boolean if a field has been set.

### GetAutoDismissMinutes

`func (o *StreamContentState) GetAutoDismissMinutes() int32`

GetAutoDismissMinutes returns the AutoDismissMinutes field if non-nil, zero value otherwise.

### GetAutoDismissMinutesOk

`func (o *StreamContentState) GetAutoDismissMinutesOk() (*int32, bool)`

GetAutoDismissMinutesOk returns a tuple with the AutoDismissMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDismissMinutes

`func (o *StreamContentState) SetAutoDismissMinutes(v int32)`

SetAutoDismissMinutes sets AutoDismissMinutes field to given value.

### HasAutoDismissMinutes

`func (o *StreamContentState) HasAutoDismissMinutes() bool`

HasAutoDismissMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


