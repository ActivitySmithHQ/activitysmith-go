# AccountResponseTrial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**EndsAt** | Pointer to **time.Time** | When a time trial ends. | [optional] 
**DaysLeft** | Pointer to **int32** | Days left in a time trial. 0 after it ends. | [optional] 
**EventsUsed** | Pointer to **int32** | Events used in an event trial. | [optional] 
**EventLimit** | Pointer to **int32** | Events included in an event trial. | [optional] 
**Ended** | **bool** | True when the trial has ended or its events are used up. | 

## Methods

### NewAccountResponseTrial

`func NewAccountResponseTrial(type_ string, ended bool, ) *AccountResponseTrial`

NewAccountResponseTrial instantiates a new AccountResponseTrial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResponseTrialWithDefaults

`func NewAccountResponseTrialWithDefaults() *AccountResponseTrial`

NewAccountResponseTrialWithDefaults instantiates a new AccountResponseTrial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccountResponseTrial) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountResponseTrial) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountResponseTrial) SetType(v string)`

SetType sets Type field to given value.


### GetEndsAt

`func (o *AccountResponseTrial) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *AccountResponseTrial) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *AccountResponseTrial) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *AccountResponseTrial) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetDaysLeft

`func (o *AccountResponseTrial) GetDaysLeft() int32`

GetDaysLeft returns the DaysLeft field if non-nil, zero value otherwise.

### GetDaysLeftOk

`func (o *AccountResponseTrial) GetDaysLeftOk() (*int32, bool)`

GetDaysLeftOk returns a tuple with the DaysLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysLeft

`func (o *AccountResponseTrial) SetDaysLeft(v int32)`

SetDaysLeft sets DaysLeft field to given value.

### HasDaysLeft

`func (o *AccountResponseTrial) HasDaysLeft() bool`

HasDaysLeft returns a boolean if a field has been set.

### GetEventsUsed

`func (o *AccountResponseTrial) GetEventsUsed() int32`

GetEventsUsed returns the EventsUsed field if non-nil, zero value otherwise.

### GetEventsUsedOk

`func (o *AccountResponseTrial) GetEventsUsedOk() (*int32, bool)`

GetEventsUsedOk returns a tuple with the EventsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUsed

`func (o *AccountResponseTrial) SetEventsUsed(v int32)`

SetEventsUsed sets EventsUsed field to given value.

### HasEventsUsed

`func (o *AccountResponseTrial) HasEventsUsed() bool`

HasEventsUsed returns a boolean if a field has been set.

### GetEventLimit

`func (o *AccountResponseTrial) GetEventLimit() int32`

GetEventLimit returns the EventLimit field if non-nil, zero value otherwise.

### GetEventLimitOk

`func (o *AccountResponseTrial) GetEventLimitOk() (*int32, bool)`

GetEventLimitOk returns a tuple with the EventLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLimit

`func (o *AccountResponseTrial) SetEventLimit(v int32)`

SetEventLimit sets EventLimit field to given value.

### HasEventLimit

`func (o *AccountResponseTrial) HasEventLimit() bool`

HasEventLimit returns a boolean if a field has been set.

### GetEnded

`func (o *AccountResponseTrial) GetEnded() bool`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *AccountResponseTrial) GetEndedOk() (*bool, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *AccountResponseTrial) SetEnded(v bool)`

SetEnded sets Ended field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


