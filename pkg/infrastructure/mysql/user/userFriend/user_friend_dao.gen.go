// Package userFriend ユーザーフレンド
package userFriend

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/keys"
	"github.com/game-core/gocrafter/pkg/domain/model/friend/userFriend"
)

type userFriendDao struct {
	ShardMysqlConn *database.ShardMysqlConn
}

func NewUserFriendDao(conn *database.MysqlHandler) userFriend.UserFriendMysqlRepository {
	return &userFriendDao{
		ShardMysqlConn: conn.User,
	}
}

func (s *userFriendDao) Find(ctx context.Context, userId string, friendUserId string) (*userFriend.UserFriend, error) {
	t := NewUserFriend()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("friend_user_id = ?", friendUserId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userFriend.SetUserFriend(t.UserId, t.FriendUserId, t.FriendType), nil
}

func (s *userFriendDao) FindOrNil(ctx context.Context, userId string, friendUserId string) (*userFriend.UserFriend, error) {
	t := NewUserFriend()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("friend_user_id = ?", friendUserId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userFriend.SetUserFriend(t.UserId, t.FriendUserId, t.FriendType), nil
}

func (s *userFriendDao) FindList(ctx context.Context, userId string) (userFriend.UserFriends, error) {
	ts := NewUserFriends()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userFriend.NewUserFriends()
	for _, t := range ts {
		ms = append(ms, userFriend.SetUserFriend(t.UserId, t.FriendUserId, t.FriendType))
	}

	return ms, nil
}

func (s *userFriendDao) Create(ctx context.Context, tx *gorm.DB, m *userFriend.UserFriend) (*userFriend.UserFriend, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserFriend{
		UserId:       m.UserId,
		FriendUserId: m.FriendUserId,
		FriendType:   m.FriendType,
	}
	res := conn.Model(NewUserFriend()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userFriend.SetUserFriend(t.UserId, t.FriendUserId, t.FriendType), nil
}

func (s *userFriendDao) CreateList(ctx context.Context, tx *gorm.DB, ms userFriend.UserFriends) (userFriend.UserFriends, error) {
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

	ts := NewUserFriends()
	for _, m := range ms {
		t := &UserFriend{
			UserId:       m.UserId,
			FriendUserId: m.FriendUserId,
			FriendType:   m.FriendType,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserFriend()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userFriendDao) Update(ctx context.Context, tx *gorm.DB, m *userFriend.UserFriend) (*userFriend.UserFriend, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserFriend{
		UserId:       m.UserId,
		FriendUserId: m.FriendUserId,
		FriendType:   m.FriendType,
	}
	res := conn.Model(NewUserFriend()).WithContext(ctx).Where("user_id = ?", m.UserId).Where("friend_user_id = ?", m.FriendUserId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userFriend.SetUserFriend(t.UserId, t.FriendUserId, t.FriendType), nil
}

func (s *userFriendDao) Delete(ctx context.Context, tx *gorm.DB, m *userFriend.UserFriend) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	res := conn.Model(NewUserFriend()).WithContext(ctx).Where("user_id = ?", m.UserId).Where("friend_user_id = ?", m.FriendUserId).Delete(NewUserFriend())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
