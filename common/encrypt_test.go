package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHash(t *testing.T) {
	key1 := "key1"
	key2 := ""

	assert.Equal(t, CreateHash(key1), CreateHash(key1))
	assert.Equal(t, CreateHash(key2), CreateHash(key2))
	assert.NotEqual(t, CreateHash(key1), CreateHash(key2))
}

func TestEncryptAndDecrypt(t *testing.T) {
	text := "this_is_password"
	bytes := []byte(text)

	var encoded string
	var decoded []byte
	var err error

	key := "key_for_decrypt"
	encoded, err = Encrypt(bytes, key)
	assert.Nil(t, err)
	decoded, err = Decrypt(encoded, key)
	assert.Nil(t, err)
	assert.Equal(t, bytes, decoded)

	key = ""
	encoded, err = Encrypt(bytes, key)
	assert.Nil(t, err)
	decoded, err = Decrypt(encoded, key)
	assert.Nil(t, err)
	assert.Equal(t, bytes, decoded)

	bytes = []byte("")
	encoded, err = Encrypt(bytes, key)
	assert.Nil(t, err)
	decoded, err = Decrypt(encoded, key)
	assert.Nil(t, err)
	assert.Equal(t, bytes, decoded)
}
