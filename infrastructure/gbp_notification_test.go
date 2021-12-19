package infrastructure

import (
	"context"
	"testing"

	"github.com/mkaiho/google-api-sample/adapter/gbpapi"
	"github.com/stretchr/testify/assert"
)

func TestNewGBPNotification(t *testing.T) {
	type args struct {
		ctx    context.Context
		config gbpapi.GBPConfig
	}
	tests := []struct {
		name    string
		args    args
		want    assert.ValueAssertionFunc
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Return GBPNotification",
			args: args{
				ctx:    context.Background(),
				config: &gbpConfig{},
			},
			want: func(tt assert.TestingT, got interface{}, _ ...interface{}) bool {
				return assert.NotNil(tt, got, "NewGBPNotification() got = %v")
			},
			wantErr: assert.NoError,
		},
		{
			name: "Return error when args.ctx is nil",
			args: args{
				ctx:    nil,
				config: &gbpConfig{},
			},
			want: func(tt assert.TestingT, got interface{}, _ ...interface{}) bool {
				return assert.Nil(tt, got, "NewGBPNotification() got = %v, want %v", got, nil)
			},
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "ctx is required"
				return assert.EqualError(tt, e, want, "NewGBPNotification() error = %v, wantErr %v", e, want)
			},
		},
		{
			name: "Return error when args.config is nil",
			args: args{
				ctx:    context.Background(),
				config: nil,
			},
			want: func(tt assert.TestingT, got interface{}, _ ...interface{}) bool {
				return assert.Nil(tt, got, "NewGBPNotification() got = %v, want %v", got, nil)
			},
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "config is required"
				return assert.EqualError(tt, e, want, "NewGBPNotification() error = %v, wantErr %v", e, want)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGBPNotification(tt.args.ctx, tt.args.config)
			if !tt.wantErr(t, err) {
				t.Errorf("NewGBPNotification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want(t, got)
		})
	}
}
