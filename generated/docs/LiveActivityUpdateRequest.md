# LiveActivityUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | **string** |  | 
**ContentState** | [**ContentStateUpdate**](ContentStateUpdate.md) |  | 

## Methods

### NewLiveActivityUpdateRequest

`func NewLiveActivityUpdateRequest(activityId string, contentState ContentStateUpdate, ) *LiveActivityUpdateRequest`

NewLiveActivityUpdateRequest instantiates a new LiveActivityUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityUpdateRequestWithDefaults

`func NewLiveActivityUpdateRequestWithDefaults() *LiveActivityUpdateRequest`

NewLiveActivityUpdateRequestWithDefaults instantiates a new LiveActivityUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityId

`func (o *LiveActivityUpdateRequest) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *LiveActivityUpdateRequest) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *LiveActivityUpdateRequest) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.


### GetContentState

`func (o *LiveActivityUpdateRequest) GetContentState() ContentStateUpdate`

GetContentState returns the ContentState field if non-nil, zero value otherwise.

### GetContentStateOk

`func (o *LiveActivityUpdateRequest) GetContentStateOk() (*ContentStateUpdate, bool)`

GetContentStateOk returns a tuple with the ContentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentState

`func (o *LiveActivityUpdateRequest) SetContentState(v ContentStateUpdate)`

SetContentState sets ContentState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


