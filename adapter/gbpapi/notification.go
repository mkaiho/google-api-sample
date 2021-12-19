package gbpapi

import (
	"context"
	"fmt"
)

/** NotificationSetting.name: accounts/{account_id}/notificationSetting **/
func NewNotificationSettingName(accountID AccountID) NotificationSettingName {
	return notificationSettingName{
		accountID: accountID,
	}
}

type NotificationSettingName interface {
	AccountID() AccountID
	String() string
}

type notificationSettingName struct {
	accountID AccountID
}

func (n notificationSettingName) AccountID() AccountID {
	return n.accountID
}

func (n notificationSettingName) String() string {
	return fmt.Sprintf("accounts/%d/notificationSetting", n.accountID)
}

/** NotificationType **/
/** See: https://developers.google.com/my-business/reference/notifications/rest/v1/NotificationSetting#NotificationType **/
const (
	NotificationTypeUnspecified           NotificationType = "NOTIFICATION_TYPE_UNSPECIFIED"
	NotificationTypeGoogleUpdated         NotificationType = "GOOGLE_UPDATE"
	NotificationTypeNewReview             NotificationType = "NEW_REVIEW"
	NotificationTypeUpdatedReview         NotificationType = "UPDATED_REVIEW"
	NotificationTypeNewCustomerMedia      NotificationType = "NEW_CUSTOMER_MEDIA"
	NotificationTypeNewQuestion           NotificationType = "NEW_QUESTION"
	NotificationTypeUpdatedQuestion       NotificationType = "UPDATED_QUESTION"
	NotificationTypeNewAnswer             NotificationType = "NEW_ANSWER"
	NotificationTypeUpdatedAnswer         NotificationType = "UPDATED_ANSWER"
	NotificationTypeDuplicateLocation     NotificationType = "DUPLICATE_LOCATION"
	NotificationTypeLossOfVoiceOfMarchant NotificationType = "LOSS_OF_VOICE_OF_MERCHANT"
)

var allNotificationTypes = NotificationTypes{
	NotificationTypeUnspecified,
	NotificationTypeGoogleUpdated,
	NotificationTypeNewReview,
	NotificationTypeUpdatedReview,
	NotificationTypeNewCustomerMedia,
	NotificationTypeNewQuestion,
	NotificationTypeUpdatedQuestion,
	NotificationTypeNewAnswer,
	NotificationTypeUpdatedAnswer,
	NotificationTypeDuplicateLocation,
	NotificationTypeLossOfVoiceOfMarchant,
}

func ParseNotificationType(value string) (*NotificationType, error) {
	notificationType := NotificationType(value)
	if !allNotificationTypes.Includes(notificationType) {
		return nil, fmt.Errorf("invalid notification type: %s", value)
	}
	return &notificationType, nil
}

func ParseNotificationTypes(values []string) (NotificationTypes, error) {
	notificationTypes := make([]NotificationType, len(values))
	for i, v := range values {
		notificationType, err := ParseNotificationType(v)
		if err != nil {
			return nil, fmt.Errorf("values include invalid notification type: %w", err)
		}
		notificationTypes[i] = *notificationType
	}
	return notificationTypes, nil
}

func NewNotificationTypes(values ...NotificationType) NotificationTypes {
	return (NotificationTypes)(values)
}

type NotificationType string

func (t NotificationType) String() string {
	return string(t)
}

type NotificationTypes []NotificationType

func (nts NotificationTypes) Includes(notificationType NotificationType) bool {
	for _, v := range nts {
		if notificationType == v {
			return true
		}
	}
	return false
}

func (nts NotificationTypes) Array() []NotificationType {
	return ([]NotificationType)(nts)
}

func (nt NotificationTypes) StringsArray() []string {
	values := make([]string, len(nt))
	for i, v := range nt {
		values[i] = v.String()
	}
	return values
}

/** NotificationSetting **/
type NotificationSetting struct {
	AccountID         AccountID
	TopicName         string
	NotificationTypes NotificationTypes
}

/** GBP Notifications APIs **/
type GBPNotification interface {
	GetNotificationSetting(ctx context.Context, accountID AccountID) (*NotificationSetting, error)
}
