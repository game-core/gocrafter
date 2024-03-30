// Package commonRankingWorld ワールドランキング
package commonRankingWorld

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingWorld"
)

type commonRankingWorldDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
}

func NewCommonRankingWorldDao(conn *database.MysqlHandler) commonRankingWorld.CommonRankingWorldMysqlRepository {
	return &commonRankingWorldDao{
		ReadMysqlConn:  conn.Common.ReadMysqlConn,
		WriteMysqlConn: conn.Common.WriteMysqlConn,
	}
}

func (s *commonRankingWorldDao) Find(ctx context.Context, masterRankingId int64, userId string) (*commonRankingWorld.CommonRankingWorld, error) {
	t := NewCommonRankingWorld()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_ranking_id = ?", masterRankingId).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return commonRankingWorld.SetCommonRankingWorld(t.MasterRankingId, t.UserId, t.Score), nil
}

func (s *commonRankingWorldDao) FindOrNil(ctx context.Context, masterRankingId int64, userId string) (*commonRankingWorld.CommonRankingWorld, error) {
	t := NewCommonRankingWorld()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_ranking_id = ?", masterRankingId).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return commonRankingWorld.SetCommonRankingWorld(t.MasterRankingId, t.UserId, t.Score), nil
}

func (s *commonRankingWorldDao) FindByMasterRankingId(ctx context.Context, masterRankingId int64) (*commonRankingWorld.CommonRankingWorld, error) {
	t := NewCommonRankingWorld()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_ranking_id = ?", masterRankingId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return commonRankingWorld.SetCommonRankingWorld(t.MasterRankingId, t.UserId, t.Score), nil
}

func (s *commonRankingWorldDao) FindOrNilByMasterRankingId(ctx context.Context, masterRankingId int64) (*commonRankingWorld.CommonRankingWorld, error) {
	t := NewCommonRankingWorld()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_ranking_id = ?", masterRankingId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return commonRankingWorld.SetCommonRankingWorld(t.MasterRankingId, t.UserId, t.Score), nil
}

func (s *commonRankingWorldDao) FindList(ctx context.Context) (commonRankingWorld.CommonRankingWorlds, error) {
	ts := NewCommonRankingWorlds()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := commonRankingWorld.NewCommonRankingWorlds()
	for _, t := range ts {
		ms = append(ms, commonRankingWorld.SetCommonRankingWorld(t.MasterRankingId, t.UserId, t.Score))
	}

	return ms, nil
}

func (s *commonRankingWorldDao) FindListByMasterRankingId(ctx context.Context, masterRankingId int64) (commonRankingWorld.CommonRankingWorlds, error) {
	ts := NewCommonRankingWorlds()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_ranking_id = ?", masterRankingId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := commonRankingWorld.NewCommonRankingWorlds()
	for _, t := range ts {
		ms = append(ms, commonRankingWorld.SetCommonRankingWorld(t.MasterRankingId, t.UserId, t.Score))
	}

	return ms, nil
}

func (s *commonRankingWorldDao) Create(ctx context.Context, tx *gorm.DB, m *commonRankingWorld.CommonRankingWorld) (*commonRankingWorld.CommonRankingWorld, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &CommonRankingWorld{
		MasterRankingId: m.MasterRankingId,
		UserId:          m.UserId,
		Score:           m.Score,
	}
	res := conn.Model(NewCommonRankingWorld()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return commonRankingWorld.SetCommonRankingWorld(t.MasterRankingId, t.UserId, t.Score), nil
}

func (s *commonRankingWorldDao) CreateList(ctx context.Context, tx *gorm.DB, ms commonRankingWorld.CommonRankingWorlds) (commonRankingWorld.CommonRankingWorlds, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewCommonRankingWorlds()
	for _, m := range ms {
		t := &CommonRankingWorld{
			MasterRankingId: m.MasterRankingId,
			UserId:          m.UserId,
			Score:           m.Score,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewCommonRankingWorld()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *commonRankingWorldDao) Update(ctx context.Context, tx *gorm.DB, m *commonRankingWorld.CommonRankingWorld) (*commonRankingWorld.CommonRankingWorld, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &CommonRankingWorld{
		MasterRankingId: m.MasterRankingId,
		UserId:          m.UserId,
		Score:           m.Score,
	}
	res := conn.Model(NewCommonRankingWorld()).WithContext(ctx).Where("master_ranking_id = ?", m.MasterRankingId).Where("user_id = ?", m.UserId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return commonRankingWorld.SetCommonRankingWorld(t.MasterRankingId, t.UserId, t.Score), nil
}

func (s *commonRankingWorldDao) Delete(ctx context.Context, tx *gorm.DB, m *commonRankingWorld.CommonRankingWorld) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewCommonRankingWorld()).WithContext(ctx).Where("master_ranking_id = ?", m.MasterRankingId).Where("user_id = ?", m.UserId).Delete(NewCommonRankingWorld())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
