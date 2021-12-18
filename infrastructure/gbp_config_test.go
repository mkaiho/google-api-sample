package infrastructure

import (
	"os"
	"testing"

	"github.com/mkaiho/google-api-sample/adapter/gbpapi"
	"github.com/stretchr/testify/assert"
)

func TestNewGBPCredentialEnv(t *testing.T) {
	type args struct {
		clientID     string
		secret       string
		refreshToken string
	}
	tests := []struct {
		name string
		args args
		want gbpapi.GBPCredential
	}{
		{
			name: "Return GBPCredential",
			args: args{
				clientID:     "dummy_id",
				secret:       "dummy_secret",
				refreshToken: "dummy_refresh_token",
			},
			want: &gbpCredential{
				clientID:     "dummy_id",
				secret:       "dummy_secret",
				refreshToken: "dummy_refresh_token",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGBPCredentialEnv(tt.args.clientID, tt.args.secret, tt.args.refreshToken)
			assert.Equal(t, tt.want, got, "NewGBPCredentialEnv() = %v, want %v", got, tt.want)
		})
	}
}

func TestLoadGBPCredentialEnv(t *testing.T) {
	type fields struct {
		clientID     string
		secret       string
		refreshToken string
	}
	tests := []struct {
		name    string
		fields  fields
		want    gbpapi.GBPCredential
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Load env and return GBPCredential",
			fields: fields{
				clientID:     "dummy_id",
				secret:       "dummy_secret",
				refreshToken: "dummy_refresh_token",
			},
			want: &gbpCredential{
				clientID:     "dummy_id",
				secret:       "dummy_secret",
				refreshToken: "dummy_refresh_token",
			},
			wantErr: assert.NoError,
		},
		{
			name: "Return error when environment valirable GBP_CLIENT_ID is unset",
			fields: fields{
				clientID:     "",
				secret:       "dummy_secret",
				refreshToken: "dummy_refresh_token",
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "failed to GBP credentials: "
				return assert.Contains(tt, e.Error(), want, "LoadGBPCredentialEnv() error = %v, wantErr %v", e, want)
			},
		},
		{
			name: "Return error when environment valirable GBP_CLIENT_SECRET is unset",
			fields: fields{
				clientID:     "dummy_id",
				secret:       "",
				refreshToken: "dummy_refresh_token",
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "failed to GBP credentials: "
				return assert.Contains(tt, e.Error(), want, "LoadGBPCredentialEnv() error = %v, wantErr %v", e, want)
			},
		},
		{
			name: "Return error when environment valirable GBP_REFRESH_TOKEN is unset",
			fields: fields{
				clientID:     "dummy_id",
				secret:       "dummy_secret",
				refreshToken: "",
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "failed to GBP credentials: "
				return assert.Contains(tt, e.Error(), want, "LoadGBPCredentialEnv() error = %v, wantErr %v", e, want)
			},
		},
	}
	for _, tt := range tests {
		if len(tt.fields.clientID) != 0 {
			os.Setenv("GBP_CLIENT_ID", tt.fields.clientID)
		}
		if len(tt.fields.secret) != 0 {
			os.Setenv("GBP_CLIENT_SECRET", tt.fields.secret)
		}
		if len(tt.fields.refreshToken) != 0 {
			os.Setenv("GBP_REFRESH_TOKEN", tt.fields.refreshToken)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadGBPCredentialEnv()
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got, "LoadGBPCredentialEnv() = %v, want %v", got, tt.want)
		})
		os.Clearenv()
	}
}

func Test_gbpCredential_ClientID(t *testing.T) {
	type fields struct {
		clientID     string
		secret       string
		refreshToken string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Return ClientID",
			fields: fields{
				clientID: "dummy_client_id",
			},
			want: "dummy_client_id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &gbpCredential{
				clientID:     tt.fields.clientID,
				secret:       tt.fields.secret,
				refreshToken: tt.fields.refreshToken,
			}
			got := c.ClientID()
			assert.Equal(t, tt.want, got, "gbpCredential.ClientID() = %v, want %v", got, tt.want)
		})
	}
}

func Test_gbpCredential_ClientSecret(t *testing.T) {
	type fields struct {
		clientID     string
		secret       string
		refreshToken string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Return ClientSecret",
			fields: fields{
				secret: "dummy_secret",
			},
			want: "dummy_secret",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &gbpCredential{
				clientID:     tt.fields.clientID,
				secret:       tt.fields.secret,
				refreshToken: tt.fields.refreshToken,
			}
			got := c.ClientSecret()
			assert.Equal(t, tt.want, got, "gbpCredential.ClientSecret() = %v, want %v", got, tt.want)
		})
	}
}

func Test_gbpCredential_RefreshToken(t *testing.T) {
	type fields struct {
		clientID     string
		secret       string
		refreshToken string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Return refresh token",
			fields: fields{
				refreshToken: "dummy_token",
			},
			want: "dummy_token",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &gbpCredential{
				clientID:     tt.fields.clientID,
				secret:       tt.fields.secret,
				refreshToken: tt.fields.refreshToken,
			}
			got := c.RefreshToken()
			assert.Equal(t, tt.want, got, "gbpCredential.RefreshToken() = %v, want %v", got, tt.want)
		})
	}
}
