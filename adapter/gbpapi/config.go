package gbpapi

type GBPCredential interface {
	ClientID() string
	ClientSecret() string
	RefreshToken() string
}
