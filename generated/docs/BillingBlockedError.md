# BillingBlockedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Message** | **string** |  | 
**TrialPeriod** | Pointer to [**BillingBlockedErrorTrialPeriod**](BillingBlockedErrorTrialPeriod.md) |  | [optional] 
**UpgradeUrl** | **string** |  | 

## Methods

### NewBillingBlockedError

`func NewBillingBlockedError(error_ string, message string, upgradeUrl string, ) *BillingBlockedError`

NewBillingBlockedError instantiates a new BillingBlockedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingBlockedErrorWithDefaults

`func NewBillingBlockedErrorWithDefaults() *BillingBlockedError`

NewBillingBlockedErrorWithDefaults instantiates a new BillingBlockedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *BillingBlockedError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BillingBlockedError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BillingBlockedError) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *BillingBlockedError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BillingBlockedError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BillingBlockedError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTrialPeriod

`func (o *BillingBlockedError) GetTrialPeriod() BillingBlockedErrorTrialPeriod`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *BillingBlockedError) GetTrialPeriodOk() (*BillingBlockedErrorTrialPeriod, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *BillingBlockedError) SetTrialPeriod(v BillingBlockedErrorTrialPeriod)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *BillingBlockedError) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetUpgradeUrl

`func (o *BillingBlockedError) GetUpgradeUrl() string`

GetUpgradeUrl returns the UpgradeUrl field if non-nil, zero value otherwise.

### GetUpgradeUrlOk

`func (o *BillingBlockedError) GetUpgradeUrlOk() (*string, bool)`

GetUpgradeUrlOk returns a tuple with the UpgradeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeUrl

`func (o *BillingBlockedError) SetUpgradeUrl(v string)`

SetUpgradeUrl sets UpgradeUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


