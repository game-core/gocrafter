package token

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

func GenerateAuthToken(uuid, name string) (string, error) {
	baseToken := jwt.New(jwt.SigningMethodHS256)
	claims := baseToken.Claims.(jwt.MapClaims)
	claims["uuid"] = uuid
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	
	token, err := baseToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
