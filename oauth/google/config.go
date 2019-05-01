package google

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	baseURL      = "http://localhost:8013" // TODO: baseURL は環境変数からの取得がいいかも
	callbackPath = "/oauth/google/callback"
)

var (
	conf *oauth2.Config
)

func InitConf() {
	conf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  baseURL + callbackPath,
		// Scopes: https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		Scopes: []string{
			"email",
			"openid",
		},
		Endpoint: google.Endpoint,
	}
}
