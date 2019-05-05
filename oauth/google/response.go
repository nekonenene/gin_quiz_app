package google

import (
	"github.com/nekonenene/gin_quiz_app/repository/user"
)

const (
	provider = "google"
)

type OAuthResponse struct {
	ProviderID    string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	PictureURL    string `json:"picture"`
}

func (res *OAuthResponse) CreateUser() (user.User, error) {
	user := user.User{
		Provider:   provider,
		ProviderID: res.ProviderID,
		Name:       res.Email,
		Email:      res.Email,
		IsAdmin:    false,
	}

	return user.Create()
}
