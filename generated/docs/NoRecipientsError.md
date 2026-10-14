# NoRecipientsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Code** | Pointer to **string** | no_paired_devices: the account has no active paired iOS devices; pair an iPhone with the ActivitySmith app. no_recipients_found: devices are paired, but none matched the targeted channels. | [optional] 
**Message** | **string** |  | 
**AppStoreUrl** | Pointer to **string** | ActivitySmith App Store link. Present when code is no_paired_devices. | [optional] 
**EffectiveChannelSlugs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNoRecipientsError

`func NewNoRecipientsError(error_ string, message string, ) *NoRecipientsError`

NewNoRecipientsError instantiates a new NoRecipientsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoRecipientsErrorWithDefaults

`func NewNoRecipientsErrorWithDefaults() *NoRecipientsError`

NewNoRecipientsErrorWithDefaults instantiates a new NoRecipientsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *NoRecipientsError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NoRecipientsError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NoRecipientsError) SetError(v string)`

SetError sets Error field to given value.


### GetCode

`func (o *NoRecipientsError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NoRecipientsError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NoRecipientsError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *NoRecipientsError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *NoRecipientsError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NoRecipientsError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NoRecipientsError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAppStoreUrl

`func (o *NoRecipientsError) GetAppStoreUrl() string`

GetAppStoreUrl returns the AppStoreUrl field if non-nil, zero value otherwise.

### GetAppStoreUrlOk

`func (o *NoRecipientsError) GetAppStoreUrlOk() (*string, bool)`

GetAppStoreUrlOk returns a tuple with the AppStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreUrl

`func (o *NoRecipientsError) SetAppStoreUrl(v string)`

SetAppStoreUrl sets AppStoreUrl field to given value.

### HasAppStoreUrl

`func (o *NoRecipientsError) HasAppStoreUrl() bool`

HasAppStoreUrl returns a boolean if a field has been set.

### GetEffectiveChannelSlugs

`func (o *NoRecipientsError) GetEffectiveChannelSlugs() []string`

GetEffectiveChannelSlugs returns the EffectiveChannelSlugs field if non-nil, zero value otherwise.

### GetEffectiveChannelSlugsOk

`func (o *NoRecipientsError) GetEffectiveChannelSlugsOk() (*[]string, bool)`

GetEffectiveChannelSlugsOk returns a tuple with the EffectiveChannelSlugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveChannelSlugs

`func (o *NoRecipientsError) SetEffectiveChannelSlugs(v []string)`

SetEffectiveChannelSlugs sets EffectiveChannelSlugs field to given value.

### HasEffectiveChannelSlugs

`func (o *NoRecipientsError) HasEffectiveChannelSlugs() bool`

HasEffectiveChannelSlugs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


