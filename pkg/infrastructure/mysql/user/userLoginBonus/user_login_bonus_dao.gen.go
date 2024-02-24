// Package userLoginBonus ユーザーログインボーナス
package userLoginBonus

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/keys"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/userLoginBonus"
)

type userLoginBonusDao struct {
	ShardConn *database.ShardConn
}

func NewUserLoginBonusDao(conn *database.SqlHandler) userLoginBonus.UserLoginBonusRepository {
	return &userLoginBonusDao{
		ShardConn: conn.User,
	}
}

func (s *userLoginBonusDao) Find(ctx context.Context, userId string, masterLoginBonusId int64) (*userLoginBonus.UserLoginBonus, error) {
	t := NewUserLoginBonus()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_login_bonus_id = ?", masterLoginBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("record does not exist")
	}

	return userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt), nil
}

func (s *userLoginBonusDao) FindOrNil(ctx context.Context, userId string, masterLoginBonusId int64) (*userLoginBonus.UserLoginBonus, error) {
	t := NewUserLoginBonus()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_login_bonus_id = ?", masterLoginBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt), nil
}

func (s *userLoginBonusDao) FindByUserId(ctx context.Context, userId string) (*userLoginBonus.UserLoginBonus, error) {
	t := NewUserLoginBonus()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("record does not exist")
	}

	return userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt), nil
}

func (s *userLoginBonusDao) FindByMasterLoginBonusIdAndUserId(ctx context.Context, userId string, masterLoginBonusId int64) (*userLoginBonus.UserLoginBonus, error) {
	t := NewUserLoginBonus()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_login_bonus_id = ?", masterLoginBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("record does not exist")
	}

	return userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt), nil
}

func (s *userLoginBonusDao) FinOrNilByUserId(ctx context.Context, userId string) (*userLoginBonus.UserLoginBonus, error) {
	t := NewUserLoginBonus()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt), nil
}

func (s *userLoginBonusDao) FinOrNilByMasterLoginBonusIdAndUserId(ctx context.Context, userId string, masterLoginBonusId int64) (*userLoginBonus.UserLoginBonus, error) {
	t := NewUserLoginBonus()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_login_bonus_id = ?", masterLoginBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt), nil
}

func (s *userLoginBonusDao) FindList(ctx context.Context, userId string) (userLoginBonus.UserLoginBonuses, error) {
	ts := NewUserLoginBonuses()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userLoginBonus.NewUserLoginBonuses()
	for _, t := range ts {
		ms = append(ms, userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt))
	}

	return ms, nil
}

func (s *userLoginBonusDao) FindListByUserId(ctx context.Context, userId string) (userLoginBonus.UserLoginBonuses, error) {
	ts := NewUserLoginBonuses()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userLoginBonus.NewUserLoginBonuses()
	for _, t := range ts {
		ms = append(ms, userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt))
	}

	return ms, nil
}

func (s *userLoginBonusDao) FindListByMasterLoginBonusIdAndUserId(ctx context.Context, userId string, masterLoginBonusId int64) (userLoginBonus.UserLoginBonuses, error) {
	ts := NewUserLoginBonuses()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_login_bonus_id = ?", masterLoginBonusId).Find(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userLoginBonus.NewUserLoginBonuses()
	for _, t := range ts {
		ms = append(ms, userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt))
	}

	return ms, nil
}

func (s *userLoginBonusDao) Create(ctx context.Context, tx *gorm.DB, m *userLoginBonus.UserLoginBonus) (*userLoginBonus.UserLoginBonus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	t := &UserLoginBonus{
		UserId:             m.UserId,
		MasterLoginBonusId: m.MasterLoginBonusId,
		ReceivedAt:         m.ReceivedAt,
	}
	res := conn.Model(NewUserLoginBonus()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt), nil
}

func (s *userLoginBonusDao) CreateList(ctx context.Context, tx *gorm.DB, ms userLoginBonus.UserLoginBonuses) (userLoginBonus.UserLoginBonuses, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	fms := ms[0]
	for _, m := range ms {
		if m.UserId != fms.UserId {
			return nil, fmt.Errorf("userId is invalid")
		}
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(fms.UserId)].WriteConn
	}

	ts := NewUserLoginBonuses()
	for _, m := range ms {
		t := &UserLoginBonus{
			UserId:             m.UserId,
			MasterLoginBonusId: m.MasterLoginBonusId,
			ReceivedAt:         m.ReceivedAt,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserLoginBonus()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userLoginBonusDao) Update(ctx context.Context, tx *gorm.DB, m *userLoginBonus.UserLoginBonus) (*userLoginBonus.UserLoginBonus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	t := &UserLoginBonus{
		UserId:             m.UserId,
		MasterLoginBonusId: m.MasterLoginBonusId,
		ReceivedAt:         m.ReceivedAt,
	}
	res := conn.Model(NewUserLoginBonus()).WithContext(ctx).Where("user_id = ?", m.UserId).Where("master_login_bonus_id = ?", m.MasterLoginBonusId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt), nil
}

func (s *userLoginBonusDao) Delete(ctx context.Context, tx *gorm.DB, m *userLoginBonus.UserLoginBonus) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	res := conn.Model(NewUserLoginBonus()).WithContext(ctx).Where("user_id = ?", m.UserId).Where("master_login_bonus_id = ?", m.MasterLoginBonusId).Delete(NewUserLoginBonus())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
