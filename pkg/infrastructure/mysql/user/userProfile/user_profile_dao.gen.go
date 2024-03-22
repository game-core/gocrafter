// Package userProfile ユーザープロフィール
package userProfile

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/keys"
	"github.com/game-core/gocrafter/pkg/domain/model/profile/userProfile"
)

type userProfileDao struct {
	ShardConn *database.ShardConn
}

func NewUserProfileDao(conn *database.MysqlHandler) userProfile.UserProfileRepository {
	return &userProfileDao{
		ShardConn: conn.User,
	}
}

func (s *userProfileDao) Find(ctx context.Context, userId string) (*userProfile.UserProfile, error) {
	t := NewUserProfile()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userProfile.SetUserProfile(t.UserId, t.Name, t.Content), nil
}

func (s *userProfileDao) FindOrNil(ctx context.Context, userId string) (*userProfile.UserProfile, error) {
	t := NewUserProfile()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userProfile.SetUserProfile(t.UserId, t.Name, t.Content), nil
}

func (s *userProfileDao) FindList(ctx context.Context, userId string) (userProfile.UserProfiles, error) {
	ts := NewUserProfiles()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userProfile.NewUserProfiles()
	for _, t := range ts {
		ms = append(ms, userProfile.SetUserProfile(t.UserId, t.Name, t.Content))
	}

	return ms, nil
}

func (s *userProfileDao) Create(ctx context.Context, tx *gorm.DB, m *userProfile.UserProfile) (*userProfile.UserProfile, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	t := &UserProfile{
		UserId:  m.UserId,
		Name:    m.Name,
		Content: m.Content,
	}
	res := conn.Model(NewUserProfile()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userProfile.SetUserProfile(t.UserId, t.Name, t.Content), nil
}

func (s *userProfileDao) CreateList(ctx context.Context, tx *gorm.DB, ms userProfile.UserProfiles) (userProfile.UserProfiles, error) {
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

	ts := NewUserProfiles()
	for _, m := range ms {
		t := &UserProfile{
			UserId:  m.UserId,
			Name:    m.Name,
			Content: m.Content,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserProfile()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userProfileDao) Update(ctx context.Context, tx *gorm.DB, m *userProfile.UserProfile) (*userProfile.UserProfile, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	t := &UserProfile{
		UserId:  m.UserId,
		Name:    m.Name,
		Content: m.Content,
	}
	res := conn.Model(NewUserProfile()).WithContext(ctx).Where("user_id = ?", m.UserId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userProfile.SetUserProfile(t.UserId, t.Name, t.Content), nil
}

func (s *userProfileDao) Delete(ctx context.Context, tx *gorm.DB, m *userProfile.UserProfile) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	res := conn.Model(NewUserProfile()).WithContext(ctx).Where("user_id = ?", m.UserId).Delete(NewUserProfile())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
