// Package userIdleBonus ユーザー放置ボーナス
package userIdleBonus

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/keys"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/userIdleBonus"
)

type userIdleBonusDao struct {
	ShardMysqlConn *database.ShardMysqlConn
}

func NewUserIdleBonusDao(conn *database.MysqlHandler) userIdleBonus.UserIdleBonusMysqlRepository {
	return &userIdleBonusDao{
		ShardMysqlConn: conn.User,
	}
}

func (s *userIdleBonusDao) Find(ctx context.Context, userId string, masterIdleBonusId int64) (*userIdleBonus.UserIdleBonus, error) {
	t := NewUserIdleBonus()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_idle_bonus_id = ?", masterIdleBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userIdleBonus.SetUserIdleBonus(t.UserId, t.MasterIdleBonusId, t.ReceivedAt), nil
}

func (s *userIdleBonusDao) FindOrNil(ctx context.Context, userId string, masterIdleBonusId int64) (*userIdleBonus.UserIdleBonus, error) {
	t := NewUserIdleBonus()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_idle_bonus_id = ?", masterIdleBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userIdleBonus.SetUserIdleBonus(t.UserId, t.MasterIdleBonusId, t.ReceivedAt), nil
}

func (s *userIdleBonusDao) FindByUserId(ctx context.Context, userId string) (*userIdleBonus.UserIdleBonus, error) {
	t := NewUserIdleBonus()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userIdleBonus.SetUserIdleBonus(t.UserId, t.MasterIdleBonusId, t.ReceivedAt), nil
}

func (s *userIdleBonusDao) FindByUserIdAndMasterIdleBonusId(ctx context.Context, userId string, masterIdleBonusId int64) (*userIdleBonus.UserIdleBonus, error) {
	t := NewUserIdleBonus()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_idle_bonus_id = ?", masterIdleBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userIdleBonus.SetUserIdleBonus(t.UserId, t.MasterIdleBonusId, t.ReceivedAt), nil
}

func (s *userIdleBonusDao) FindOrNilByUserId(ctx context.Context, userId string) (*userIdleBonus.UserIdleBonus, error) {
	t := NewUserIdleBonus()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userIdleBonus.SetUserIdleBonus(t.UserId, t.MasterIdleBonusId, t.ReceivedAt), nil
}

func (s *userIdleBonusDao) FindOrNilByUserIdAndMasterIdleBonusId(ctx context.Context, userId string, masterIdleBonusId int64) (*userIdleBonus.UserIdleBonus, error) {
	t := NewUserIdleBonus()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_idle_bonus_id = ?", masterIdleBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userIdleBonus.SetUserIdleBonus(t.UserId, t.MasterIdleBonusId, t.ReceivedAt), nil
}

func (s *userIdleBonusDao) FindList(ctx context.Context, userId string) (userIdleBonus.UserIdleBonuses, error) {
	ts := NewUserIdleBonuses()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userIdleBonus.NewUserIdleBonuses()
	for _, t := range ts {
		ms = append(ms, userIdleBonus.SetUserIdleBonus(t.UserId, t.MasterIdleBonusId, t.ReceivedAt))
	}

	return ms, nil
}

func (s *userIdleBonusDao) FindListByUserId(ctx context.Context, userId string) (userIdleBonus.UserIdleBonuses, error) {
	ts := NewUserIdleBonuses()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userIdleBonus.NewUserIdleBonuses()
	for _, t := range ts {
		ms = append(ms, userIdleBonus.SetUserIdleBonus(t.UserId, t.MasterIdleBonusId, t.ReceivedAt))
	}

	return ms, nil
}

func (s *userIdleBonusDao) FindListByUserIdAndMasterIdleBonusId(ctx context.Context, userId string, masterIdleBonusId int64) (userIdleBonus.UserIdleBonuses, error) {
	ts := NewUserIdleBonuses()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_idle_bonus_id = ?", masterIdleBonusId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userIdleBonus.NewUserIdleBonuses()
	for _, t := range ts {
		ms = append(ms, userIdleBonus.SetUserIdleBonus(t.UserId, t.MasterIdleBonusId, t.ReceivedAt))
	}

	return ms, nil
}

func (s *userIdleBonusDao) Create(ctx context.Context, tx *gorm.DB, m *userIdleBonus.UserIdleBonus) (*userIdleBonus.UserIdleBonus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserIdleBonus{
		UserId:            m.UserId,
		MasterIdleBonusId: m.MasterIdleBonusId,
		ReceivedAt:        m.ReceivedAt,
	}
	res := conn.Model(NewUserIdleBonus()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userIdleBonus.SetUserIdleBonus(t.UserId, t.MasterIdleBonusId, t.ReceivedAt), nil
}

func (s *userIdleBonusDao) CreateList(ctx context.Context, tx *gorm.DB, ms userIdleBonus.UserIdleBonuses) (userIdleBonus.UserIdleBonuses, error) {
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

	ts := NewUserIdleBonuses()
	for _, m := range ms {
		t := &UserIdleBonus{
			UserId:            m.UserId,
			MasterIdleBonusId: m.MasterIdleBonusId,
			ReceivedAt:        m.ReceivedAt,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserIdleBonus()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userIdleBonusDao) Update(ctx context.Context, tx *gorm.DB, m *userIdleBonus.UserIdleBonus) (*userIdleBonus.UserIdleBonus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserIdleBonus{
		UserId:            m.UserId,
		MasterIdleBonusId: m.MasterIdleBonusId,
		ReceivedAt:        m.ReceivedAt,
	}
	res := conn.Model(NewUserIdleBonus()).WithContext(ctx).Where("user_id = ?", m.UserId).Where("master_idle_bonus_id = ?", m.MasterIdleBonusId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userIdleBonus.SetUserIdleBonus(t.UserId, t.MasterIdleBonusId, t.ReceivedAt), nil
}

func (s *userIdleBonusDao) Delete(ctx context.Context, tx *gorm.DB, m *userIdleBonus.UserIdleBonus) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	res := conn.Model(NewUserIdleBonus()).WithContext(ctx).Where("user_id = ?", m.UserId).Where("master_idle_bonus_id = ?", m.MasterIdleBonusId).Delete(NewUserIdleBonus())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
