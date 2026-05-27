# LiveActivityAlertIcon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Apple SF Symbol name. | 
**Color** | Pointer to [**LiveActivityColor**](LiveActivityColor.md) | Optional icon color. | [optional] 

## Methods

### NewLiveActivityAlertIcon

`func NewLiveActivityAlertIcon(symbol string, ) *LiveActivityAlertIcon`

NewLiveActivityAlertIcon instantiates a new LiveActivityAlertIcon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityAlertIconWithDefaults

`func NewLiveActivityAlertIconWithDefaults() *LiveActivityAlertIcon`

NewLiveActivityAlertIconWithDefaults instantiates a new LiveActivityAlertIcon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *LiveActivityAlertIcon) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *LiveActivityAlertIcon) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *LiveActivityAlertIcon) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetColor

`func (o *LiveActivityAlertIcon) GetColor() LiveActivityColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *LiveActivityAlertIcon) GetColorOk() (*LiveActivityColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *LiveActivityAlertIcon) SetColor(v LiveActivityColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *LiveActivityAlertIcon) HasColor() bool`

HasColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


