// Package userItemBox ユーザーアイテムボックス
package userItemBox

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/keys"
	"github.com/game-core/gocrafter/pkg/domain/model/item/userItemBox"
)

type userItemBoxDao struct {
	ShardConn *database.ShardConn
}

func NewUserItemBoxDao(conn *database.SqlHandler) userItemBox.UserItemBoxRepository {
	return &userItemBoxDao{
		ShardConn: conn.User,
	}
}

func (s *userItemBoxDao) Find(ctx context.Context, userId string, masterItemId int64) (*userItemBox.UserItemBox, error) {
	t := NewUserItemBox()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_item_id = ?", masterItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userItemBox.SetUserItemBox(t.UserId, t.MasterItemId, t.Count), nil
}

func (s *userItemBoxDao) FindOrNil(ctx context.Context, userId string, masterItemId int64) (*userItemBox.UserItemBox, error) {
	t := NewUserItemBox()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_item_id = ?", masterItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userItemBox.SetUserItemBox(t.UserId, t.MasterItemId, t.Count), nil
}

func (s *userItemBoxDao) FindList(ctx context.Context, userId string) (userItemBox.UserItemBoxes, error) {
	ts := NewUserItemBoxes()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userItemBox.NewUserItemBoxes()
	for _, t := range ts {
		ms = append(ms, userItemBox.SetUserItemBox(t.UserId, t.MasterItemId, t.Count))
	}

	return ms, nil
}

func (s *userItemBoxDao) Create(ctx context.Context, tx *gorm.DB, m *userItemBox.UserItemBox) (*userItemBox.UserItemBox, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	t := &UserItemBox{
		UserId:       m.UserId,
		MasterItemId: m.MasterItemId,
		Count:        m.Count,
	}
	res := conn.Model(NewUserItemBox()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userItemBox.SetUserItemBox(t.UserId, t.MasterItemId, t.Count), nil
}

func (s *userItemBoxDao) CreateList(ctx context.Context, tx *gorm.DB, ms userItemBox.UserItemBoxes) (userItemBox.UserItemBoxes, error) {
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

	ts := NewUserItemBoxes()
	for _, m := range ms {
		t := &UserItemBox{
			UserId:       m.UserId,
			MasterItemId: m.MasterItemId,
			Count:        m.Count,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserItemBox()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userItemBoxDao) Update(ctx context.Context, tx *gorm.DB, m *userItemBox.UserItemBox) (*userItemBox.UserItemBox, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	t := &UserItemBox{
		UserId:       m.UserId,
		MasterItemId: m.MasterItemId,
		Count:        m.Count,
	}
	res := conn.Model(NewUserItemBox()).WithContext(ctx).Where("user_id = ?", m.UserId).Where("master_item_id = ?", m.MasterItemId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userItemBox.SetUserItemBox(t.UserId, t.MasterItemId, t.Count), nil
}

func (s *userItemBoxDao) Delete(ctx context.Context, tx *gorm.DB, m *userItemBox.UserItemBox) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	res := conn.Model(NewUserItemBox()).WithContext(ctx).Where("user_id = ?", m.UserId).Where("master_item_id = ?", m.MasterItemId).Delete(NewUserItemBox())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
