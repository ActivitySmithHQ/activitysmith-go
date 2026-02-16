# AlertPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 

## Methods

### NewAlertPayload

`func NewAlertPayload() *AlertPayload`

NewAlertPayload instantiates a new AlertPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertPayloadWithDefaults

`func NewAlertPayloadWithDefaults() *AlertPayload`

NewAlertPayloadWithDefaults instantiates a new AlertPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AlertPayload) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlertPayload) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlertPayload) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AlertPayload) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBody

`func (o *AlertPayload) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AlertPayload) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AlertPayload) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *AlertPayload) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


