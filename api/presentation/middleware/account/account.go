package account

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
	"time"
)

type AccountMiddleware interface {
	AccountMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type accountMiddleware struct{}

func NewAccountMiddleware() AccountMiddleware {
	return &accountMiddleware{}
}

// AccountMiddleware トークンを検証する
func (a *accountMiddleware) AccountMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString, err := extractTokenFromHeader(c)
		if err != nil {
			return err
		}

		token, err := parseToken(tokenString)
		if err != nil {
			return err
		}

		if isValidToken(token) {
			return next(c)
		}

		return fmt.Errorf("Invalid token")
	}
}

func extractTokenFromHeader(c echo.Context) (string, error) {
	tokenString := strings.ReplaceAll(c.Request().Header.Get("Authorization"), "Bearer ", "")
	if tokenString == "" {
		return "", fmt.Errorf("Authorization header is missing")
	}

	return tokenString, nil
}

func parseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if signingMethod, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		} else if signingMethod != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", signingMethod.Alg())
		}

		return []byte(os.Getenv("AUTH_SECRET")), nil
	})
}

func isValidToken(token *jwt.Token) bool {
	claims, ok := token.Claims.(jwt.MapClaims)

	return ok && token.Valid && claims.VerifyExpiresAt(time.Now().Unix(), true)
}
