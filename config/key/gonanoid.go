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
