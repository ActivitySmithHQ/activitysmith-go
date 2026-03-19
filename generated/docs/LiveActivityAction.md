# LiveActivityAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Button title displayed in the Live Activity UI. | 
**Type** | [**LiveActivityActionType**](LiveActivityActionType.md) |  | 
**Url** | **string** | HTTPS URL. For open_url it is opened in browser. For webhook it is called by ActivitySmith backend. | 
**Method** | Pointer to [**LiveActivityWebhookMethod**](LiveActivityWebhookMethod.md) | Webhook HTTP method. Used only when type&#x3D;webhook. | [optional] [default to LIVEACTIVITYWEBHOOKMETHOD_POST]
**Body** | Pointer to **map[string]interface{}** | Optional webhook payload body. Used only when type&#x3D;webhook. | [optional] 

## Methods

### NewLiveActivityAction

`func NewLiveActivityAction(title string, type_ LiveActivityActionType, url string, ) *LiveActivityAction`

NewLiveActivityAction instantiates a new LiveActivityAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityActionWithDefaults

`func NewLiveActivityActionWithDefaults() *LiveActivityAction`

NewLiveActivityActionWithDefaults instantiates a new LiveActivityAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *LiveActivityAction) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LiveActivityAction) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LiveActivityAction) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *LiveActivityAction) GetType() LiveActivityActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LiveActivityAction) GetTypeOk() (*LiveActivityActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LiveActivityAction) SetType(v LiveActivityActionType)`

SetType sets Type field to given value.


### GetUrl

`func (o *LiveActivityAction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LiveActivityAction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LiveActivityAction) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMethod

`func (o *LiveActivityAction) GetMethod() LiveActivityWebhookMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LiveActivityAction) GetMethodOk() (*LiveActivityWebhookMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LiveActivityAction) SetMethod(v LiveActivityWebhookMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LiveActivityAction) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetBody

`func (o *LiveActivityAction) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *LiveActivityAction) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *LiveActivityAction) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *LiveActivityAction) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


