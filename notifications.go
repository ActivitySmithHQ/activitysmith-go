package activitysmith

import "github.com/ActivitySmithHQ/activitysmith-go/generated"

type NotificationsService struct {
	client *Client
}

func (s *NotificationsService) Send(request generated.PushNotificationRequest) (*generated.PushNotificationResponse, error) {
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
