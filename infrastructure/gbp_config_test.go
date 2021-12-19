package infrastructure

import (
	"os"
	"testing"

	"github.com/mkaiho/google-api-sample/adapter/gbpapi"
	"github.com/stretchr/testify/assert"
)

func TestNewGBPConfigEnv(t *testing.T) {
	type args struct {
		clientID     string
		secret       string
		refreshToken string
		redirectURL  string
	}
	tests := []struct {
		name string
		args args
		want gbpapi.GBPConfig
	}{
		{
			name: "Return GBPConfig",
			args: args{
				clientID:     "dummy_id",
				secret:       "dummy_secret",
				refreshToken: "dummy_refresh_token",
				redirectURL:  "http://xxx.yyy.com",
			},
			want: &gbpConfig{
				clientID:     "dummy_id",
				secret:       "dummy_secret",
				refreshToken: "dummy_refresh_token",
				redirectURL:  "http://xxx.yyy.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGBPConfigEnv(tt.args.clientID, tt.args.secret, tt.args.refreshToken, tt.args.redirectURL)
			assert.Equal(t, tt.want, got, "NewGBPConfigEnv() = %v, want %v", got, tt.want)
		})
	}
}

func TestLoadGBPConfigEnv(t *testing.T) {
	type fields struct {
		clientID     string
		secret       string
		refreshToken string
		redirectURL  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    gbpapi.GBPConfig
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Load env and return GBPConfig",
			fields: fields{
				clientID:     "dummy_id",
				secret:       "dummy_secret",
				refreshToken: "dummy_refresh_token",
				redirectURL:  "http://xxx.yyy.com",
			},
			want: &gbpConfig{
				clientID:     "dummy_id",
				secret:       "dummy_secret",
				refreshToken: "dummy_refresh_token",
				redirectURL:  "http://xxx.yyy.com",
			},
			wantErr: assert.NoError,
		},
		{
			name: "Return error when environment valirable GBP_CLIENT_ID is unset",
			fields: fields{
				clientID:     "",
				secret:       "dummy_secret",
				refreshToken: "dummy_refresh_token",
				redirectURL:  "http://xxx.yyy.com",
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "failed to GBP configuration: "
				return assert.Contains(tt, e.Error(), want, "LoadGBPConfigEnv() error = %v, wantErr %v", e, want)
			},
		},
		{
			name: "Return error when environment valirable GBP_CLIENT_SECRET is unset",
			fields: fields{
				clientID:     "dummy_id",
				secret:       "",
				refreshToken: "dummy_refresh_token",
				redirectURL:  "http://xxx.yyy.com",
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "failed to GBP configuration: "
				return assert.Contains(tt, e.Error(), want, "LoadGBPConfigEnv() error = %v, wantErr %v", e, want)
			},
		},
		{
			name: "Return error when environment valirable GBP_REFRESH_TOKEN is unset",
			fields: fields{
				clientID:     "dummy_id",
				secret:       "dummy_secret",
				refreshToken: "",
				redirectURL:  "http://xxx.yyy.com",
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "failed to GBP configuration: "
				return assert.Contains(tt, e.Error(), want, "LoadGBPConfigEnv() error = %v, wantErr %v", e, want)
			},
		},
		{
			name: "Return error when environment valirable GBP_REDIRECT_URL is unset",
			fields: fields{
				clientID:     "dummy_id",
				secret:       "dummy_secret",
				refreshToken: "dummy_refresh_token",
				redirectURL:  "",
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "failed to GBP configuration: "
				return assert.Contains(tt, e.Error(), want, "LoadGBPConfigEnv() error = %v, wantErr %v", e, want)
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
		if len(tt.fields.redirectURL) != 0 {
			os.Setenv("GBP_REDIRECT_URL", tt.fields.redirectURL)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadGBPConfigEnv()
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got, "LoadGBPConfigEnv() = %v, want %v", got, tt.want)
		})
		os.Clearenv()
	}
}

func Test_gbpConfig_ClientID(t *testing.T) {
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
			c := &gbpConfig{
				clientID:     tt.fields.clientID,
				secret:       tt.fields.secret,
				refreshToken: tt.fields.refreshToken,
			}
			got := c.ClientID()
			assert.Equal(t, tt.want, got, "gbpConfig.ClientID() = %v, want %v", got, tt.want)
		})
	}
}

func Test_gbpConfig_ClientSecret(t *testing.T) {
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
			c := &gbpConfig{
				clientID:     tt.fields.clientID,
				secret:       tt.fields.secret,
				refreshToken: tt.fields.refreshToken,
			}
			got := c.ClientSecret()
			assert.Equal(t, tt.want, got, "gbpConfig.ClientSecret() = %v, want %v", got, tt.want)
		})
	}
}

func Test_gbpConfig_RefreshToken(t *testing.T) {
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
			c := &gbpConfig{
				clientID:     tt.fields.clientID,
				secret:       tt.fields.secret,
				refreshToken: tt.fields.refreshToken,
			}
			got := c.RefreshToken()
			assert.Equal(t, tt.want, got, "gbpConfig.RefreshToken() = %v, want %v", got, tt.want)
		})
	}
}

func Test_gbpConfig_RedirectURL(t *testing.T) {
	type fields struct {
		clientID     string
		secret       string
		refreshToken string
		redirectURL  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Return redirect URL",
			fields: fields{
				redirectURL: "http://xxx.yyy.com",
			},
			want: "http://xxx.yyy.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &gbpConfig{
				clientID:     tt.fields.clientID,
				secret:       tt.fields.secret,
				refreshToken: tt.fields.refreshToken,
				redirectURL:  tt.fields.redirectURL,
			}
			if got := c.RedirectURL(); got != tt.want {
				t.Errorf("gbpConfig.RedirectURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
