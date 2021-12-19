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

var _ gbpapi.GBPCredential = (*gbpCredential)(nil)

/** Credentials for GBP API **/

func NewGBPCredentialEnv(clientID string, secret string, refreshToken string) gbpapi.GBPCredential {
	return &gbpCredential{
		clientID:     clientID,
		secret:       secret,
		refreshToken: refreshToken,
	}
}

func LoadGBPCredentialEnv() (gbpapi.GBPCredential, error) {
	var env struct {
		ClientID     string `envconfig:"gbp_client_id" required:"true"`
		Secret       string `envconfig:"gbp_client_secret" required:"true"`
		RefreshToken string `envconfig:"gbp_refresh_token" required:"true"`
	}
	if err := LoadEnvToStruct("", &env); err != nil {
		return nil, fmt.Errorf("failed to GBP credentials: %w", err)
	}
	return &gbpCredential{
		clientID:     env.ClientID,
		secret:       env.Secret,
		refreshToken: env.RefreshToken,
	}, nil
}

type gbpCredential struct {
	clientID     string
	secret       string
	refreshToken string
}

func (c *gbpCredential) ClientID() string {
	return c.clientID
}

func (c *gbpCredential) ClientSecret() string {
	return c.secret
}

func (c *gbpCredential) RefreshToken() string {
	return c.refreshToken
}

/** Client option for GBP API **/

func newGBPOption(ctx context.Context, credentials gbpapi.GBPCredential) option.ClientOption {
	config := oauth2.Config{
		ClientID:     credentials.ClientID(),
		ClientSecret: credentials.ClientSecret(),
		Endpoint:     google.Endpoint,
		RedirectURL:  "",
		Scopes:       gbpScopes,
	}
	tokenSource := config.TokenSource(ctx, &oauth2.Token{
		RefreshToken: credentials.RefreshToken(),
	})

	return option.WithTokenSource(tokenSource)
}
