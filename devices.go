package activitysmith

import "github.com/ActivitySmithHQ/activitysmith-go/generated"

type DevicesService struct {
	client *Client
}

// DeviceListInput filters the paired device list. Zero values are not sent,
// so the API defaults apply: up to 100 active devices from the first page.
type DeviceListInput struct {
	Limit           int
	Offset          int
	IncludeInactive bool
}

// List lists the iPhones and iPads paired to the account, active devices first.
func (s *DevicesService) List(input DeviceListInput) (*generated.AccountDeviceListResponse, error) {
	request := s.client.apiClient.DevicesAPI.ListAccountDevices(s.client.ctx)
	if input.Limit != 0 {
		request = request.Limit(int32(input.Limit))
	}
	if input.Offset != 0 {
		request = request.Offset(int32(input.Offset))
	}
	if input.IncludeInactive {
		request = request.IncludeInactive(true)
	}

	response, _, err := request.Execute()
	return response, err
}

// Generated-style alias.
func (s *DevicesService) ListAccountDevices(input DeviceListInput) (*generated.AccountDeviceListResponse, error) {
	return s.List(input)
}
