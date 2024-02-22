package userAccount

import (
	"context"
)

type UserAccountRepository interface {
	Find(ctx context.Context, userId string) (*UserAccount, error)
}
