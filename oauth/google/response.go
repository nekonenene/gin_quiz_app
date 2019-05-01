package google

type OAuthResponse struct {
	OpenID        string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	PictureURL    string `json:"picture"`
}
