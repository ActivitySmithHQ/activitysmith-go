# LiveActivityEndRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | **string** |  | 
**ContentState** | [**ContentStateEnd**](ContentStateEnd.md) |  | 

## Methods

### NewLiveActivityEndRequest

`func NewLiveActivityEndRequest(activityId string, contentState ContentStateEnd, ) *LiveActivityEndRequest`

NewLiveActivityEndRequest instantiates a new LiveActivityEndRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveActivityEndRequestWithDefaults

`func NewLiveActivityEndRequestWithDefaults() *LiveActivityEndRequest`

NewLiveActivityEndRequestWithDefaults instantiates a new LiveActivityEndRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityId

`func (o *LiveActivityEndRequest) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *LiveActivityEndRequest) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *LiveActivityEndRequest) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.


### GetContentState

`func (o *LiveActivityEndRequest) GetContentState() ContentStateEnd`

GetContentState returns the ContentState field if non-nil, zero value otherwise.

### GetContentStateOk

`func (o *LiveActivityEndRequest) GetContentStateOk() (*ContentStateEnd, bool)`

GetContentStateOk returns a tuple with the ContentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentState

`func (o *LiveActivityEndRequest) SetContentState(v ContentStateEnd)`

SetContentState sets ContentState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


