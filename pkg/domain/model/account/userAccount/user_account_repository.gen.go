package userAccount

import (
	"context"

	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userAccount"
)

type UserAccountRepository interface {
	Find(ctx context.Context, userId string) (*userAccount.UserAccount, error)
}
