package common

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const (
	cookiePath = "/"
)

// cookie に保存。maxAge に負の値が渡されたときは削除がおこなわれる
// Ref: https://golang.org/pkg/net/http/#Cookie
func SetCookie(c *gin.Context, name string, value string, maxAge int) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		Path:     cookiePath,
		MaxAge:   maxAge,
		Secure:   c.Request.URL.Scheme == "https",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}

// Ref: https://github.com/gin-gonic/gin/blob/893c6ca/context.go#L760-L767
func GetCookie(c *gin.Context, name string) (string, error) {
	return c.Cookie(name)
}
