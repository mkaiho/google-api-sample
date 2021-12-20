package infrastructure

import (
	"fmt"

	"github.com/mkaiho/google-api-sample/adapter/oauth2"
)

/** Configurations for GCP OAuth2 **/

func NewGCPOAuth2ConfigEnv(clientID string, secret string, refreshToken string, redirectURL string) oauth2.OAuth2Config {
	return &gcpOAuth2Config{
		clientID:     clientID,
		secret:       secret,
		refreshToken: refreshToken,
		redirectURL:  redirectURL,
	}
}

func LoadGCPOAuth2ConfigEnv() (oauth2.OAuth2Config, error) {
	var env struct {
		ClientID     string `envconfig:"gcp_client_id" required:"true"`
		Secret       string `envconfig:"gcp_client_secret" required:"true"`
		RefreshToken string `envconfig:"gbp_refresh_token" required:"true"`
		RedirectURL  string `envconfig:"gcp_oauth2_token_redirect_url" required:"true"`
	}
	if err := LoadEnvToStruct("", &env); err != nil {
		return nil, fmt.Errorf("failed to GBP configuration: %w", err)
	}
	return &gcpOAuth2Config{
		clientID:     env.ClientID,
		secret:       env.Secret,
		refreshToken: env.RefreshToken,
		redirectURL:  env.RedirectURL,
	}, nil
}

type gcpOAuth2Config struct {
	clientID     string
	secret       string
	refreshToken string
	redirectURL  string
}

func (c *gcpOAuth2Config) ClientID() string {
	return c.clientID
}

func (c *gcpOAuth2Config) ClientSecret() string {
	return c.secret
}

func (c *gcpOAuth2Config) RefreshToken() string {
	return c.refreshToken
}

func (c *gcpOAuth2Config) RedirectURL() string {
	return c.redirectURL
}
