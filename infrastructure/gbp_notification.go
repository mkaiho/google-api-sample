package infrastructure

import (
	"context"
	"errors"
	"fmt"

	"github.com/mkaiho/google-api-sample/adapter/gbpapi"
	"google.golang.org/api/mybusinessnotifications/v1"
)

var _ gbpapi.GBPNotification = (*gbpNotification)(nil)

/** My Business Notifications **/
/** See: https://developers.google.com/my-business/reference/notifications/rest **/
type gbpNotification struct {
	client *mybusinessnotifications.Service
}

func NewGBPNotification(ctx context.Context, config gbpapi.GBPConfig) (*gbpNotification, error) {
	if ctx == nil {
		return nil, errors.New("ctx is required")
	}
	if config == nil {
		return nil, errors.New("config is required")
	}
	opts := newGBPOption(ctx, config)
	client, err := mybusinessnotifications.NewService(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &gbpNotification{
		client: client,
	}, nil
}

/** Get notification settings **/
/** See: https://developers.google.com/my-business/reference/notifications/rest/v1/accounts/getNotificationSetting **/
func (n *gbpNotification) GetNotificationSetting(ctx context.Context, accountID gbpapi.AccountID) (*gbpapi.NotificationSetting, error) {
	name := gbpapi.NewNotificationSettingName(accountID)
	res, err := n.client.Accounts.GetNotificationSetting(name.String()).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to Notifications.accounts.getNotificationSetting request: %w", err)
	}
	notificationTypes, err := gbpapi.ParseNotificationTypes(res.NotificationTypes)
	if err != nil {
		return nil, fmt.Errorf("failed to convert Notifications.accounts.getNotificationSetting response: %w", err)
	}

	return &gbpapi.NotificationSetting{
		AccountID:         accountID,
		TopicName:         res.PubsubTopic,
		NotificationTypes: notificationTypes,
	}, nil
}
