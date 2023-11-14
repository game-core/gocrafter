package middleware

import (
	"os"
	"fmt"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type UserMiddleware interface {
	UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type userMiddleware struct {}

func NewUserMiddleware() UserMiddleware {
    return &userMiddleware{}
}

// UserMiddleware トークンを検証する
func (userMiddleware *userMiddleware) UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
		if tokenString == "" {
			return fmt.Errorf("Invalid token")
		}
	
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
		} else {
			return fmt.Errorf("Invalid token")
		}
    }
}
