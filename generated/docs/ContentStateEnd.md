# ContentStateEnd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Subtitle** | Pointer to **string** |  | [optional] 
**NumberOfSteps** | Pointer to **int32** | Total number of steps. Use for type&#x3D;segmented_progress. Optional on end, and safe to change if the final workflow used more or fewer steps than originally planned. | [optional] 
**CurrentStep** | Pointer to **int32** | Current step. Use for type&#x3D;segmented_progress. | [optional] 
**Percentage** | Pointer to **float32** | Progress percentage (0–100). Use for type&#x3D;progress. Takes precedence over value/upper_limit if both are provided. | [optional] 
**Value** | Pointer to **float32** | Current progress value. Use with upper_limit for type&#x3D;progress. | [optional] 
**UpperLimit** | Pointer to **float32** | Maximum progress value. Use with value for type&#x3D;progress. | [optional] 
**Type** | Pointer to **string** | Optional. When omitted, the API uses the existing Live Activity type. | [optional] 
**Color** | Pointer to **string** | Optional. Accent color for the Live Activity. Defaults to blue. | [optional] [default to "blue"]
**StepColor** | Pointer to **string** | Optional. Overrides color for the current step. Only applies to type&#x3D;segmented_progress. | [optional] 
**AutoDismissMinutes** | Pointer to **int32** | Optional. Minutes before the ended Live Activity is dismissed. Default 3. Set 0 for immediate dismissal. iOS will dismiss ended Live Activities after ~4 hours max. | [optional] [default to 3]

## Methods

### NewContentStateEnd

`func NewContentStateEnd(title string, ) *ContentStateEnd`

NewContentStateEnd instantiates a new ContentStateEnd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentStateEndWithDefaults

`func NewContentStateEndWithDefaults() *ContentStateEnd`

NewContentStateEndWithDefaults instantiates a new ContentStateEnd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ContentStateEnd) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentStateEnd) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentStateEnd) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSubtitle

`func (o *ContentStateEnd) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *ContentStateEnd) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *ContentStateEnd) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *ContentStateEnd) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetNumberOfSteps

`func (o *ContentStateEnd) GetNumberOfSteps() int32`

GetNumberOfSteps returns the NumberOfSteps field if non-nil, zero value otherwise.

### GetNumberOfStepsOk

`func (o *ContentStateEnd) GetNumberOfStepsOk() (*int32, bool)`

GetNumberOfStepsOk returns a tuple with the NumberOfSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSteps

`func (o *ContentStateEnd) SetNumberOfSteps(v int32)`

SetNumberOfSteps sets NumberOfSteps field to given value.

### HasNumberOfSteps

`func (o *ContentStateEnd) HasNumberOfSteps() bool`

HasNumberOfSteps returns a boolean if a field has been set.

### GetCurrentStep

`func (o *ContentStateEnd) GetCurrentStep() int32`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *ContentStateEnd) GetCurrentStepOk() (*int32, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *ContentStateEnd) SetCurrentStep(v int32)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *ContentStateEnd) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### GetPercentage

`func (o *ContentStateEnd) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ContentStateEnd) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ContentStateEnd) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *ContentStateEnd) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetValue

`func (o *ContentStateEnd) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContentStateEnd) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContentStateEnd) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ContentStateEnd) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUpperLimit

`func (o *ContentStateEnd) GetUpperLimit() float32`

GetUpperLimit returns the UpperLimit field if non-nil, zero value otherwise.

### GetUpperLimitOk

`func (o *ContentStateEnd) GetUpperLimitOk() (*float32, bool)`

GetUpperLimitOk returns a tuple with the UpperLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperLimit

`func (o *ContentStateEnd) SetUpperLimit(v float32)`

SetUpperLimit sets UpperLimit field to given value.

### HasUpperLimit

`func (o *ContentStateEnd) HasUpperLimit() bool`

HasUpperLimit returns a boolean if a field has been set.

### GetType

`func (o *ContentStateEnd) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentStateEnd) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentStateEnd) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContentStateEnd) HasType() bool`

HasType returns a boolean if a field has been set.

### GetColor

`func (o *ContentStateEnd) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ContentStateEnd) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ContentStateEnd) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ContentStateEnd) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetStepColor

`func (o *ContentStateEnd) GetStepColor() string`

GetStepColor returns the StepColor field if non-nil, zero value otherwise.

### GetStepColorOk

`func (o *ContentStateEnd) GetStepColorOk() (*string, bool)`

GetStepColorOk returns a tuple with the StepColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepColor

`func (o *ContentStateEnd) SetStepColor(v string)`

SetStepColor sets StepColor field to given value.

### HasStepColor

`func (o *ContentStateEnd) HasStepColor() bool`

HasStepColor returns a boolean if a field has been set.

### GetAutoDismissMinutes

`func (o *ContentStateEnd) GetAutoDismissMinutes() int32`

GetAutoDismissMinutes returns the AutoDismissMinutes field if non-nil, zero value otherwise.

### GetAutoDismissMinutesOk

`func (o *ContentStateEnd) GetAutoDismissMinutesOk() (*int32, bool)`

GetAutoDismissMinutesOk returns a tuple with the AutoDismissMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDismissMinutes

`func (o *ContentStateEnd) SetAutoDismissMinutes(v int32)`

SetAutoDismissMinutes sets AutoDismissMinutes field to given value.

### HasAutoDismissMinutes

`func (o *ContentStateEnd) HasAutoDismissMinutes() bool`

HasAutoDismissMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


