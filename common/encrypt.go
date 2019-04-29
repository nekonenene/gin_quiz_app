package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

func CreateHash(key string) string {
	salt := os.Getenv("SECRET_SALT")
	hasher := md5.New()
	hasher.Write([]byte(key + salt))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Encrypt(text []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(CreateHash(key))) // cipher block の作成
	if err != nil {
		return nil, err
	}

	b := base64.StdEncoding.EncodeToString(text)     // Base64 Encoding をおこない文字列 b とする
	ciphertext := make([]byte, aes.BlockSize+len(b)) // 16 + b の長さの空文字列を作る

	iv := ciphertext[:aes.BlockSize] // 最初の16文字を iv とする
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)                // CFB = Ciphertext feedback
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b)) // 各バイトをXORする
	return ciphertext, nil
}

func Decrypt(text []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(CreateHash(key)))
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)

	plaintext, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
