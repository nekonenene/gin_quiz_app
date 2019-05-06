package cookie

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
func Set(c *gin.Context, name string, value string, maxAge int) http.Cookie {
	cookie := http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		Path:     cookiePath,
		MaxAge:   maxAge,
		Secure:   c.Request.URL.Scheme == "https",
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}

	http.SetCookie(c.Writer, &cookie)
	return cookie
}

// Ref: https://github.com/gin-gonic/gin/blob/893c6ca/context.go#L760-L767
func GetValue(c *gin.Context, name string) (string, error) {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return "", err
	}

	val, _ := url.QueryUnescape(cookie.Value)
	return val, nil
}
