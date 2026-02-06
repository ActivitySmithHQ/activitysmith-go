# ContentStateEnd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Subtitle** | Pointer to **string** |  | [optional] 
**NumberOfSteps** | Pointer to **int32** |  | [optional] 
**CurrentStep** | **int32** |  | 
**Color** | Pointer to **string** | Optional. Accent color for the Live Activity. Defaults to blue. | [optional] [default to "blue"]
**StepColor** | Pointer to **string** | Optional. Overrides color for the current step. | [optional] 
**AutoDismissMinutes** | Pointer to **int32** | Optional. Minutes before the ended Live Activity is dismissed. Default 3. Set 0 for immediate dismissal. iOS will dismiss ended Live Activities after ~4 hours max. | [optional] [default to 3]

## Methods

### NewContentStateEnd

`func NewContentStateEnd(title string, currentStep int32, ) *ContentStateEnd`

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


