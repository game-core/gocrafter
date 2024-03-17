// Package userAction ユーザーアクション
package userAction

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/keys"
	"github.com/game-core/gocrafter/pkg/domain/model/action/userAction"
)

type userActionDao struct {
	ShardConn *database.ShardConn
}

func NewUserActionDao(conn *database.SqlHandler) userAction.UserActionRepository {
	return &userActionDao{
		ShardConn: conn.User,
	}
}

func (s *userActionDao) Find(ctx context.Context, userId string, masterActionId int64) (*userAction.UserAction, error) {
	t := NewUserAction()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_action_id = ?", masterActionId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userAction.SetUserAction(t.UserId, t.Name, t.MasterActionId, t.StartedAt), nil
}

func (s *userActionDao) FindOrNil(ctx context.Context, userId string, masterActionId int64) (*userAction.UserAction, error) {
	t := NewUserAction()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_action_id = ?", masterActionId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userAction.SetUserAction(t.UserId, t.Name, t.MasterActionId, t.StartedAt), nil
}

func (s *userActionDao) FindList(ctx context.Context, userId string) (userAction.UserActions, error) {
	ts := NewUserActions()
	res := s.ShardConn.Shards[keys.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userAction.NewUserActions()
	for _, t := range ts {
		ms = append(ms, userAction.SetUserAction(t.UserId, t.Name, t.MasterActionId, t.StartedAt))
	}

	return ms, nil
}

func (s *userActionDao) Create(ctx context.Context, tx *gorm.DB, m *userAction.UserAction) (*userAction.UserAction, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	t := &UserAction{
		UserId:         m.UserId,
		Name:           m.Name,
		MasterActionId: m.MasterActionId,
		StartedAt:      m.StartedAt,
	}
	res := conn.Model(NewUserAction()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userAction.SetUserAction(t.UserId, t.Name, t.MasterActionId, t.StartedAt), nil
}

func (s *userActionDao) CreateList(ctx context.Context, tx *gorm.DB, ms userAction.UserActions) (userAction.UserActions, error) {
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

	ts := NewUserActions()
	for _, m := range ms {
		t := &UserAction{
			UserId:         m.UserId,
			Name:           m.Name,
			MasterActionId: m.MasterActionId,
			StartedAt:      m.StartedAt,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserAction()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userActionDao) Update(ctx context.Context, tx *gorm.DB, m *userAction.UserAction) (*userAction.UserAction, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	t := &UserAction{
		UserId:         m.UserId,
		Name:           m.Name,
		MasterActionId: m.MasterActionId,
		StartedAt:      m.StartedAt,
	}
	res := conn.Model(NewUserAction()).WithContext(ctx).Where("user_id = ?", m.UserId).Where("master_action_id = ?", m.MasterActionId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userAction.SetUserAction(t.UserId, t.Name, t.MasterActionId, t.StartedAt), nil
}

func (s *userActionDao) Delete(ctx context.Context, tx *gorm.DB, m *userAction.UserAction) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteConn
	}

	res := conn.Model(NewUserAction()).WithContext(ctx).Where("user_id = ?", m.UserId).Where("master_action_id = ?", m.MasterActionId).Delete(NewUserAction())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
