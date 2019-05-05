package google

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	callbackPath = "/oauth/google/callback"
)

type Config struct {
	ClientID     string
	ClientSecret string
	BaseURL      string
}

func InitConf(conf Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		RedirectURL:  conf.BaseURL + callbackPath,
		// Scopes: https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		Scopes: []string{
			"email",
			"openid",
		},
		Endpoint: google.Endpoint,
	}
}
