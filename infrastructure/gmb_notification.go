package infrastructure

import (
	"context"

	"google.golang.org/api/mybusinessnotifications/v1"
	"google.golang.org/api/option"
)

/** Credentials for My Business Notifications **/
func NewGBPNotificationOptions() *GBPNotificationOptions {
	return &GBPNotificationOptions{}
}

type GBPNotificationOptions struct {
	APIKey string
}

/** My Business Notifications **/
/** See: https://developers.google.com/my-business/reference/notifications/rest **/
type GBPNotification struct {
	client *mybusinessnotifications.Service
}

func NewGBPNotification(ctx context.Context, opts GBPNotificationOptions) (*GBPNotification, error) {
	client, err := mybusinessnotifications.NewService(ctx, option.WithAPIKey(opts.APIKey))
	if err != nil {
		return nil, err
	}

	return &GBPNotification{
		client: client,
	}, nil
}
