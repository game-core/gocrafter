// Package userAccount ユーザーアカウント
package userAccount

import (
	"context"
	"fmt"

	"gorm.io/gorm"

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

	return userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountDao) FindOrNil(ctx context.Context, userId string) (*userAccount.UserAccount, error) {
	t := NewUserAccount()
	res := s.ShardConn.Shards[internal.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountDao) FindList(ctx context.Context, userId string) (userAccount.UserAccounts, error) {
	ts := NewUserAccounts()
	res := s.ShardConn.Shards[internal.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userAccount.NewUserAccounts()
	for _, t := range ts {
		ms = append(ms, userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt))
	}

	return ms, nil
}

func (s *userAccountDao) Create(ctx context.Context, tx *gorm.DB, m *userAccount.UserAccount) (*userAccount.UserAccount, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[internal.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	t := &userAccount.UserAccount{
		UserId:   m.UserId,
		Name:     m.Name,
		Password: m.Password,
		LoginAt:  m.LoginAt,
		LogoutAt: m.LogoutAt,
	}
	res := conn.Model(NewUserAccount()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountDao) Update(ctx context.Context, tx *gorm.DB, m *userAccount.UserAccount) (*userAccount.UserAccount, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[internal.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	t := &userAccount.UserAccount{
		UserId:   m.UserId,
		Name:     m.Name,
		Password: m.Password,
		LoginAt:  m.LoginAt,
		LogoutAt: m.LogoutAt,
	}
	res := conn.Model(NewUserAccount()).WithContext(ctx).Where("user_id = ?", m.UserId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountDao) Delete(ctx context.Context, tx *gorm.DB, m *userAccount.UserAccount) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[internal.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	res := conn.Model(NewUserAccount()).WithContext(ctx).Where("user_id = ?", m.UserId).Delete(NewUserAccount())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
