package common

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

// length byte のデータを Base64 Encode した後の左から length 文字を得る
func RandomString(length int) string {
	b := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	str := base64.URLEncoding.EncodeToString(b)
	return str[:length]
}
