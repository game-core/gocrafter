// Package commonRoom ルーム
//
//go:generate mockgen -source=./common_room_mysql_repository.gen.go -destination=./common_room_mysql_repository_mock.gen.go -package=commonRoom
package commonRoom

import (
	"context"

	"gorm.io/gorm"
)

type CommonRoomMysqlRepository interface {
	Find(ctx context.Context, roomId string) (*CommonRoom, error)
	FindOrNil(ctx context.Context, roomId string) (*CommonRoom, error)
	FindByHostUserId(ctx context.Context, hostUserId string) (*CommonRoom, error)
	FindByRoomIdAndHostUserId(ctx context.Context, roomId string, hostUserId string) (*CommonRoom, error)
	FindByName(ctx context.Context, name string) (*CommonRoom, error)
	FindOrNilByHostUserId(ctx context.Context, hostUserId string) (*CommonRoom, error)
	FindOrNilByRoomIdAndHostUserId(ctx context.Context, roomId string, hostUserId string) (*CommonRoom, error)
	FindOrNilByName(ctx context.Context, name string) (*CommonRoom, error)
	FindList(ctx context.Context) (CommonRooms, error)
	FindListByHostUserId(ctx context.Context, hostUserId string) (CommonRooms, error)
	FindListByRoomIdAndHostUserId(ctx context.Context, roomId string, hostUserId string) (CommonRooms, error)
	FindListByName(ctx context.Context, name string) (CommonRooms, error)
	Create(ctx context.Context, tx *gorm.DB, m *CommonRoom) (*CommonRoom, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms CommonRooms) (CommonRooms, error)
	Update(ctx context.Context, tx *gorm.DB, m *CommonRoom) (*CommonRoom, error)
	Delete(ctx context.Context, tx *gorm.DB, m *CommonRoom) error
}
