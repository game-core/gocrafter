// Package userAction ユーザーアカウント
package userAction

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/keys"
)

type userAccountDao struct {
	ShardConn *database.ShardConn
}

func NewUserAccountDao(conn *database.SqlHandler) userAction.UserAccountRepository {
	return &userAccountDao{
		ShardConn: conn.User,
	}
}

func (s *userAccountDao) Find(ctx context.Context, userId string) (*userAction.UserAccount, error) {
	t := NewUserAccount()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userAction.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountDao) FindOrNil(ctx context.Context, userId string) (*userAction.UserAccount, error) {
	t := NewUserAccount()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userAction.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountDao) FindList(ctx context.Context, userId string) (userAction.UserAccounts, error) {
	ts := NewUserAccounts()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userAction.NewUserAccounts()
	for _, t := range ts {
		ms = append(ms, userAction.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt))
	}

	return ms, nil
}

func (s *userAccountDao) Create(ctx context.Context, tx *gorm.DB, m *userAction.UserAccount) (*userAction.UserAccount, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	t := &UserAccount{
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

	return userAction.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountDao) CreateList(ctx context.Context, tx *gorm.DB, ms userAction.UserAccounts) (userAction.UserAccounts, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	fms := ms[0]
	for _, m := range ms {
		if m.UserId != fms.UserId {
			return nil, errors.NewError("userId is invalid")
		}
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(fms.UserId)].WriteConn
	}

	ts := NewUserAccounts()
	for _, m := range ms {
		t := &UserAccount{
			UserId:   m.UserId,
			Name:     m.Name,
			Password: m.Password,
			LoginAt:  m.LoginAt,
			LogoutAt: m.LogoutAt,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserAccount()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userAccountDao) Update(ctx context.Context, tx *gorm.DB, m *userAction.UserAccount) (*userAction.UserAccount, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	t := &UserAccount{
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

	return userAction.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountDao) Delete(ctx context.Context, tx *gorm.DB, m *userAction.UserAccount) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	res := conn.Model(NewUserAccount()).WithContext(ctx).Where("user_id = ?", m.UserId).Delete(NewUserAccount())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
