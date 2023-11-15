package middleware

import (
	"os"
	"fmt"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type AccountMiddleware interface {
	AccountMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type accountMiddleware struct {}

func NewAccountMiddleware() AccountMiddleware {
    return &accountMiddleware{}
}

// AccountMiddleware トークンを検証する
func (a *accountMiddleware) AccountMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")	
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("AUTH_SECRET")), nil
		})
		if err != nil {
			return fmt.Errorf("Invalid token")
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return next(c)
		}

		return fmt.Errorf("Invalid token")
    }
}
