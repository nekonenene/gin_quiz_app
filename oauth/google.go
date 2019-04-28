package oauth

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
		RedirectURL:  "http://127.0.0.1:8013/oauth/google/callback",
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
	// state = randToken()
	// session := sessions.Default(c)
	// session.Set("state", state)
	// session.Save()

	fmt.Printf("%v", getLoginURL(state))
	common.OkResponse(c, "url", getLoginURL(state))
}

// func randToken() string {
// 	b := make([]byte, 32)
// 	rand.Read(b)
// 	return base64.StdEncoding.EncodeToString(b)
// }

func callbackHandler(c *gin.Context) {
	context := context.Background()
	token, err := conf.Exchange(context, c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := conf.Client(context, token)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer email.Body.Close()

	data, _ := ioutil.ReadAll(email.Body)
	log.Println("Email body: ", string(data))
	c.Status(http.StatusOK)
}

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}
