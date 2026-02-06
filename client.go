package activitysmith

import (
	"context"
	"errors"
	"strings"

	"github.com/ActivitySmithHQ/activitysmith-go/generated"
)

var ErrAPIKeyRequired = errors.New("activitysmith: apiKey is required")

type Options struct {
	BaseURL string
	Context context.Context
}

type Client struct {
	apiClient      *generated.APIClient
	ctx            context.Context
	Notifications  *NotificationsService
	LiveActivities *LiveActivitiesService
}

func New(apiKey string, opts *Options) (*Client, error) {
	if strings.TrimSpace(apiKey) == "" {
		return nil, ErrAPIKeyRequired
	}

	cfg := generated.NewConfiguration()
	ctx := context.Background()

	if opts != nil {
		if opts.BaseURL != "" {
			cfg.Servers = generated.ServerConfigurations{{URL: strings.TrimRight(opts.BaseURL, "/")}}
		}
		if opts.Context != nil {
			ctx = opts.Context
		}
	}

	ctx = context.WithValue(ctx, generated.ContextAccessToken, apiKey)

	apiClient := generated.NewAPIClient(cfg)
	client := &Client{apiClient: apiClient, ctx: ctx}
	client.Notifications = &NotificationsService{client: client}
	client.LiveActivities = &LiveActivitiesService{client: client}

	return client, nil
}

func (c *Client) APIClient() *generated.APIClient {
	return c.apiClient
}

func (c *Client) Context() context.Context {
	return c.ctx
}
