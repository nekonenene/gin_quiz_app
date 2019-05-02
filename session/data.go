package session

import (
	"encoding/json"
	"os"

	"github.com/nekonenene/gin_quiz_app/common"
)

type Data struct {
	UserID uint64 `json:"user_id"`
}

func (data *Data) Encode() (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return common.Encrypt(bytes, os.Getenv("SESSION_SECRET_KEY"))
}

func Decode(str string) (Data, error) {
	decryptedBytes, err := common.Decrypt(str, os.Getenv("SESSION_SECRET_KEY"))
	if err != nil {
		return Data{}, err
	}

	var data Data
	if err := json.Unmarshal(decryptedBytes, &data); err != nil {
		return Data{}, err
	}

	return data, nil
}
