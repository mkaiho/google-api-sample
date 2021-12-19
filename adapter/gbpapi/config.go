package gbpapi

type GBPConfig interface {
	ClientID() string
	ClientSecret() string
	RefreshToken() string
	RedirectURL() string
}
