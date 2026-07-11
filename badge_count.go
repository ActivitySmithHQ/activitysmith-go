package activitysmith

import (
	"fmt"
	"math"

	"github.com/ActivitySmithHQ/activitysmith-go/generated"
)

// BadgeCount sets the count shown on the ActivitySmith app icon. Pass 0 to clear it.
func (c *Client) BadgeCount(value int, channels ...string) (*generated.AppIconBadgeCountUpdateResponse, error) {
	if value < 0 || value > math.MaxInt32 {
		return nil, fmt.Errorf("activitysmith: badge count must be between 0 and %d", math.MaxInt32)
	}

	request := *generated.NewAppIconBadgeCountUpdateRequest(int32(value))
	if len(channels) > 0 {
		request.SetTarget(generated.ChannelTarget{Channels: append([]string{}, channels...)})
	}

	response, _, err := c.apiClient.AppIconBadgesAPI.
		UpdateAppIconBadgeCount(c.ctx).
		AppIconBadgeCountUpdateRequest(request).
		Execute()

	return response, err
}
