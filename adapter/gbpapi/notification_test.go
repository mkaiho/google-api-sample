package gbpapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNotificationSettingName(t *testing.T) {
	type args struct {
		accountID AccountID
	}
	tests := []struct {
		name string
		args args
		want NotificationSettingName
	}{
		{
			name: "Return notification setting name",
			args: args{
				accountID: AccountID(1),
			},
			want: notificationSettingName{
				accountID: AccountID(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNotificationSettingName(tt.args.accountID)
			assert.Equal(t, tt.want, got, "NewNotificationSettingName() = %v, want %v", got, tt.want)
		})
	}
}

func Test_notificationSettingName_AccountID(t *testing.T) {
	type fields struct {
		accountID AccountID
	}
	tests := []struct {
		name   string
		fields fields
		want   AccountID
	}{
		{
			name: "Return account ID",
			fields: fields{
				accountID: AccountID(1),
			},
			want: AccountID(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := notificationSettingName{
				accountID: tt.fields.accountID,
			}
			got := n.AccountID()
			assert.Equal(t, tt.want, got, "notificationSettingName.AccountID() = %v, want %v", got, tt.want)
		})
	}
}

func Test_notificationSettingName_String(t *testing.T) {
	type fields struct {
		accountID AccountID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Return account ID",
			fields: fields{
				accountID: AccountID(1),
			},
			want: "accounts/1/notificationSetting",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := notificationSettingName{
				accountID: tt.fields.accountID,
			}
			got := n.String()
			assert.Equal(t, tt.want, got, "notificationSettingName.String() = %v, want %v", got, tt.want)
		})
	}
}

func TestParseNotificationType(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *NotificationType
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Return \"NOTIFICATION_TYPE_UNSPECIFIED\"",
			args: args{
				value: "NOTIFICATION_TYPE_UNSPECIFIED",
			},
			want:    func() *NotificationType { value := NotificationTypeUnspecified; return &value }(),
			wantErr: assert.NoError,
		},
		{
			name: "Return \"GOOGLE_UPDATE\"",
			args: args{
				value: "GOOGLE_UPDATE",
			},
			want:    func() *NotificationType { value := NotificationTypeGoogleUpdated; return &value }(),
			wantErr: assert.NoError,
		},
		{
			name: "Return \"NEW_REVIEW\"",
			args: args{
				value: "NEW_REVIEW",
			},
			want:    func() *NotificationType { value := NotificationTypeNewReview; return &value }(),
			wantErr: assert.NoError,
		},
		{
			name: "Return \"UPDATED_REVIEW\"",
			args: args{
				value: "UPDATED_REVIEW",
			},
			want:    func() *NotificationType { value := NotificationTypeUpdatedReview; return &value }(),
			wantErr: assert.NoError,
		},
		{
			name: "Return \"NEW_CUSTOMER_MEDIA\"",
			args: args{
				value: "NEW_CUSTOMER_MEDIA",
			},
			want:    func() *NotificationType { value := NotificationTypeNewCustomerMedia; return &value }(),
			wantErr: assert.NoError,
		},
		{
			name: "Return \"NEW_QUESTION\"",
			args: args{
				value: "NEW_QUESTION",
			},
			want:    func() *NotificationType { value := NotificationTypeNewQuestion; return &value }(),
			wantErr: assert.NoError,
		},
		{
			name: "Return \"UPDATED_QUESTION\"",
			args: args{
				value: "UPDATED_QUESTION",
			},
			want:    func() *NotificationType { value := NotificationTypeUpdatedQuestion; return &value }(),
			wantErr: assert.NoError,
		},
		{
			name: "Return \"NEW_ANSWER\"",
			args: args{
				value: "NEW_ANSWER",
			},
			want:    func() *NotificationType { value := NotificationTypeNewAnswer; return &value }(),
			wantErr: assert.NoError,
		},
		{
			name: "Return \"UPDATED_ANSWER\"",
			args: args{
				value: "UPDATED_ANSWER",
			},
			want:    func() *NotificationType { value := NotificationTypeUpdatedAnswer; return &value }(),
			wantErr: assert.NoError,
		},
		{
			name: "Return \"DUPLICATE_LOCATION\"",
			args: args{
				value: "DUPLICATE_LOCATION",
			},
			want:    func() *NotificationType { value := NotificationTypeDuplicateLocation; return &value }(),
			wantErr: assert.NoError,
		},
		{
			name: "Return \"LOSS_OF_VOICE_OF_MERCHANT\"",
			args: args{
				value: "LOSS_OF_VOICE_OF_MERCHANT",
			},
			want:    func() *NotificationType { value := NotificationTypeLossOfVoiceOfMarchant; return &value }(),
			wantErr: assert.NoError,
		},
		{
			name: "Return error when value is invalid",
			args: args{
				value: "XXXXX",
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "invalid notification type: XXXXX"
				return assert.EqualError(tt, e, want, "ParseNotificationType() error = %v, wantErr %v", e, want)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNotificationType(tt.args.value)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got, "ParseNotificationType() = %v, want %v", got, tt.want)
		})
	}
}

func TestParseNotificationTypes(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name    string
		args    args
		want    NotificationTypes
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Return notification types",
			args: args{
				values: []string{
					"NOTIFICATION_TYPE_UNSPECIFIED",
					"GOOGLE_UPDATE",
					"NEW_REVIEW",
					"UPDATED_REVIEW",
					"NEW_CUSTOMER_MEDIA",
					"NEW_QUESTION",
					"UPDATED_QUESTION",
					"NEW_ANSWER",
					"UPDATED_ANSWER",
					"DUPLICATE_LOCATION",
					"LOSS_OF_VOICE_OF_MERCHANT",
				},
			},
			want: NotificationTypes{
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
			},
			wantErr: assert.NoError,
		},
		{
			name: "Return empty notification types when values are empty",
			args: args{
				values: []string{},
			},
			want:    NotificationTypes{},
			wantErr: assert.NoError,
		},
		{
			name: "Return error when values include invalid value",
			args: args{
				values: []string{
					"NOTIFICATION_TYPE_UNSPECIFIED",
					"GOOGLE_UPDATE",
					"NEW_REVIEW",
					"UPDATED_REVIEW",
					"NEW_CUSTOMER_MEDIA",
					"NEW_QUESTION",
					"UPDATED_QUESTION",
					"NEW_ANSWER",
					"UPDATED_ANSWER",
					"DUPLICATE_LOCATION",
					"LOSS_OF_VOICE_OF_MERCHANT",
					"XXXXX",
				},
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "values include invalid notification type: invalid notification type: XXXXX"
				return assert.EqualError(tt, e, want, "ParseNotificationTypes() error = %v, wantErr %v", e, want)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNotificationTypes(tt.args.values)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got, "ParseNotificationTypes() = %v, want %v", got, tt.want)
		})
	}
}

func TestNewNotificationTypes(t *testing.T) {
	type args struct {
		values []NotificationType
	}
	tests := []struct {
		name string
		args args
		want NotificationTypes
	}{
		{
			name: "Return notification types",
			args: args{
				values: []NotificationType{
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
				},
			},
			want: NotificationTypes{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNotificationTypes(tt.args.values...)
			assert.Equal(t, tt.want, got, "NewNotificationTypes() = %v, want %v", got, tt.want)
		})
	}
}

func TestNotificationType_String(t *testing.T) {
	tests := []struct {
		name string
		tr   NotificationType
		want string
	}{
		{
			name: "Return \"NOTIFICATION_TYPE_UNSPECIFIED\"",
			tr:   NotificationTypeUnspecified,
			want: "NOTIFICATION_TYPE_UNSPECIFIED",
		},
		{
			name: "Return \"GOOGLE_UPDATE\"",
			tr:   NotificationTypeGoogleUpdated,
			want: "GOOGLE_UPDATE",
		},
		{
			name: "Return \"NEW_REVIEW\"",
			tr:   NotificationTypeNewReview,
			want: "NEW_REVIEW",
		},
		{
			name: "Return \"UPDATED_REVIEW\"",
			tr:   NotificationTypeUpdatedReview,
			want: "UPDATED_REVIEW",
		},
		{
			name: "Return \"NEW_CUSTOMER_MEDIA\"",
			tr:   NotificationTypeNewCustomerMedia,
			want: "NEW_CUSTOMER_MEDIA",
		},
		{
			name: "Return \"NEW_QUESTION\"",
			tr:   NotificationTypeNewQuestion,
			want: "NEW_QUESTION",
		},
		{
			name: "Return \"UPDATED_QUESTION\"",
			tr:   NotificationTypeUpdatedQuestion,
			want: "UPDATED_QUESTION",
		},
		{
			name: "Return \"NEW_ANSWER\"",
			tr:   NotificationTypeNewAnswer,
			want: "NEW_ANSWER",
		},
		{
			name: "Return \"UPDATED_ANSWER\"",
			tr:   NotificationTypeUpdatedAnswer,
			want: "UPDATED_ANSWER",
		},
		{
			name: "Return \"DUPLICATE_LOCATION\"",
			tr:   NotificationTypeDuplicateLocation,
			want: "DUPLICATE_LOCATION",
		},
		{
			name: "Return \"LOSS_OF_VOICE_OF_MERCHANT\"",
			tr:   NotificationTypeLossOfVoiceOfMarchant,
			want: "LOSS_OF_VOICE_OF_MERCHANT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.String()
			assert.Equal(t, tt.want, got, "NotificationType.String() = %v, want %v", got, tt.want)
		})
	}
}

func TestNotificationTypes_Includes(t *testing.T) {
	type args struct {
		notificationType NotificationType
	}
	tests := []struct {
		name string
		nts  NotificationTypes
		args args
		want bool
	}{
		{
			name: "Return true when notification types include value",
			nts: NotificationTypes{
				NotificationTypeGoogleUpdated,
				NotificationTypeNewReview,
				NotificationTypeUpdatedReview,
			},
			args: args{
				NotificationTypeGoogleUpdated,
			},
			want: true,
		},
		{
			name: "Return false when notification types do not include value",
			nts: NotificationTypes{
				NotificationTypeGoogleUpdated,
				NotificationTypeNewReview,
				NotificationTypeUpdatedReview,
			},
			args: args{
				NotificationTypeNewCustomerMedia,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.nts.Includes(tt.args.notificationType)
			assert.Equal(t, tt.want, got, "NotificationTypes.Includes() = %v, want %v", got, tt.want)
		})
	}
}

func TestNotificationTypes_Array(t *testing.T) {
	tests := []struct {
		name string
		nts  NotificationTypes
		want []NotificationType
	}{
		{
			name: "Return notification types array",
			nts: NotificationTypes{
				NotificationTypeGoogleUpdated,
				NotificationTypeNewReview,
				NotificationTypeUpdatedReview,
			},
			want: []NotificationType{
				NotificationTypeGoogleUpdated,
				NotificationTypeNewReview,
				NotificationTypeUpdatedReview,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.nts.Array()
			assert.Equal(t, tt.want, got, "NotificationTypes.Array() = %v, want %v", got, tt.want)
		})
	}
}

func TestNotificationTypes_StringsArray(t *testing.T) {
	tests := []struct {
		name string
		nt   NotificationTypes
		want []string
	}{
		{
			name: "Return strings array",
			nt: NotificationTypes{
				NotificationTypeGoogleUpdated,
				NotificationTypeNewReview,
				NotificationTypeUpdatedReview,
			},
			want: []string{
				"GOOGLE_UPDATE",
				"NEW_REVIEW",
				"UPDATED_REVIEW",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.nt.StringsArray()
			assert.Equal(t, tt.want, got, "NotificationTypes.StringsArray() = %v, want %v", got, tt.want)
		})
	}
}
