# ChannelTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | **[]string** | Channel slugs. When omitted, API key scope determines recipients. | 

## Methods

### NewChannelTarget

`func NewChannelTarget(channels []string, ) *ChannelTarget`

NewChannelTarget instantiates a new ChannelTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelTargetWithDefaults

`func NewChannelTargetWithDefaults() *ChannelTarget`

NewChannelTargetWithDefaults instantiates a new ChannelTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *ChannelTarget) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ChannelTarget) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ChannelTarget) SetChannels(v []string)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


