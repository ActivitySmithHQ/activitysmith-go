# PushNotificationAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Button title displayed in iOS expanded notification UI. | 
**Type** | [**PushNotificationActionType**](PushNotificationActionType.md) |  | 
**Url** | **string** | HTTPS URL. For open_url it is opened in browser. For webhook it is called by ActivitySmith backend. | 
**Method** | Pointer to [**PushNotificationWebhookMethod**](PushNotificationWebhookMethod.md) | Webhook HTTP method. Used only when type&#x3D;webhook. | [optional] [default to PUSHNOTIFICATIONWEBHOOKMETHOD_POST]
**Body** | Pointer to **map[string]interface{}** | Optional webhook payload body. Used only when type&#x3D;webhook. | [optional] 

## Methods

### NewPushNotificationAction

`func NewPushNotificationAction(title string, type_ PushNotificationActionType, url string, ) *PushNotificationAction`

NewPushNotificationAction instantiates a new PushNotificationAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushNotificationActionWithDefaults

`func NewPushNotificationActionWithDefaults() *PushNotificationAction`

NewPushNotificationActionWithDefaults instantiates a new PushNotificationAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PushNotificationAction) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PushNotificationAction) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PushNotificationAction) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *PushNotificationAction) GetType() PushNotificationActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PushNotificationAction) GetTypeOk() (*PushNotificationActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PushNotificationAction) SetType(v PushNotificationActionType)`

SetType sets Type field to given value.


### GetUrl

`func (o *PushNotificationAction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PushNotificationAction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PushNotificationAction) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMethod

`func (o *PushNotificationAction) GetMethod() PushNotificationWebhookMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PushNotificationAction) GetMethodOk() (*PushNotificationWebhookMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PushNotificationAction) SetMethod(v PushNotificationWebhookMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PushNotificationAction) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetBody

`func (o *PushNotificationAction) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PushNotificationAction) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PushNotificationAction) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *PushNotificationAction) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


