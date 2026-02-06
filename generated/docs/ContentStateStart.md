# ContentStateStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Subtitle** | Pointer to **string** |  | [optional] 
**NumberOfSteps** | **int32** |  | 
**CurrentStep** | **int32** |  | 
**Type** | **string** |  | 
**Color** | Pointer to **string** | Optional. Accent color for the Live Activity. Defaults to blue. | [optional] [default to "blue"]
**StepColor** | Pointer to **string** | Optional. Overrides color for the current step. | [optional] 

## Methods

### NewContentStateStart

`func NewContentStateStart(title string, numberOfSteps int32, currentStep int32, type_ string, ) *ContentStateStart`

NewContentStateStart instantiates a new ContentStateStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentStateStartWithDefaults

`func NewContentStateStartWithDefaults() *ContentStateStart`

NewContentStateStartWithDefaults instantiates a new ContentStateStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ContentStateStart) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentStateStart) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentStateStart) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSubtitle

`func (o *ContentStateStart) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *ContentStateStart) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *ContentStateStart) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *ContentStateStart) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetNumberOfSteps

`func (o *ContentStateStart) GetNumberOfSteps() int32`

GetNumberOfSteps returns the NumberOfSteps field if non-nil, zero value otherwise.

### GetNumberOfStepsOk

`func (o *ContentStateStart) GetNumberOfStepsOk() (*int32, bool)`

GetNumberOfStepsOk returns a tuple with the NumberOfSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSteps

`func (o *ContentStateStart) SetNumberOfSteps(v int32)`

SetNumberOfSteps sets NumberOfSteps field to given value.


### GetCurrentStep

`func (o *ContentStateStart) GetCurrentStep() int32`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *ContentStateStart) GetCurrentStepOk() (*int32, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *ContentStateStart) SetCurrentStep(v int32)`

SetCurrentStep sets CurrentStep field to given value.


### GetType

`func (o *ContentStateStart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentStateStart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentStateStart) SetType(v string)`

SetType sets Type field to given value.


### GetColor

`func (o *ContentStateStart) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ContentStateStart) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ContentStateStart) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ContentStateStart) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetStepColor

`func (o *ContentStateStart) GetStepColor() string`

GetStepColor returns the StepColor field if non-nil, zero value otherwise.

### GetStepColorOk

`func (o *ContentStateStart) GetStepColorOk() (*string, bool)`

GetStepColorOk returns a tuple with the StepColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepColor

`func (o *ContentStateStart) SetStepColor(v string)`

SetStepColor sets StepColor field to given value.

### HasStepColor

`func (o *ContentStateStart) HasStepColor() bool`

HasStepColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


