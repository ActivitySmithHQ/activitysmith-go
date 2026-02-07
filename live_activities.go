package activitysmith

import "github.com/ActivitySmithHQ/activitysmith-go/generated"

type LiveActivitiesService struct {
	client *Client
}

func (s *LiveActivitiesService) Start(request generated.LiveActivityStartRequest) (*generated.LiveActivityStartResponse, error) {
	response, _, err := s.client.apiClient.LiveActivitiesAPI.
		StartLiveActivity(s.client.ctx).
		LiveActivityStartRequest(request).
		Execute()

	return response, err
}

func (s *LiveActivitiesService) Update(request generated.LiveActivityUpdateRequest) (*generated.LiveActivityUpdateResponse, error) {
	response, _, err := s.client.apiClient.LiveActivitiesAPI.
		UpdateLiveActivity(s.client.ctx).
		LiveActivityUpdateRequest(request).
		Execute()

	return response, err
}

func (s *LiveActivitiesService) End(request generated.LiveActivityEndRequest) (*generated.LiveActivityEndResponse, error) {
	response, _, err := s.client.apiClient.LiveActivitiesAPI.
		EndLiveActivity(s.client.ctx).
		LiveActivityEndRequest(request).
		Execute()

	return response, err
}

// Backward-compatible aliases.
func (s *LiveActivitiesService) StartLiveActivity(request generated.LiveActivityStartRequest) (*generated.LiveActivityStartResponse, error) {
	return s.Start(request)
}

func (s *LiveActivitiesService) UpdateLiveActivity(request generated.LiveActivityUpdateRequest) (*generated.LiveActivityUpdateResponse, error) {
	return s.Update(request)
}

func (s *LiveActivitiesService) EndLiveActivity(request generated.LiveActivityEndRequest) (*generated.LiveActivityEndResponse, error) {
	return s.End(request)
}
