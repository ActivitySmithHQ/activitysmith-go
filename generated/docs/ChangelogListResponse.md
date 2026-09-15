# ChangelogListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changelogs** | [**[]ChangelogEntry**](ChangelogEntry.md) |  | 

## Methods

### NewChangelogListResponse

`func NewChangelogListResponse(changelogs []ChangelogEntry, ) *ChangelogListResponse`

NewChangelogListResponse instantiates a new ChangelogListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangelogListResponseWithDefaults

`func NewChangelogListResponseWithDefaults() *ChangelogListResponse`

NewChangelogListResponseWithDefaults instantiates a new ChangelogListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangelogs

`func (o *ChangelogListResponse) GetChangelogs() []ChangelogEntry`

GetChangelogs returns the Changelogs field if non-nil, zero value otherwise.

### GetChangelogsOk

`func (o *ChangelogListResponse) GetChangelogsOk() (*[]ChangelogEntry, bool)`

GetChangelogsOk returns a tuple with the Changelogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogs

`func (o *ChangelogListResponse) SetChangelogs(v []ChangelogEntry)`

SetChangelogs sets Changelogs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


