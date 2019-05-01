package common

import (
	crand "crypto/rand"
	"encoding/base64"
	"io"
)

func RandomString(length int) string {
	b := make([]byte, length)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}

	str := base64.URLEncoding.EncodeToString(b)
	return str[:length]
}
