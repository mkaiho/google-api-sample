package infrastructure

import (
	"context"
	"errors"

	"github.com/mkaiho/google-api-sample/adapter/gbpapi"
	"google.golang.org/api/mybusinessnotifications/v1"
)

var _ gbpapi.GBPNotification = (*gbpNotification)(nil)

/** My Business Notifications **/
/** See: https://developers.google.com/my-business/reference/notifications/rest **/
type gbpNotification struct {
	client *mybusinessnotifications.Service
}

func NewGBPNotification(ctx context.Context, credentials gbpapi.GBPCredential) (*gbpNotification, error) {
	if ctx == nil {
		return nil, errors.New("ctx is required")
	}
	if credentials == nil {
		return nil, errors.New("credentials is required")
	}
	opts := newGBPOption(ctx, credentials)
	client, err := mybusinessnotifications.NewService(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &gbpNotification{
		client: client,
	}, nil
}
