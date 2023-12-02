package key

import (
	"golang.org/x/crypto/bcrypt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// GenerateUUID UUIDを生成する
func GenerateUUID() (string, error) {
	uuid, err := gonanoid.New(20)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

// GeneratePassword パスワードを生成する
func GeneratePassword() (string, string, error) {
	password, err := gonanoid.New(20)
	if err != nil {
		return "", "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	return password, string(hashedPassword), nil
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
