package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/nekonenene/gin_quiz_app/common"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	conf  *oauth2.Config
	state string
)

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
	state = randToken()
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()

	// common.OkResponse(c, "url", getLoginURL(state))
	c.Redirect(302, getLoginURL(state))
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func callbackHandler(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("state") != c.Query("state") {
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

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("userinfo: %v", string(data))
	common.OkResponse(c, "response", string(data))
}

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}
