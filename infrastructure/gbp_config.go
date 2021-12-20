package infrastructure

import (
	"context"
	"fmt"

	"github.com/mkaiho/google-api-sample/adapter/gbpapi"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

var gbpScopes = []string{
	"https://www.googleapis.com/auth/business.manage",
}

var _ gbpapi.GBPConfig = (*gbpConfig)(nil)

/** Configurations for GBP API **/

func NewGBPConfigEnv(clientID string, secret string, refreshToken string) gbpapi.GBPConfig {
	return &gbpConfig{
		clientID:     clientID,
		secret:       secret,
		refreshToken: refreshToken,
	}
}

func LoadGBPConfigEnv() (gbpapi.GBPConfig, error) {
	var env struct {
		ClientID     string `envconfig:"gcp_client_id" required:"true"`
		Secret       string `envconfig:"gcp_client_secret" required:"true"`
		RefreshToken string `envconfig:"gbp_refresh_token" required:"true"`
	}
	if err := LoadEnvToStruct("", &env); err != nil {
		return nil, fmt.Errorf("failed to GBP configuration: %w", err)
	}
	return &gbpConfig{
		clientID:     env.ClientID,
		secret:       env.Secret,
		refreshToken: env.RefreshToken,
	}, nil
}

type gbpConfig struct {
	clientID     string
	secret       string
	refreshToken string
}

func (c *gbpConfig) ClientID() string {
	return c.clientID
}

func (c *gbpConfig) ClientSecret() string {
	return c.secret
}

func (c *gbpConfig) RefreshToken() string {
	return c.refreshToken
}

/** Client option for GBP API **/

func newGBPOption(ctx context.Context, config gbpapi.GBPConfig) option.ClientOption {
	oauth2config := oauth2.Config{
		ClientID:     config.ClientID(),
		ClientSecret: config.ClientSecret(),
		Endpoint:     google.Endpoint,
		Scopes:       gbpScopes,
	}
	tokenSource := oauth2config.TokenSource(ctx, &oauth2.Token{
		RefreshToken: config.RefreshToken(),
	})

	return option.WithTokenSource(tokenSource)
}
