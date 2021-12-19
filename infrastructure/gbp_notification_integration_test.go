//go:build gbpapi
// +build gbpapi

package infrastructure

import (
	"context"
	"reflect"
	"testing"

	"github.com/mkaiho/google-api-sample/adapter/gbpapi"
	"google.golang.org/api/mybusinessnotifications/v1"
)

func Test_gbpNotification_GetNotificationSetting(t *testing.T) {
	type fields struct {
		client *mybusinessnotifications.Service
	}
	type args struct {
		ctx       context.Context
		accountID gbpapi.AccountID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *gbpapi.NotificationSetting
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &gbpNotification{
				client: tt.fields.client,
			}
			got, err := n.GetNotificationSetting(tt.args.ctx, tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("gbpNotification.GetNotificationSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gbpNotification.GetNotificationSetting() = %v, want %v", got, tt.want)
			}
		})
	}
}
