package activitysmith

import "github.com/ActivitySmithHQ/activitysmith-go/generated"

type LiveActivitiesService struct {
	client *Client
}

func (s *LiveActivitiesService) StartLiveActivity(request generated.LiveActivityStartRequest) (generated.LiveActivityStartResponse, error) {
	response, _, err := s.client.apiClient.LiveActivitiesAPI.
		StartLiveActivity(s.client.ctx).
		LiveActivityStartRequest(request).
		Execute()

	return response, err
}

func (s *LiveActivitiesService) UpdateLiveActivity(request generated.LiveActivityUpdateRequest) (generated.LiveActivityUpdateResponse, error) {
	response, _, err := s.client.apiClient.LiveActivitiesAPI.
		UpdateLiveActivity(s.client.ctx).
		LiveActivityUpdateRequest(request).
		Execute()

	return response, err
}

func (s *LiveActivitiesService) EndLiveActivity(request generated.LiveActivityEndRequest) (generated.LiveActivityEndResponse, error) {
	response, _, err := s.client.apiClient.LiveActivitiesAPI.
		EndLiveActivity(s.client.ctx).
		LiveActivityEndRequest(request).
		Execute()

	return response, err
}
