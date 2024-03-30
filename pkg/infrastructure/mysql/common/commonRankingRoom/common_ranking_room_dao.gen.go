// Package commonRankingRoom ルームランキング
package commonRankingRoom

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingRoom"
)

type commonRankingRoomDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
}

func NewCommonRankingRoomDao(conn *database.MysqlHandler) commonRankingRoom.CommonRankingRoomMysqlRepository {
	return &commonRankingRoomDao{
		ReadMysqlConn:  conn.Common.ReadMysqlConn,
		WriteMysqlConn: conn.Common.WriteMysqlConn,
	}
}

func (s *commonRankingRoomDao) Find(ctx context.Context, masterRankingId int64, roomId string, userId string) (*commonRankingRoom.CommonRankingRoom, error) {
	t := NewCommonRankingRoom()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_ranking_id = ?", masterRankingId).Where("room_id = ?", roomId).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return commonRankingRoom.SetCommonRankingRoom(t.MasterRankingId, t.RoomId, t.UserId, t.Score), nil
}

func (s *commonRankingRoomDao) FindOrNil(ctx context.Context, masterRankingId int64, roomId string, userId string) (*commonRankingRoom.CommonRankingRoom, error) {
	t := NewCommonRankingRoom()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_ranking_id = ?", masterRankingId).Where("room_id = ?", roomId).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return commonRankingRoom.SetCommonRankingRoom(t.MasterRankingId, t.RoomId, t.UserId, t.Score), nil
}

func (s *commonRankingRoomDao) FindByMasterRankingIdAndRoomId(ctx context.Context, masterRankingId int64, roomId string) (*commonRankingRoom.CommonRankingRoom, error) {
	t := NewCommonRankingRoom()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_ranking_id = ?", masterRankingId).Where("room_id = ?", roomId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return commonRankingRoom.SetCommonRankingRoom(t.MasterRankingId, t.RoomId, t.UserId, t.Score), nil
}

func (s *commonRankingRoomDao) FindOrNilByMasterRankingIdAndRoomId(ctx context.Context, masterRankingId int64, roomId string) (*commonRankingRoom.CommonRankingRoom, error) {
	t := NewCommonRankingRoom()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_ranking_id = ?", masterRankingId).Where("room_id = ?", roomId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return commonRankingRoom.SetCommonRankingRoom(t.MasterRankingId, t.RoomId, t.UserId, t.Score), nil
}

func (s *commonRankingRoomDao) FindList(ctx context.Context) (commonRankingRoom.CommonRankingRooms, error) {
	ts := NewCommonRankingRooms()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := commonRankingRoom.NewCommonRankingRooms()
	for _, t := range ts {
		ms = append(ms, commonRankingRoom.SetCommonRankingRoom(t.MasterRankingId, t.RoomId, t.UserId, t.Score))
	}

	return ms, nil
}

func (s *commonRankingRoomDao) FindListByMasterRankingIdAndRoomId(ctx context.Context, masterRankingId int64, roomId string) (commonRankingRoom.CommonRankingRooms, error) {
	ts := NewCommonRankingRooms()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_ranking_id = ?", masterRankingId).Where("room_id = ?", roomId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := commonRankingRoom.NewCommonRankingRooms()
	for _, t := range ts {
		ms = append(ms, commonRankingRoom.SetCommonRankingRoom(t.MasterRankingId, t.RoomId, t.UserId, t.Score))
	}

	return ms, nil
}

func (s *commonRankingRoomDao) Create(ctx context.Context, tx *gorm.DB, m *commonRankingRoom.CommonRankingRoom) (*commonRankingRoom.CommonRankingRoom, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &CommonRankingRoom{
		MasterRankingId: m.MasterRankingId,
		RoomId:          m.RoomId,
		UserId:          m.UserId,
		Score:           m.Score,
	}
	res := conn.Model(NewCommonRankingRoom()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return commonRankingRoom.SetCommonRankingRoom(t.MasterRankingId, t.RoomId, t.UserId, t.Score), nil
}

func (s *commonRankingRoomDao) CreateList(ctx context.Context, tx *gorm.DB, ms commonRankingRoom.CommonRankingRooms) (commonRankingRoom.CommonRankingRooms, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewCommonRankingRooms()
	for _, m := range ms {
		t := &CommonRankingRoom{
			MasterRankingId: m.MasterRankingId,
			RoomId:          m.RoomId,
			UserId:          m.UserId,
			Score:           m.Score,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewCommonRankingRoom()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *commonRankingRoomDao) Update(ctx context.Context, tx *gorm.DB, m *commonRankingRoom.CommonRankingRoom) (*commonRankingRoom.CommonRankingRoom, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &CommonRankingRoom{
		MasterRankingId: m.MasterRankingId,
		RoomId:          m.RoomId,
		UserId:          m.UserId,
		Score:           m.Score,
	}
	res := conn.Model(NewCommonRankingRoom()).WithContext(ctx).Where("master_ranking_id = ?", m.MasterRankingId).Where("room_id = ?", m.RoomId).Where("user_id = ?", m.UserId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return commonRankingRoom.SetCommonRankingRoom(t.MasterRankingId, t.RoomId, t.UserId, t.Score), nil
}

func (s *commonRankingRoomDao) Delete(ctx context.Context, tx *gorm.DB, m *commonRankingRoom.CommonRankingRoom) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewCommonRankingRoom()).WithContext(ctx).Where("master_ranking_id = ?", m.MasterRankingId).Where("room_id = ?", m.RoomId).Where("user_id = ?", m.UserId).Delete(NewCommonRankingRoom())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
