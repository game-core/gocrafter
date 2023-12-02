package token

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateAuthTokenByUUID(uuid, name string) (string, error) {
	baseToken := jwt.New(jwt.SigningMethodHS256)
	claims := baseToken.Claims.(jwt.MapClaims)
	claims["uuid"] = uuid
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token, err := baseToken.SignedString([]byte(os.Getenv("AUTH_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateAuthTokenByEmail(email, name string) (string, error) {
	baseToken := jwt.New(jwt.SigningMethodHS256)
	claims := baseToken.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token, err := baseToken.SignedString([]byte(os.Getenv("AUTH_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}
