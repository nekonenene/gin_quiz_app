package google

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nekonenene/gin_quiz_app/common"
	"github.com/nekonenene/gin_quiz_app/session"
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
	// 遷移異常がないか確認
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

	// Google の OAuth 2.0 API から情報を取得
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

	// Google OAuth の結果から user を検索。なければ作成
	var u user.User
	u, err = user.FindByOpenID(provider, res.ProviderID)
	if gorm.IsRecordNotFoundError(err) {
		u, err = res.CreateUser()
		if err != nil {
			common.ErrorResponse(c, err.Error())
			return
		}
	} else if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}

	// User ID をセッションに保存
	data := session.Data{UserID: u.ID}
	encoded, err := data.Encode()
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}
	session.DestroySession(c)
	session.StartNewSession(c, encoded)

	c.Redirect(302, "/")
}
