# LiveActivityAlertBadge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Color** | Pointer to [**LiveActivityColor**](LiveActivityColor.md) | Optional badge color. | [optional] 

## Methods

### NewLiveActivityAlertBadge

`func NewLiveActivityAlertBadge(title string, ) *LiveActivityAlertBadge`

NewLiveActivityAlertBadge instantiates a new LiveActivityAlertBadge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityAlertBadgeWithDefaults

`func NewLiveActivityAlertBadgeWithDefaults() *LiveActivityAlertBadge`

NewLiveActivityAlertBadgeWithDefaults instantiates a new LiveActivityAlertBadge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *LiveActivityAlertBadge) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LiveActivityAlertBadge) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LiveActivityAlertBadge) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetColor

`func (o *LiveActivityAlertBadge) GetColor() LiveActivityColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *LiveActivityAlertBadge) GetColorOk() (*LiveActivityColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *LiveActivityAlertBadge) SetColor(v LiveActivityColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *LiveActivityAlertBadge) HasColor() bool`

HasColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


