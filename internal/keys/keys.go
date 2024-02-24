package keys

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/game-core/gocrafter/internal/errors"
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
