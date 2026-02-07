package activitysmith

import (
	"fmt"

	"github.com/ActivitySmithHQ/activitysmith-go/generated"
)

type LiveActivitiesService struct {
	client *Client
}

func (s *LiveActivitiesService) Start(input any) (*generated.LiveActivityStartResponse, error) {
	request, err := normalizeLiveActivityStartRequest(input)
	if err != nil {
		return nil, err
	}

	response, _, err := s.client.apiClient.LiveActivitiesAPI.
		StartLiveActivity(s.client.ctx).
		LiveActivityStartRequest(request).
		Execute()

	return response, err
}

func (s *LiveActivitiesService) Update(input any) (*generated.LiveActivityUpdateResponse, error) {
	request, err := normalizeLiveActivityUpdateRequest(input)
	if err != nil {
		return nil, err
	}

	response, _, err := s.client.apiClient.LiveActivitiesAPI.
		UpdateLiveActivity(s.client.ctx).
		LiveActivityUpdateRequest(request).
		Execute()

	return response, err
}

func (s *LiveActivitiesService) End(input any) (*generated.LiveActivityEndResponse, error) {
	request, err := normalizeLiveActivityEndRequest(input)
	if err != nil {
		return nil, err
	}

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

func normalizeLiveActivityStartRequest(input any) (generated.LiveActivityStartRequest, error) {
	switch v := input.(type) {
	case LiveActivityStartInput:
		return v.toGenerated(), nil
	case *LiveActivityStartInput:
		if v == nil {
			return generated.LiveActivityStartRequest{}, fmt.Errorf("activitysmith: input cannot be nil")
		}
		return v.toGenerated(), nil
	case generated.LiveActivityStartRequest:
		return v, nil
	case *generated.LiveActivityStartRequest:
		if v == nil {
			return generated.LiveActivityStartRequest{}, fmt.Errorf("activitysmith: input cannot be nil")
		}
		return *v, nil
	default:
		return generated.LiveActivityStartRequest{}, fmt.Errorf(
			"activitysmith: unsupported live activity start input type %T",
			input,
		)
	}
}

func normalizeLiveActivityUpdateRequest(input any) (generated.LiveActivityUpdateRequest, error) {
	switch v := input.(type) {
	case LiveActivityUpdateInput:
		return v.toGenerated(), nil
	case *LiveActivityUpdateInput:
		if v == nil {
			return generated.LiveActivityUpdateRequest{}, fmt.Errorf("activitysmith: input cannot be nil")
		}
		return v.toGenerated(), nil
	case generated.LiveActivityUpdateRequest:
		return v, nil
	case *generated.LiveActivityUpdateRequest:
		if v == nil {
			return generated.LiveActivityUpdateRequest{}, fmt.Errorf("activitysmith: input cannot be nil")
		}
		return *v, nil
	default:
		return generated.LiveActivityUpdateRequest{}, fmt.Errorf(
			"activitysmith: unsupported live activity update input type %T",
			input,
		)
	}
}

func normalizeLiveActivityEndRequest(input any) (generated.LiveActivityEndRequest, error) {
	switch v := input.(type) {
	case LiveActivityEndInput:
		return v.toGenerated(), nil
	case *LiveActivityEndInput:
		if v == nil {
			return generated.LiveActivityEndRequest{}, fmt.Errorf("activitysmith: input cannot be nil")
		}
		return v.toGenerated(), nil
	case generated.LiveActivityEndRequest:
		return v, nil
	case *generated.LiveActivityEndRequest:
		if v == nil {
			return generated.LiveActivityEndRequest{}, fmt.Errorf("activitysmith: input cannot be nil")
		}
		return *v, nil
	default:
		return generated.LiveActivityEndRequest{}, fmt.Errorf(
			"activitysmith: unsupported live activity end input type %T",
			input,
		)
	}
}
