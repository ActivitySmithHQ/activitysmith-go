# AccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** | Email of the signed-in person. Null for API keys. | 
**IsAccountOwner** | **NullableBool** | True when the signed-in person owns the account. Null for API keys. | 
**Plan** | [**AccountResponsePlan**](AccountResponsePlan.md) |  | 
**Trial** | [**AccountResponseTrial**](AccountResponseTrial.md) |  | 
**CanSend** | **bool** | False when Push Notifications, Live Activities, App Icon Badge updates, and Lock Screen Widget values are paused until the account upgrades. | 
**WhenTrialEnds** | Pointer to **string** | What stops when the trial ends. Present only while a trial is running. | [optional] 
**UpgradeUrl** | **NullableString** | Where to upgrade. Set during a trial or without an active plan, otherwise null. | 

## Methods

### NewAccountResponse

`func NewAccountResponse(email NullableString, isAccountOwner NullableBool, plan AccountResponsePlan, trial AccountResponseTrial, canSend bool, upgradeUrl NullableString, ) *AccountResponse`

NewAccountResponse instantiates a new AccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResponseWithDefaults

`func NewAccountResponseWithDefaults() *AccountResponse`

NewAccountResponseWithDefaults instantiates a new AccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AccountResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *AccountResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AccountResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetIsAccountOwner

`func (o *AccountResponse) GetIsAccountOwner() bool`

GetIsAccountOwner returns the IsAccountOwner field if non-nil, zero value otherwise.

### GetIsAccountOwnerOk

`func (o *AccountResponse) GetIsAccountOwnerOk() (*bool, bool)`

GetIsAccountOwnerOk returns a tuple with the IsAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountOwner

`func (o *AccountResponse) SetIsAccountOwner(v bool)`

SetIsAccountOwner sets IsAccountOwner field to given value.


### SetIsAccountOwnerNil

`func (o *AccountResponse) SetIsAccountOwnerNil(b bool)`

 SetIsAccountOwnerNil sets the value for IsAccountOwner to be an explicit nil

### UnsetIsAccountOwner
`func (o *AccountResponse) UnsetIsAccountOwner()`

UnsetIsAccountOwner ensures that no value is present for IsAccountOwner, not even an explicit nil
### GetPlan

`func (o *AccountResponse) GetPlan() AccountResponsePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AccountResponse) GetPlanOk() (*AccountResponsePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AccountResponse) SetPlan(v AccountResponsePlan)`

SetPlan sets Plan field to given value.


### GetTrial

`func (o *AccountResponse) GetTrial() AccountResponseTrial`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *AccountResponse) GetTrialOk() (*AccountResponseTrial, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *AccountResponse) SetTrial(v AccountResponseTrial)`

SetTrial sets Trial field to given value.


### GetCanSend

`func (o *AccountResponse) GetCanSend() bool`

GetCanSend returns the CanSend field if non-nil, zero value otherwise.

### GetCanSendOk

`func (o *AccountResponse) GetCanSendOk() (*bool, bool)`

GetCanSendOk returns a tuple with the CanSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSend

`func (o *AccountResponse) SetCanSend(v bool)`

SetCanSend sets CanSend field to given value.


### GetWhenTrialEnds

`func (o *AccountResponse) GetWhenTrialEnds() string`

GetWhenTrialEnds returns the WhenTrialEnds field if non-nil, zero value otherwise.

### GetWhenTrialEndsOk

`func (o *AccountResponse) GetWhenTrialEndsOk() (*string, bool)`

GetWhenTrialEndsOk returns a tuple with the WhenTrialEnds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenTrialEnds

`func (o *AccountResponse) SetWhenTrialEnds(v string)`

SetWhenTrialEnds sets WhenTrialEnds field to given value.

### HasWhenTrialEnds

`func (o *AccountResponse) HasWhenTrialEnds() bool`

HasWhenTrialEnds returns a boolean if a field has been set.

### GetUpgradeUrl

`func (o *AccountResponse) GetUpgradeUrl() string`

GetUpgradeUrl returns the UpgradeUrl field if non-nil, zero value otherwise.

### GetUpgradeUrlOk

`func (o *AccountResponse) GetUpgradeUrlOk() (*string, bool)`

GetUpgradeUrlOk returns a tuple with the UpgradeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeUrl

`func (o *AccountResponse) SetUpgradeUrl(v string)`

SetUpgradeUrl sets UpgradeUrl field to given value.


### SetUpgradeUrlNil

`func (o *AccountResponse) SetUpgradeUrlNil(b bool)`

 SetUpgradeUrlNil sets the value for UpgradeUrl to be an explicit nil

### UnsetUpgradeUrl
`func (o *AccountResponse) UnsetUpgradeUrl()`

UnsetUpgradeUrl ensures that no value is present for UpgradeUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


