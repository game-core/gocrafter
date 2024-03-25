// Package commonRoom ルーム
package commonRoom

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoom"
)

type commonRoomDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
}

func NewCommonRoomDao(conn *database.MysqlHandler) commonRoom.CommonRoomMysqlRepository {
	return &commonRoomDao{
		ReadMysqlConn:  conn.Common.ReadMysqlConn,
		WriteMysqlConn: conn.Common.WriteMysqlConn,
	}
}

func (s *commonRoomDao) Find(ctx context.Context, roomId string) (*commonRoom.CommonRoom, error) {
	t := NewCommonRoom()
	res := s.ReadMysqlConn.WithContext(ctx).Where("room_id = ?", roomId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return commonRoom.SetCommonRoom(t.RoomId, t.HostUserId, t.RoomNumber, t.Name, t.UserCount), nil
}

func (s *commonRoomDao) FindOrNil(ctx context.Context, roomId string) (*commonRoom.CommonRoom, error) {
	t := NewCommonRoom()
	res := s.ReadMysqlConn.WithContext(ctx).Where("room_id = ?", roomId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return commonRoom.SetCommonRoom(t.RoomId, t.HostUserId, t.RoomNumber, t.Name, t.UserCount), nil
}

func (s *commonRoomDao) FindByHostUserId(ctx context.Context, hostUserId string) (*commonRoom.CommonRoom, error) {
	t := NewCommonRoom()
	res := s.ReadMysqlConn.WithContext(ctx).Where("host_user_id = ?", hostUserId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return commonRoom.SetCommonRoom(t.RoomId, t.HostUserId, t.RoomNumber, t.Name, t.UserCount), nil
}

func (s *commonRoomDao) FindOrNilByHostUserId(ctx context.Context, hostUserId string) (*commonRoom.CommonRoom, error) {
	t := NewCommonRoom()
	res := s.ReadMysqlConn.WithContext(ctx).Where("host_user_id = ?", hostUserId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return commonRoom.SetCommonRoom(t.RoomId, t.HostUserId, t.RoomNumber, t.Name, t.UserCount), nil
}

func (s *commonRoomDao) FindList(ctx context.Context) (commonRoom.CommonRooms, error) {
	ts := NewCommonRooms()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := commonRoom.NewCommonRooms()
	for _, t := range ts {
		ms = append(ms, commonRoom.SetCommonRoom(t.RoomId, t.HostUserId, t.RoomNumber, t.Name, t.UserCount))
	}

	return ms, nil
}

func (s *commonRoomDao) FindListByHostUserId(ctx context.Context, hostUserId string) (commonRoom.CommonRooms, error) {
	ts := NewCommonRooms()
	res := s.ReadMysqlConn.WithContext(ctx).Where("host_user_id = ?", hostUserId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := commonRoom.NewCommonRooms()
	for _, t := range ts {
		ms = append(ms, commonRoom.SetCommonRoom(t.RoomId, t.HostUserId, t.RoomNumber, t.Name, t.UserCount))
	}

	return ms, nil
}

func (s *commonRoomDao) Create(ctx context.Context, tx *gorm.DB, m *commonRoom.CommonRoom) (*commonRoom.CommonRoom, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &CommonRoom{
		RoomId:     m.RoomId,
		HostUserId: m.HostUserId,
		RoomNumber: m.RoomNumber,
		Name:       m.Name,
		UserCount:  m.UserCount,
	}
	res := conn.Model(NewCommonRoom()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return commonRoom.SetCommonRoom(t.RoomId, t.HostUserId, t.RoomNumber, t.Name, t.UserCount), nil
}

func (s *commonRoomDao) CreateList(ctx context.Context, tx *gorm.DB, ms commonRoom.CommonRooms) (commonRoom.CommonRooms, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewCommonRooms()
	for _, m := range ms {
		t := &CommonRoom{
			RoomId:     m.RoomId,
			HostUserId: m.HostUserId,
			RoomNumber: m.RoomNumber,
			Name:       m.Name,
			UserCount:  m.UserCount,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewCommonRoom()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *commonRoomDao) Update(ctx context.Context, tx *gorm.DB, m *commonRoom.CommonRoom) (*commonRoom.CommonRoom, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &CommonRoom{
		RoomId:     m.RoomId,
		HostUserId: m.HostUserId,
		RoomNumber: m.RoomNumber,
		Name:       m.Name,
		UserCount:  m.UserCount,
	}
	res := conn.Model(NewCommonRoom()).WithContext(ctx).Where("room_id = ?", m.RoomId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return commonRoom.SetCommonRoom(t.RoomId, t.HostUserId, t.RoomNumber, t.Name, t.UserCount), nil
}

func (s *commonRoomDao) Delete(ctx context.Context, tx *gorm.DB, m *commonRoom.CommonRoom) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewCommonRoom()).WithContext(ctx).Where("room_id = ?", m.RoomId).Delete(NewCommonRoom())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
