# ContentStateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Subtitle** | Pointer to **string** |  | [optional] 
**NumberOfSteps** | Pointer to **int32** |  | [optional] 
**CurrentStep** | **int32** |  | 
**Color** | Pointer to **string** | Optional. Accent color for the Live Activity. Defaults to blue. | [optional] [default to "blue"]
**StepColor** | Pointer to **string** | Optional. Overrides color for the current step. | [optional] 

## Methods

### NewContentStateUpdate

`func NewContentStateUpdate(title string, currentStep int32, ) *ContentStateUpdate`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


