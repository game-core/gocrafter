// Package commonRankingRoom ルームランキング
//
//go:generate mockgen -source=./common_ranking_room_mysql_repository.gen.go -destination=./common_ranking_room_mysql_repository_mock.gen.go -package=commonRankingRoom
package commonRankingRoom

import (
	"context"

	"gorm.io/gorm"
)

type CommonRankingRoomMysqlRepository interface {
	Find(ctx context.Context, masterRankingId int64, roomId string, userId string) (*CommonRankingRoom, error)
	FindOrNil(ctx context.Context, masterRankingId int64, roomId string, userId string) (*CommonRankingRoom, error)
	FindByMasterRankingIdAndRoomId(ctx context.Context, masterRankingId int64, roomId string) (*CommonRankingRoom, error)
	FindOrNilByMasterRankingIdAndRoomId(ctx context.Context, masterRankingId int64, roomId string) (*CommonRankingRoom, error)
	FindList(ctx context.Context) (CommonRankingRooms, error)
	FindListByMasterRankingIdAndRoomId(ctx context.Context, masterRankingId int64, roomId string) (CommonRankingRooms, error)
	Create(ctx context.Context, tx *gorm.DB, m *CommonRankingRoom) (*CommonRankingRoom, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms CommonRankingRooms) (CommonRankingRooms, error)
	Update(ctx context.Context, tx *gorm.DB, m *CommonRankingRoom) (*CommonRankingRoom, error)
	Delete(ctx context.Context, tx *gorm.DB, m *CommonRankingRoom) error
}
