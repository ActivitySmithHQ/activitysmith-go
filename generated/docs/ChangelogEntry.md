# ChangelogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Platform** | **string** |  | 
**Version** | **string** |  | 
**Title** | **string** |  | 
**Subtitle** | **NullableString** |  | 
**HeroImageUrl** | **NullableString** |  | 
**CtaTitle** | **string** |  | 
**PublishedAt** | **NullableTime** |  | 
**Items** | [**[]ChangelogItem**](ChangelogItem.md) |  | 

## Methods

### NewChangelogEntry

`func NewChangelogEntry(id string, platform string, version string, title string, subtitle NullableString, heroImageUrl NullableString, ctaTitle string, publishedAt NullableTime, items []ChangelogItem, ) *ChangelogEntry`

NewChangelogEntry instantiates a new ChangelogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangelogEntryWithDefaults

`func NewChangelogEntryWithDefaults() *ChangelogEntry`

NewChangelogEntryWithDefaults instantiates a new ChangelogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChangelogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangelogEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangelogEntry) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *ChangelogEntry) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ChangelogEntry) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ChangelogEntry) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetVersion

`func (o *ChangelogEntry) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ChangelogEntry) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ChangelogEntry) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetTitle

`func (o *ChangelogEntry) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChangelogEntry) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChangelogEntry) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSubtitle

`func (o *ChangelogEntry) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *ChangelogEntry) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *ChangelogEntry) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.


### SetSubtitleNil

`func (o *ChangelogEntry) SetSubtitleNil(b bool)`

 SetSubtitleNil sets the value for Subtitle to be an explicit nil

### UnsetSubtitle
`func (o *ChangelogEntry) UnsetSubtitle()`

UnsetSubtitle ensures that no value is present for Subtitle, not even an explicit nil
### GetHeroImageUrl

`func (o *ChangelogEntry) GetHeroImageUrl() string`

GetHeroImageUrl returns the HeroImageUrl field if non-nil, zero value otherwise.

### GetHeroImageUrlOk

`func (o *ChangelogEntry) GetHeroImageUrlOk() (*string, bool)`

GetHeroImageUrlOk returns a tuple with the HeroImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroImageUrl

`func (o *ChangelogEntry) SetHeroImageUrl(v string)`

SetHeroImageUrl sets HeroImageUrl field to given value.


### SetHeroImageUrlNil

`func (o *ChangelogEntry) SetHeroImageUrlNil(b bool)`

 SetHeroImageUrlNil sets the value for HeroImageUrl to be an explicit nil

### UnsetHeroImageUrl
`func (o *ChangelogEntry) UnsetHeroImageUrl()`

UnsetHeroImageUrl ensures that no value is present for HeroImageUrl, not even an explicit nil
### GetCtaTitle

`func (o *ChangelogEntry) GetCtaTitle() string`

GetCtaTitle returns the CtaTitle field if non-nil, zero value otherwise.

### GetCtaTitleOk

`func (o *ChangelogEntry) GetCtaTitleOk() (*string, bool)`

GetCtaTitleOk returns a tuple with the CtaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtaTitle

`func (o *ChangelogEntry) SetCtaTitle(v string)`

SetCtaTitle sets CtaTitle field to given value.


### GetPublishedAt

`func (o *ChangelogEntry) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *ChangelogEntry) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *ChangelogEntry) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.


### SetPublishedAtNil

`func (o *ChangelogEntry) SetPublishedAtNil(b bool)`

 SetPublishedAtNil sets the value for PublishedAt to be an explicit nil

### UnsetPublishedAt
`func (o *ChangelogEntry) UnsetPublishedAt()`

UnsetPublishedAt ensures that no value is present for PublishedAt, not even an explicit nil
### GetItems

`func (o *ChangelogEntry) GetItems() []ChangelogItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ChangelogEntry) GetItemsOk() (*[]ChangelogItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ChangelogEntry) SetItems(v []ChangelogItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


