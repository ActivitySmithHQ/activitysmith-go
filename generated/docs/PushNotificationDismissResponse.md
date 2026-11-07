# PushNotificationDismissResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | True when the dismissal is recorded. | 
**NotificationId** | **string** | ID of the dismissed Push Notification. | 
**DismissedAt** | **time.Time** | When the notification was first dismissed. Repeat calls return the same time. | 

## Methods

### NewPushNotificationDismissResponse

`func NewPushNotificationDismissResponse(success bool, notificationId string, dismissedAt time.Time, ) *PushNotificationDismissResponse`

NewPushNotificationDismissResponse instantiates a new PushNotificationDismissResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushNotificationDismissResponseWithDefaults

`func NewPushNotificationDismissResponseWithDefaults() *PushNotificationDismissResponse`

NewPushNotificationDismissResponseWithDefaults instantiates a new PushNotificationDismissResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *PushNotificationDismissResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PushNotificationDismissResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PushNotificationDismissResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetNotificationId

`func (o *PushNotificationDismissResponse) GetNotificationId() string`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *PushNotificationDismissResponse) GetNotificationIdOk() (*string, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *PushNotificationDismissResponse) SetNotificationId(v string)`

SetNotificationId sets NotificationId field to given value.


### GetDismissedAt

`func (o *PushNotificationDismissResponse) GetDismissedAt() time.Time`

GetDismissedAt returns the DismissedAt field if non-nil, zero value otherwise.

### GetDismissedAtOk

`func (o *PushNotificationDismissResponse) GetDismissedAtOk() (*time.Time, bool)`

GetDismissedAtOk returns a tuple with the DismissedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedAt

`func (o *PushNotificationDismissResponse) SetDismissedAt(v time.Time)`

SetDismissedAt sets DismissedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


