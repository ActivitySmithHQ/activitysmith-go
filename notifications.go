package activitysmith

import (
	"fmt"

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
	switch v := input.(type) {
	case PushNotificationInput:
		return v.toGenerated(), nil
	case *PushNotificationInput:
		if v == nil {
			return generated.PushNotificationRequest{}, fmt.Errorf("activitysmith: input cannot be nil")
		}
		return v.toGenerated(), nil
	case generated.PushNotificationRequest:
		return v, nil
	case *generated.PushNotificationRequest:
		if v == nil {
			return generated.PushNotificationRequest{}, fmt.Errorf("activitysmith: input cannot be nil")
		}
		return *v, nil
	default:
		return generated.PushNotificationRequest{}, fmt.Errorf(
			"activitysmith: unsupported push notification input type %T",
			input,
		)
	}
}
