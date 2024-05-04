package auth

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	accountService "github.com/game-core/gocrafter/pkg/domain/model/account"
)

type AuthInterceptor interface {
	JwtAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
}

type authInterceptor struct {
	accountService accountService.AccountService
}

func NewAuthInterceptor(
	accountService accountService.AccountService,
) AuthInterceptor {
	return &authInterceptor{
		accountService: accountService,
	}
}

// JwtAuth 認証
func (i *authInterceptor) JwtAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if i.isPublic(info.FullMethod) {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata is not provided")
	}

	claims, err := i.check(ctx, strings.ReplaceAll(strings.Join(md.Get("authorization"), " "), "Bearer ", ""))
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %s", err)
	}

	return handler(context.WithValue(ctx, "jwtClaims", claims), req)
}

// isPublic 認証しないpath
func (i *authInterceptor) isPublic(fullMethod string) bool {
	return fullMethod == "/proto.Account/Login" || fullMethod == "/proto.Account/Create" || fullMethod == "/proto.Health/Check"
}

// check JWTトークンの検証
func (i *authInterceptor) check(ctx context.Context, tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if signingMethod, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", fmt.Sprint(token.Header["alg"]))
		} else if signingMethod != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %s", signingMethod.Alg())
		}

		return []byte(os.Getenv("AUTH_SECRET")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to jwt.Parse: %s", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, fmt.Errorf("invalid token")
	}

	userAccountToken, err := i.accountService.CheckToken(ctx, accountService.SetAccountCheckTokenRequest(claims["userId"].(string)))
	if err != nil {
		return nil, fmt.Errorf("i.accountService.CheckToken: %s", err)
	}

	if tokenString != userAccountToken.UserAccountToken.Token {
		return nil, fmt.Errorf("invalid userAccountToken")
	}

	return claims, nil
}
