//go:generate mockgen -source=./room_service.go -destination=./room_service_mock.gen.go -package=room
package room

import (
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoomUser"
)

type RoomService interface {
}

type roomService struct {
	commonRoomMysqlRepository     commonRoom.CommonRoomMysqlRepository
	commonRoomUserMysqlRepository commonRoomUser.CommonRoomUserMysqlRepository
}

func NewRoomService(
	commonRoomMysqlRepository commonRoom.CommonRoomMysqlRepository,
	commonRoomUserMysqlRepository commonRoomUser.CommonRoomUserMysqlRepository,
) RoomService {
	return &roomService{
		commonRoomMysqlRepository:     commonRoomMysqlRepository,
		commonRoomUserMysqlRepository: commonRoomUserMysqlRepository,
	}
}
