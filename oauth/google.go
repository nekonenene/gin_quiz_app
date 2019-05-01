package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nekonenene/gin_quiz_app/common"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	stateCookieName = "google_oauth_state"
	stateMaxAge     = 600
)

var (
	conf *oauth2.Config
)

type OAuthResponse struct {
	OpenID        string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	PictureURL    string `json:"picture"`
}

func InitGoogleOAuth() {
	conf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8013/oauth/google/callback",
		// Scopes: https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		Scopes: []string{
			"email",
			"openid",
		},
		Endpoint: google.Endpoint,
	}
}

func GoogleRouter(router *gin.RouterGroup) {
	router.GET("/login", login)
	router.GET("/callback", callbackHandler)
}

func login(c *gin.Context) {
	state := randToken()
	common.SetCookie(c, stateCookieName, state, stateMaxAge)

	c.Redirect(302, conf.AuthCodeURL(state))
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func callbackHandler(c *gin.Context) {
	state, _ := common.GetCookieValue(c, stateCookieName)
	common.SetCookie(c, stateCookieName, "", -1) // Delete Cookie
	if state != c.Query("state") {
		common.BadRequestErrorResponse(c, "invalid session state")
		return
	}

	context := context.Background()
	token, err := conf.Exchange(context, c.Query("code"))
	if err != nil {
		common.BadRequestErrorResponse(c, err.Error())
		return
	}
	if !token.Valid() {
		common.BadRequestErrorResponse(c, "invalid token")
		return
	}

	client := conf.Client(context, token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		common.BadRequestErrorResponse(c, err.Error())
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	var res OAuthResponse
	json.Unmarshal(body, &res)
	// log.Printf("userinfo: %v\n\n", string(body))
	log.Printf("userinfo: %v\n\n", res)
	common.OkResponse(c, "response", string(body))
}
