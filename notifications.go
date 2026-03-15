package activitysmith

import (
	"fmt"
	"strings"

	"github.com/ActivitySmithHQ/activitysmith-go/generated"
)

type NotificationsService struct {
	client *Client
}

func (s *NotificationsService) Send(input any) (*generated.PushNotificationResponse, error) {
	request, err := normalizePushNotificationRequest(input)
	if err != nil {
		return nil, err
	}

	response, _, err := s.client.apiClient.PushNotificationsAPI.
		SendPushNotification(s.client.ctx).
		PushNotificationRequest(request).
		Execute()

	return response, err
}

// Backward-compatible alias.
func (s *NotificationsService) SendPushNotification(request generated.PushNotificationRequest) (*generated.PushNotificationResponse, error) {
	return s.Send(request)
}

func normalizePushNotificationRequest(input any) (generated.PushNotificationRequest, error) {
	var request generated.PushNotificationRequest

	switch v := input.(type) {
	case PushNotificationInput:
		request = v.toGenerated()
	case *PushNotificationInput:
		if v == nil {
			return generated.PushNotificationRequest{}, fmt.Errorf("activitysmith: input cannot be nil")
		}
		request = v.toGenerated()
	case generated.PushNotificationRequest:
		request = v
	case *generated.PushNotificationRequest:
		if v == nil {
			return generated.PushNotificationRequest{}, fmt.Errorf("activitysmith: input cannot be nil")
		}
		request = *v
	default:
		return generated.PushNotificationRequest{}, fmt.Errorf(
			"activitysmith: unsupported push notification input type %T",
			input,
		)
	}

	if err := validatePushNotificationRequest(request); err != nil {
		return generated.PushNotificationRequest{}, err
	}

	return request, nil
}

func validatePushNotificationRequest(request generated.PushNotificationRequest) error {
	if request.Media == nil {
		return nil
	}

	if strings.TrimSpace(*request.Media) == "" {
		return nil
	}

	if len(request.Actions) > 0 {
		return ErrPushNotificationMediaActionsConflict
	}

	return nil
}
