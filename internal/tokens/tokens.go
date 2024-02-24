package tokens

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"strings"
	"time"

	"github.com/game-core/gocrafter/internal/errors"
)

// GenerateAuthTokenByUserId userIdからAuthトークンを発行する
func GenerateAuthTokenByUserId(userId, name string) (string, error) {
	baseToken := jwt.New(jwt.SigningMethodHS256)
	claims := baseToken.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token, err := baseToken.SignedString([]byte(os.Getenv("AUTH_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

// CheckAuthToken Authトークンを検証する
func CheckAuthToken(userId, token string) error {
	tokenUserId, err := extractUserIDFromAuthToken(token)
	if err != nil {
		return err
	}

	if userId != tokenUserId {
		return fmt.Errorf("userId or token are invalid")
	}

	return nil
}

func CheckJwtClaims(ctx context.Context, userId string) error {
	jwtClaims, ok := ctx.Value("jwtClaims").(map[string]interface{})
	if !ok {
		return errors.NewError("failed to get jwtClaims from context")
	}
	jwtUserId, ok := jwtClaims["userId"].(string)
	if !ok {
		return errors.NewError("failed to get userId from jwtClaims")
	}
	if userId != jwtUserId {
		return errors.NewError("userId is invalid")
	}

	return nil
}

// extractUserIDFromAuthToken AuthトークンからuserIdを抽出する
func extractUserIDFromAuthToken(tokenString string) (string, error) {
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("AUTH_SECRET")), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.NewError("invalid token claims")
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		return "", errors.NewError("userId not found in token claims")
	}

	return userId, nil
}
