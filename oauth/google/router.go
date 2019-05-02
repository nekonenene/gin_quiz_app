package google

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nekonenene/gin_quiz_app/common"
	"github.com/nekonenene/gin_quiz_app/user"
)

const (
	stateLength     = 48
	stateCookieName = "google_oauth_state"
	stateMaxAge     = 600
)

func GoogleOAuthRouter(router *gin.RouterGroup) {
	router.GET("/login", login)
	router.GET("/callback", callbackHandler)
}

func login(c *gin.Context) {
	state := common.RandomString(stateLength)
	common.SetCookie(c, stateCookieName, state, stateMaxAge)

	c.Redirect(302, conf.AuthCodeURL(state))
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

	// TODO: エラーの表示についてはあとで書く
	var u user.User
	u, err = user.FindByOpenID(provider, res.ProviderID)
	if gorm.IsRecordNotFoundError(err) {
		u, err = res.CreateUser()
		if err != nil {
			log.Println("ユーザーの作成に失敗しました")
		}
	} else if err != nil {
		log.Println("不明なエラーが発生しました")
	}

	log.Printf("user: %v\n", u)
	log.Printf("userinfo: %v\n", res)
	common.OkResponse(c, "response", string(body))
}
