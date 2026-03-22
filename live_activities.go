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

func (s *LiveActivitiesService) Stream(streamKey string, input any) (*generated.LiveActivityStreamPutResponse, error) {
	request, err := normalizeLiveActivityStreamRequest(input)
	if err != nil {
		return nil, err
	}

	response, _, err := s.client.apiClient.LiveActivitiesAPI.
		ReconcileLiveActivityStream(s.client.ctx, streamKey).
		LiveActivityStreamRequest(request).
		Execute()

	return response, err
}

func (s *LiveActivitiesService) EndStream(streamKey string, input any) (*generated.LiveActivityStreamDeleteResponse, error) {
	request, hasRequest, err := normalizeLiveActivityStreamDeleteRequest(input)
	if err != nil {
		return nil, err
	}

	call := s.client.apiClient.LiveActivitiesAPI.EndLiveActivityStream(s.client.ctx, streamKey)
	if hasRequest {
		call = call.LiveActivityStreamDeleteRequest(request)
	}

	response, _, err := call.Execute()
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

func (s *LiveActivitiesService) ReconcileLiveActivityStream(streamKey string, request generated.LiveActivityStreamRequest) (*generated.LiveActivityStreamPutResponse, error) {
	return s.Stream(streamKey, request)
}

func (s *LiveActivitiesService) EndLiveActivityStream(streamKey string, request generated.LiveActivityStreamDeleteRequest) (*generated.LiveActivityStreamDeleteResponse, error) {
	return s.EndStream(streamKey, request)
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

func normalizeLiveActivityStreamRequest(input any) (generated.LiveActivityStreamRequest, error) {
	switch v := input.(type) {
	case LiveActivityStreamInput:
		return v.toGenerated(), nil
	case *LiveActivityStreamInput:
		if v == nil {
			return generated.LiveActivityStreamRequest{}, fmt.Errorf("activitysmith: input cannot be nil")
		}
		return v.toGenerated(), nil
	case generated.LiveActivityStreamRequest:
		return v, nil
	case *generated.LiveActivityStreamRequest:
		if v == nil {
			return generated.LiveActivityStreamRequest{}, fmt.Errorf("activitysmith: input cannot be nil")
		}
		return *v, nil
	default:
		return generated.LiveActivityStreamRequest{}, fmt.Errorf(
			"activitysmith: unsupported live activity stream input type %T",
			input,
		)
	}
}

func normalizeLiveActivityStreamDeleteRequest(input any) (generated.LiveActivityStreamDeleteRequest, bool, error) {
	switch v := input.(type) {
	case nil:
		return generated.LiveActivityStreamDeleteRequest{}, false, nil
	case LiveActivityStreamEndInput:
		return v.toGenerated(), true, nil
	case *LiveActivityStreamEndInput:
		if v == nil {
			return generated.LiveActivityStreamDeleteRequest{}, false, nil
		}
		return v.toGenerated(), true, nil
	case generated.LiveActivityStreamDeleteRequest:
		return v, true, nil
	case *generated.LiveActivityStreamDeleteRequest:
		if v == nil {
			return generated.LiveActivityStreamDeleteRequest{}, false, nil
		}
		return *v, true, nil
	default:
		return generated.LiveActivityStreamDeleteRequest{}, false, fmt.Errorf(
			"activitysmith: unsupported live activity stream end input type %T",
			input,
		)
	}
}
