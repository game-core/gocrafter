// Package userAccount ユーザーアカウント
package userAccount

import (
	"context"
	"fmt"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal"
	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccount"
)

type userAccountDao struct {
	ShardConn *database.ShardConn
}

func NewUserAccountDao(conn *database.SqlHandler) userAccount.UserAccountRepository {
	return &userAccountDao{
		ShardConn: conn.User,
	}
}

func (s *userAccountDao) Find(ctx context.Context, userId string) (*userAccount.UserAccount, error) {
	t := NewUserAccount()
	res := s.ShardConn.Shards[internal.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("record does not exist")
	}

	return &userAccount.UserAccount{
		UserId:   t.UserId,
		Name:     t.Name,
		Password: t.Password,
		LoginAt:  t.LoginAt,
		LogoutAt: t.LogoutAt,
	}, nil
}
