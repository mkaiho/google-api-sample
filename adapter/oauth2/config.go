package oauth2

type OAuth2Config interface {
	ClientID() string
	ClientSecret() string
	RefreshToken() string
	RedirectURL() string
}
