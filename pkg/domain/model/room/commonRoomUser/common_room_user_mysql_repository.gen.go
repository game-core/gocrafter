// Package commonRoomUser ルームユーザー
//
//go:generate mockgen -source=./common_room_user_mysql_repository.gen.go -destination=./common_room_user_mysql_repository_mock.gen.go -package=commonRoomUser
package commonRoomUser

import (
	"context"

	"gorm.io/gorm"
)

type CommonRoomUserMysqlRepository interface {
	Find(ctx context.Context, roomId string, userId string) (*CommonRoomUser, error)
	FindOrNil(ctx context.Context, roomId string, userId string) (*CommonRoomUser, error)
	FindByRoomId(ctx context.Context, roomId string) (*CommonRoomUser, error)
	FindOrNilByRoomId(ctx context.Context, roomId string) (*CommonRoomUser, error)
	FindList(ctx context.Context) (CommonRoomUsers, error)
	FindListByRoomId(ctx context.Context, roomId string) (CommonRoomUsers, error)
	Create(ctx context.Context, tx *gorm.DB, m *CommonRoomUser) (*CommonRoomUser, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms CommonRoomUsers) (CommonRoomUsers, error)
	Update(ctx context.Context, tx *gorm.DB, m *CommonRoomUser) (*CommonRoomUser, error)
	Delete(ctx context.Context, tx *gorm.DB, m *CommonRoomUser) error
}
