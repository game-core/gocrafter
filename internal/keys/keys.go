package keys

import (
	"fmt"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// GenerateUserID UserIDを生成する
func GenerateUserID(shardKey string) (string, error) {
	uuid, err := gonanoid.New(20)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%s", shardKey, uuid), nil
}

// GetShardKeyByUserId ユーザIDからシャードキーを取得する
func GetShardKeyByUserId(userID string) string {
	return strings.Split(userID, ":")[0]
}

// GeneratePassword パスワードを生成する
func GeneratePassword() (string, error) {
	password, err := gonanoid.New(20)
	if err != nil {
		return "", err
	}

	return password, nil
}

// GenerateHashPassword ハッシュパスワードを生成する
func GenerateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// CheckPassword パスワードを検証する
func CheckPassword(password, hashPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); err != nil {
		return false
	}

	return true
}
