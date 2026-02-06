package generated

import "context"

type contextKey string

const ContextAccessToken contextKey = "accessToken"

type ServerConfiguration struct {
	URL string
}

type ServerConfigurations []ServerConfiguration

type Configuration struct {
	Servers ServerConfigurations
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

type APIClient struct {
	PushNotificationsAPI *PushNotificationsAPIService
	LiveActivitiesAPI    *LiveActivitiesAPIService
}

func NewAPIClient(_ *Configuration) *APIClient {
	return &APIClient{
		PushNotificationsAPI: &PushNotificationsAPIService{},
		LiveActivitiesAPI:    &LiveActivitiesAPIService{},
	}
}

type PushNotificationRequest struct{}
type PushNotificationResponse struct{}
type LiveActivityStartRequest struct{}
type LiveActivityStartResponse struct{}
type LiveActivityUpdateRequest struct{}
type LiveActivityUpdateResponse struct{}
type LiveActivityEndRequest struct{}
type LiveActivityEndResponse struct{}

type PushNotificationsAPIService struct{}
type LiveActivitiesAPIService struct{}

type ApiSendPushNotificationRequest struct{}
type ApiStartLiveActivityRequest struct{}
type ApiUpdateLiveActivityRequest struct{}
type ApiEndLiveActivityRequest struct{}

func (s *PushNotificationsAPIService) SendPushNotification(_ context.Context) ApiSendPushNotificationRequest {
	return ApiSendPushNotificationRequest{}
}

func (r ApiSendPushNotificationRequest) PushNotificationRequest(_ PushNotificationRequest) ApiSendPushNotificationRequest {
	return r
}

func (r ApiSendPushNotificationRequest) Execute() (PushNotificationResponse, any, error) {
	return PushNotificationResponse{}, nil, nil
}

func (s *LiveActivitiesAPIService) StartLiveActivity(_ context.Context) ApiStartLiveActivityRequest {
	return ApiStartLiveActivityRequest{}
}

func (r ApiStartLiveActivityRequest) LiveActivityStartRequest(_ LiveActivityStartRequest) ApiStartLiveActivityRequest {
	return r
}

func (r ApiStartLiveActivityRequest) Execute() (LiveActivityStartResponse, any, error) {
	return LiveActivityStartResponse{}, nil, nil
}

func (s *LiveActivitiesAPIService) UpdateLiveActivity(_ context.Context) ApiUpdateLiveActivityRequest {
	return ApiUpdateLiveActivityRequest{}
}

func (r ApiUpdateLiveActivityRequest) LiveActivityUpdateRequest(_ LiveActivityUpdateRequest) ApiUpdateLiveActivityRequest {
	return r
}

func (r ApiUpdateLiveActivityRequest) Execute() (LiveActivityUpdateResponse, any, error) {
	return LiveActivityUpdateResponse{}, nil, nil
}

func (s *LiveActivitiesAPIService) EndLiveActivity(_ context.Context) ApiEndLiveActivityRequest {
	return ApiEndLiveActivityRequest{}
}

func (r ApiEndLiveActivityRequest) LiveActivityEndRequest(_ LiveActivityEndRequest) ApiEndLiveActivityRequest {
	return r
}

func (r ApiEndLiveActivityRequest) Execute() (LiveActivityEndResponse, any, error) {
	return LiveActivityEndResponse{}, nil, nil
}
