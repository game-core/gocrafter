// Package userAccount ユーザーアカウント
package userAccount

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/keys"
	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccount"
)

type userAccountDao struct {
	ShardMysqlConn *database.ShardMysqlConn
}

func NewUserAccountDao(conn *database.MysqlHandler) userAccount.UserAccountRepository {
	return &userAccountDao{
		ShardMysqlConn: conn.User,
	}
}

func (s *userAccountDao) Find(ctx context.Context, userId string) (*userAccount.UserAccount, error) {
	t := NewUserAccount()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountDao) FindOrNil(ctx context.Context, userId string) (*userAccount.UserAccount, error) {
	t := NewUserAccount()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
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
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
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
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
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

	return userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountDao) CreateList(ctx context.Context, tx *gorm.DB, ms userAccount.UserAccounts) (userAccount.UserAccounts, error) {
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
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(fms.UserId)].WriteMysqlConn
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

func (s *userAccountDao) Update(ctx context.Context, tx *gorm.DB, m *userAccount.UserAccount) (*userAccount.UserAccount, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
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

	return userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountDao) Delete(ctx context.Context, tx *gorm.DB, m *userAccount.UserAccount) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	res := conn.Model(NewUserAccount()).WithContext(ctx).Where("user_id = ?", m.UserId).Delete(NewUserAccount())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
