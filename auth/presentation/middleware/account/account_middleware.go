package account

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
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
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		token, err := parseToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		if isValidToken(token) {
			return next(c)
		}

		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}
}

func extractTokenFromHeader(c echo.Context) (string, error) {
	tokenString := strings.ReplaceAll(c.Request().Header.Get("Authorization"), "Bearer ", "")
	if tokenString == "" {
		return "", errors.New("Authorization header is missing")
	}

	return tokenString, nil
}

func parseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if signingMethod, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method: " + fmt.Sprint(token.Header["alg"]))
		} else if signingMethod != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method: " + signingMethod.Alg())
		}

		return []byte(os.Getenv("AUTH_SECRET")), nil
	})
}

func isValidToken(token *jwt.Token) bool {
	claims, ok := token.Claims.(jwt.MapClaims)

	return ok && token.Valid && claims.VerifyExpiresAt(time.Now().Unix(), true)
}
