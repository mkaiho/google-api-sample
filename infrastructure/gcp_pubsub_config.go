package infrastructure

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

/** TODO: Refactoring **/

var gcpPubsubSocpes = []string{
	pubsub.ScopePubSub,
	pubsub.ScopeCloudPlatform,
}

/** Configurations for GCP **/

func NewGCPConfigEnv(
	clientID string,
	secret string,
	refreshToken string,
	redirectURL string,
	topic string,
	subscription string,
) gcpConfig {
	return gcpConfig{
		clientID:     clientID,
		secret:       secret,
		refreshToken: refreshToken,
		redirectURL:  redirectURL,
		topic:        topic,
		subscription: subscription,
	}
}

func LoadGCPConfigEnv() (*gcpConfig, error) {
	var env struct {
		ClientID     string `envconfig:"gcp_client_id" required:"true"`
		Secret       string `envconfig:"gcp_client_secret" required:"true"`
		RefreshToken string `envconfig:"gcp_refresh_token" required:"true"`
		RedirectURL  string `envconfig:"gcp_redirect_url" required:"true"`
		ProjectID    string `envconfig:"gcp_project_id" required:"true"`
		Topic        string `envconfig:"gcp_pubsub_topic" required:"true"`
		Subscription string `envconfig:"gcp_pubsub_subscription" required:"true"`
	}
	if err := LoadEnvToStruct("", &env); err != nil {
		return nil, fmt.Errorf("failed to GCP configuration: %w", err)
	}
	return &gcpConfig{
		clientID:     env.ClientID,
		secret:       env.Secret,
		refreshToken: env.RefreshToken,
		redirectURL:  env.RedirectURL,
		projectID:    env.ProjectID,
		topic:        env.Topic,
		subscription: env.Subscription,
	}, nil
}

type gcpConfig struct {
	clientID     string
	secret       string
	refreshToken string
	redirectURL  string
	projectID    string
	topic        string
	subscription string
}

func (c *gcpConfig) ClientID() string {
	return c.clientID
}

func (c *gcpConfig) ClientSecret() string {
	return c.secret
}

func (c *gcpConfig) RefreshToken() string {
	return c.refreshToken
}

func (c *gcpConfig) RedirectURL() string {
	return c.redirectURL
}

func (c *gcpConfig) ProjectID() string {
	return c.projectID
}

func (c *gcpConfig) PubsubTopic() string {
	return c.topic
}

func (c *gcpConfig) PubsubSubscription() string {
	return c.subscription
}

/** Client option for GCP **/

func newGCPOption(ctx context.Context, config gcpConfig) option.ClientOption {
	oauth2config := oauth2.Config{
		ClientID:     config.ClientID(),
		ClientSecret: config.ClientSecret(),
		Endpoint:     google.Endpoint,
		RedirectURL:  config.RedirectURL(),
		Scopes:       gcpPubsubSocpes,
	}
	tokenSource := oauth2config.TokenSource(ctx, &oauth2.Token{
		RefreshToken: config.RefreshToken(),
	})

	return option.WithTokenSource(tokenSource)
}
