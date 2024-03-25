// Package commonRoomUser ルームユーザー
package commonRoomUser

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoomUser"
)

type commonRoomUserDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
}

func NewCommonRoomUserDao(conn *database.MysqlHandler) commonRoomUser.CommonRoomUserMysqlRepository {
	return &commonRoomUserDao{
		ReadMysqlConn:  conn.Common.ReadMysqlConn,
		WriteMysqlConn: conn.Common.WriteMysqlConn,
	}
}

func (s *commonRoomUserDao) Find(ctx context.Context, roomId string, userId string) (*commonRoomUser.CommonRoomUser, error) {
	t := NewCommonRoomUser()
	res := s.ReadMysqlConn.WithContext(ctx).Where("room_id = ?", roomId).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return commonRoomUser.SetCommonRoomUser(t.RoomId, t.UserId, t.RoomUserPositionType), nil
}

func (s *commonRoomUserDao) FindOrNil(ctx context.Context, roomId string, userId string) (*commonRoomUser.CommonRoomUser, error) {
	t := NewCommonRoomUser()
	res := s.ReadMysqlConn.WithContext(ctx).Where("room_id = ?", roomId).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return commonRoomUser.SetCommonRoomUser(t.RoomId, t.UserId, t.RoomUserPositionType), nil
}

func (s *commonRoomUserDao) FindByRoomId(ctx context.Context, roomId string) (*commonRoomUser.CommonRoomUser, error) {
	t := NewCommonRoomUser()
	res := s.ReadMysqlConn.WithContext(ctx).Where("room_id = ?", roomId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return commonRoomUser.SetCommonRoomUser(t.RoomId, t.UserId, t.RoomUserPositionType), nil
}

func (s *commonRoomUserDao) FindOrNilByRoomId(ctx context.Context, roomId string) (*commonRoomUser.CommonRoomUser, error) {
	t := NewCommonRoomUser()
	res := s.ReadMysqlConn.WithContext(ctx).Where("room_id = ?", roomId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return commonRoomUser.SetCommonRoomUser(t.RoomId, t.UserId, t.RoomUserPositionType), nil
}

func (s *commonRoomUserDao) FindList(ctx context.Context) (commonRoomUser.CommonRoomUsers, error) {
	ts := NewCommonRoomUsers()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := commonRoomUser.NewCommonRoomUsers()
	for _, t := range ts {
		ms = append(ms, commonRoomUser.SetCommonRoomUser(t.RoomId, t.UserId, t.RoomUserPositionType))
	}

	return ms, nil
}

func (s *commonRoomUserDao) FindListByRoomId(ctx context.Context, roomId string) (commonRoomUser.CommonRoomUsers, error) {
	ts := NewCommonRoomUsers()
	res := s.ReadMysqlConn.WithContext(ctx).Where("room_id = ?", roomId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := commonRoomUser.NewCommonRoomUsers()
	for _, t := range ts {
		ms = append(ms, commonRoomUser.SetCommonRoomUser(t.RoomId, t.UserId, t.RoomUserPositionType))
	}

	return ms, nil
}

func (s *commonRoomUserDao) Create(ctx context.Context, tx *gorm.DB, m *commonRoomUser.CommonRoomUser) (*commonRoomUser.CommonRoomUser, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &CommonRoomUser{
		RoomId:               m.RoomId,
		UserId:               m.UserId,
		RoomUserPositionType: m.RoomUserPositionType,
	}
	res := conn.Model(NewCommonRoomUser()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return commonRoomUser.SetCommonRoomUser(t.RoomId, t.UserId, t.RoomUserPositionType), nil
}

func (s *commonRoomUserDao) CreateList(ctx context.Context, tx *gorm.DB, ms commonRoomUser.CommonRoomUsers) (commonRoomUser.CommonRoomUsers, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewCommonRoomUsers()
	for _, m := range ms {
		t := &CommonRoomUser{
			RoomId:               m.RoomId,
			UserId:               m.UserId,
			RoomUserPositionType: m.RoomUserPositionType,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewCommonRoomUser()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *commonRoomUserDao) Update(ctx context.Context, tx *gorm.DB, m *commonRoomUser.CommonRoomUser) (*commonRoomUser.CommonRoomUser, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &CommonRoomUser{
		RoomId:               m.RoomId,
		UserId:               m.UserId,
		RoomUserPositionType: m.RoomUserPositionType,
	}
	res := conn.Model(NewCommonRoomUser()).WithContext(ctx).Where("room_id = ?", m.RoomId).Where("user_id = ?", m.UserId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return commonRoomUser.SetCommonRoomUser(t.RoomId, t.UserId, t.RoomUserPositionType), nil
}

func (s *commonRoomUserDao) Delete(ctx context.Context, tx *gorm.DB, m *commonRoomUser.CommonRoomUser) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewCommonRoomUser()).WithContext(ctx).Where("room_id = ?", m.RoomId).Where("user_id = ?", m.UserId).Delete(NewCommonRoomUser())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
