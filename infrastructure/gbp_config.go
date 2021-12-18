package infrastructure

import (
	"fmt"

	"github.com/mkaiho/google-api-sample/adapter/gbpapi"
)

var _ gbpapi.GBPCredential = (*gbpCredential)(nil)

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
