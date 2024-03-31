// Package userAccountToken ユーザーアカウントトークン
package userAccountToken

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/keys"
	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccountToken"
)

type userAccountTokenDao struct {
	ShardMysqlConn *database.ShardMysqlConn
}

func NewUserAccountTokenDao(conn *database.MysqlHandler) userAccountToken.UserAccountTokenMysqlRepository {
	return &userAccountTokenDao{
		ShardMysqlConn: conn.User,
	}
}

func (s *userAccountTokenDao) Find(ctx context.Context, userId string) (*userAccountToken.UserAccountToken, error) {
	t := NewUserAccountToken()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userAccountToken.SetUserAccountToken(t.UserId, t.Token), nil
}

func (s *userAccountTokenDao) FindOrNil(ctx context.Context, userId string) (*userAccountToken.UserAccountToken, error) {
	t := NewUserAccountToken()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userAccountToken.SetUserAccountToken(t.UserId, t.Token), nil
}

func (s *userAccountTokenDao) FindList(ctx context.Context, userId string) (userAccountToken.UserAccountTokens, error) {
	ts := NewUserAccountTokens()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userAccountToken.NewUserAccountTokens()
	for _, t := range ts {
		ms = append(ms, userAccountToken.SetUserAccountToken(t.UserId, t.Token))
	}

	return ms, nil
}

func (s *userAccountTokenDao) Create(ctx context.Context, tx *gorm.DB, m *userAccountToken.UserAccountToken) (*userAccountToken.UserAccountToken, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserAccountToken{
		UserId: m.UserId,
		Token:  m.Token,
	}
	res := conn.Model(NewUserAccountToken()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userAccountToken.SetUserAccountToken(t.UserId, t.Token), nil
}

func (s *userAccountTokenDao) CreateList(ctx context.Context, tx *gorm.DB, ms userAccountToken.UserAccountTokens) (userAccountToken.UserAccountTokens, error) {
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

	ts := NewUserAccountTokens()
	for _, m := range ms {
		t := &UserAccountToken{
			UserId: m.UserId,
			Token:  m.Token,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserAccountToken()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userAccountTokenDao) Update(ctx context.Context, tx *gorm.DB, m *userAccountToken.UserAccountToken) (*userAccountToken.UserAccountToken, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserAccountToken{
		UserId: m.UserId,
		Token:  m.Token,
	}
	res := conn.Model(NewUserAccountToken()).WithContext(ctx).Where("user_id = ?", m.UserId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userAccountToken.SetUserAccountToken(t.UserId, t.Token), nil
}

func (s *userAccountTokenDao) Delete(ctx context.Context, tx *gorm.DB, m *userAccountToken.UserAccountToken) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	res := conn.Model(NewUserAccountToken()).WithContext(ctx).Where("user_id = ?", m.UserId).Delete(NewUserAccountToken())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
