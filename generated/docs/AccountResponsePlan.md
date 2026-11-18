# AccountResponsePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | trial: the free trial. subscription: a monthly or yearly plan. lifetime: a Lifetime Membership. grandfathered: an early account with free access. inactive: no trial and no active plan. | 
**PlanId** | **NullableString** | Plan ID, such as pro. Null during a trial or without a plan. | 
**BillingInterval** | **NullableString** | month or year for subscriptions, otherwise null. | 
**RenewsAt** | **NullableTime** | Next payment date of a subscription, otherwise null. | 
**CancelAtPeriodEnd** | **bool** | True when the subscription ends at the end of the current period. | 

## Methods

### NewAccountResponsePlan

`func NewAccountResponsePlan(status string, planId NullableString, billingInterval NullableString, renewsAt NullableTime, cancelAtPeriodEnd bool, ) *AccountResponsePlan`

NewAccountResponsePlan instantiates a new AccountResponsePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResponsePlanWithDefaults

`func NewAccountResponsePlanWithDefaults() *AccountResponsePlan`

NewAccountResponsePlanWithDefaults instantiates a new AccountResponsePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AccountResponsePlan) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountResponsePlan) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountResponsePlan) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlanId

`func (o *AccountResponsePlan) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *AccountResponsePlan) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *AccountResponsePlan) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### SetPlanIdNil

`func (o *AccountResponsePlan) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *AccountResponsePlan) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetBillingInterval

`func (o *AccountResponsePlan) GetBillingInterval() string`

GetBillingInterval returns the BillingInterval field if non-nil, zero value otherwise.

### GetBillingIntervalOk

`func (o *AccountResponsePlan) GetBillingIntervalOk() (*string, bool)`

GetBillingIntervalOk returns a tuple with the BillingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInterval

`func (o *AccountResponsePlan) SetBillingInterval(v string)`

SetBillingInterval sets BillingInterval field to given value.


### SetBillingIntervalNil

`func (o *AccountResponsePlan) SetBillingIntervalNil(b bool)`

 SetBillingIntervalNil sets the value for BillingInterval to be an explicit nil

### UnsetBillingInterval
`func (o *AccountResponsePlan) UnsetBillingInterval()`

UnsetBillingInterval ensures that no value is present for BillingInterval, not even an explicit nil
### GetRenewsAt

`func (o *AccountResponsePlan) GetRenewsAt() time.Time`

GetRenewsAt returns the RenewsAt field if non-nil, zero value otherwise.

### GetRenewsAtOk

`func (o *AccountResponsePlan) GetRenewsAtOk() (*time.Time, bool)`

GetRenewsAtOk returns a tuple with the RenewsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewsAt

`func (o *AccountResponsePlan) SetRenewsAt(v time.Time)`

SetRenewsAt sets RenewsAt field to given value.


### SetRenewsAtNil

`func (o *AccountResponsePlan) SetRenewsAtNil(b bool)`

 SetRenewsAtNil sets the value for RenewsAt to be an explicit nil

### UnsetRenewsAt
`func (o *AccountResponsePlan) UnsetRenewsAt()`

UnsetRenewsAt ensures that no value is present for RenewsAt, not even an explicit nil
### GetCancelAtPeriodEnd

`func (o *AccountResponsePlan) GetCancelAtPeriodEnd() bool`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *AccountResponsePlan) GetCancelAtPeriodEndOk() (*bool, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *AccountResponsePlan) SetCancelAtPeriodEnd(v bool)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


