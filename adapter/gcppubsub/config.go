package gcppubsub

type GCPPubsubConfig interface {
	ProjectID() string
	ClientID() string
	ClientSecret() string
	RefreshToken() string
}
