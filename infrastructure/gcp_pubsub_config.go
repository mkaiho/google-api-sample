package infrastructure

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/mkaiho/google-api-sample/adapter/gcppubsub"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

var gcpPubsubSocpes = []string{
	pubsub.ScopePubSub,
	pubsub.ScopeCloudPlatform,
}

/** Configurations for GCP Pub/Sub **/

func NewGCPPubsubConfigEnv(
	clientID string,
	secret string,
	refreshToken string,
	topic string,
	subscription string,
) gcpPubsubConfig {
	return gcpPubsubConfig{
		clientID:     clientID,
		secret:       secret,
		refreshToken: refreshToken,
		topic:        topic,
		subscription: subscription,
	}
}

func LoadGCPPubsubConfigEnv() (*gcpPubsubConfig, error) {
	var env struct {
		ProjectID    string `envconfig:"gcp_project_id" required:"true"`
		ClientID     string `envconfig:"gcp_client_id" required:"true"`
		Secret       string `envconfig:"gcp_client_secret" required:"true"`
		RefreshToken string `envconfig:"gcp_pubsub_refresh_token" required:"true"`
		Topic        string `envconfig:"gcp_pubsub_topic" required:"true"`
		Subscription string `envconfig:"gcp_pubsub_subscription" required:"true"`
	}
	if err := LoadEnvToStruct("", &env); err != nil {
		return nil, fmt.Errorf("failed to GCP configuration: %w", err)
	}
	return &gcpPubsubConfig{
		clientID:     env.ClientID,
		secret:       env.Secret,
		refreshToken: env.RefreshToken,
		projectID:    env.ProjectID,
		topic:        env.Topic,
		subscription: env.Subscription,
	}, nil
}

type gcpPubsubConfig struct {
	clientID     string
	secret       string
	refreshToken string
	redirectURL  string
	projectID    string
	topic        string
	subscription string
}

func (c *gcpPubsubConfig) ProjectID() string {
	return c.projectID
}

func (c *gcpPubsubConfig) ClientID() string {
	return c.clientID
}

func (c *gcpPubsubConfig) ClientSecret() string {
	return c.secret
}

func (c *gcpPubsubConfig) RefreshToken() string {
	return c.refreshToken
}

func (c *gcpPubsubConfig) Topic() string {
	return c.topic
}

func (c *gcpPubsubConfig) Subscription() string {
	return c.subscription
}

/** Client option for GCP **/

func newGCPPubsubOption(ctx context.Context, config gcppubsub.GCPPubsubConfig) option.ClientOption {
	oauth2config := oauth2.Config{
		ClientID:     config.ClientID(),
		ClientSecret: config.ClientSecret(),
		Endpoint:     google.Endpoint,
		Scopes:       gcpPubsubSocpes,
	}
	tokenSource := oauth2config.TokenSource(ctx, &oauth2.Token{
		RefreshToken: config.RefreshToken(),
	})

	return option.WithTokenSource(tokenSource)
}
