//go:generate mockgen -source=./room_service.go -destination=./room_service_mock.gen.go -package=room
package room

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/keys"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	configService "github.com/game-core/gocrafter/pkg/domain/model/config"
	friendService "github.com/game-core/gocrafter/pkg/domain/model/friend"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoomUser"
)

type RoomService interface {
	Search(ctx context.Context, req *RoomSearchRequest) (*RoomSearchResponse, error)
	Create(ctx context.Context, tx *gorm.DB, req *RoomCreateRequest) (*RoomCreateResponse, error)
	Delete(ctx context.Context, tx *gorm.DB, req *RoomDeleteRequest) (*RoomDeleteResponse, error)
	CheckIn(ctx context.Context, tx *gorm.DB, req *RoomCheckInRequest) (*RoomCheckInResponse, error)
	CheckOut(ctx context.Context, tx *gorm.DB, req *RoomCheckOutRequest) (*RoomCheckOutResponse, error)
}

type roomService struct {
	configService                 configService.ConfigService
	friendService                 friendService.FriendService
	commonRoomMysqlRepository     commonRoom.CommonRoomMysqlRepository
	commonRoomUserMysqlRepository commonRoomUser.CommonRoomUserMysqlRepository
}

func NewRoomService(
	configService configService.ConfigService,
	friendService friendService.FriendService,
	commonRoomMysqlRepository commonRoom.CommonRoomMysqlRepository,
	commonRoomUserMysqlRepository commonRoomUser.CommonRoomUserMysqlRepository,
) RoomService {
	return &roomService{
		configService:                 configService,
		friendService:                 friendService,
		commonRoomMysqlRepository:     commonRoomMysqlRepository,
		commonRoomUserMysqlRepository: commonRoomUserMysqlRepository,
	}
}

// Search ルームを検索する
func (s *roomService) Search(ctx context.Context, req *RoomSearchRequest) (*RoomSearchResponse, error) {
	commonRoomModels, err := s.commonRoomMysqlRepository.FindListByName(ctx, req.Name)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRoomMysqlRepository.FindListByName", err)
	}

	return SetRoomSearchResponse(commonRoomModels), nil
}

// Create ルームを作成する
func (s *roomService) Create(ctx context.Context, tx *gorm.DB, req *RoomCreateRequest) (*RoomCreateResponse, error) {
	roomId, err := s.generateRoomId(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.generateRoomId", err)
	}

	commonRoomModel, err := s.commonRoomMysqlRepository.Create(ctx, tx, commonRoom.SetCommonRoom(roomId, req.UserId, req.RoomReleaseType, req.Name, 1))
	if err != nil {
		return nil, errors.NewMethodError("s.commonRoomMysqlRepository.Create", err)
	}

	if _, err := s.commonRoomUserMysqlRepository.Create(ctx, tx, commonRoomUser.SetCommonRoomUser(roomId, req.UserId, enum.RoomUserPositionType_Leader)); err != nil {
		return nil, errors.NewMethodError("s.commonRoomUserMysqlRepository.Create", err)
	}

	return SetRoomCreateResponse(commonRoomModel), nil
}

// Delete ルームを削除する
func (s *roomService) Delete(ctx context.Context, tx *gorm.DB, req *RoomDeleteRequest) (*RoomDeleteResponse, error) {
	commonRoomModel, err := s.commonRoomMysqlRepository.FindByRoomIdAndHostUserId(ctx, req.RoomId, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRoomMysqlRepository.FindByRoomIdAndHostUserId", err)
	}

	if err := s.commonRoomMysqlRepository.Delete(ctx, tx, commonRoomModel); err != nil {
		return nil, errors.NewMethodError("s.commonRoomMysqlRepository.Delete", err)
	}

	commonRoomUserModels, err := s.commonRoomUserMysqlRepository.FindListByRoomId(ctx, req.RoomId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRoomUserMysqlRepository.FindListByRoomId", err)
	}

	for _, commonRoomUserModel := range commonRoomUserModels {
		if err := s.commonRoomUserMysqlRepository.Delete(ctx, tx, commonRoomUserModel); err != nil {
			return nil, errors.NewMethodError("s.commonRoomUserMysqlRepository.Delete", err)
		}
	}

	return SetRoomDeleteResponse(commonRoomModel), nil
}

// CheckIn ルームに参加する
func (s *roomService) CheckIn(ctx context.Context, tx *gorm.DB, req *RoomCheckInRequest) (*RoomCheckInResponse, error) {
	commonRoomModel, err := s.commonRoomMysqlRepository.Find(ctx, req.RoomId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRoomMysqlRepository.Find", err)
	}

	switch commonRoomModel.RoomReleaseType {
	case enum.RoomReleaseType_Private:
		if _, err := s.friendService.Check(ctx, friendService.SetFriendCheckRequest(commonRoomModel.HostUserId, req.UserId)); err != nil {
			return nil, errors.NewMethodError("s.friendService.Check", err)
		}
	case enum.RoomReleaseType_Public:
	}

	commonRoomModel.UserCount += 1
	newCommonRoomModel, err := s.commonRoomMysqlRepository.Update(ctx, tx, commonRoomModel)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRoomMysqlRepository.Update", err)
	}

	commonRoomUserModel, err := s.commonRoomUserMysqlRepository.Create(ctx, tx, commonRoomUser.SetCommonRoomUser(req.RoomId, req.UserId, enum.RoomUserPositionType_General))
	if err != nil {
		return nil, errors.NewMethodError("s.commonRoomMysqlRepository.Find", err)
	}

	return SetRoomCheckInResponse(newCommonRoomModel, commonRoomUserModel), nil
}

// CheckOut ルームを退出する
func (s *roomService) CheckOut(ctx context.Context, tx *gorm.DB, req *RoomCheckOutRequest) (*RoomCheckOutResponse, error) {
	commonRoomModel, err := s.commonRoomMysqlRepository.Find(ctx, req.RoomId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRoomMysqlRepository.Find", err)
	}

	if commonRoomModel.HostUserId == req.UserId {
		return nil, errors.NewError("host user cannot check out")
	}

	commonRoomUserModel, err := s.commonRoomUserMysqlRepository.Find(ctx, req.RoomId, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRoomUserMysqlRepository.Find", err)
	}

	commonRoomModel.UserCount -= 1
	newCommonRoomModel, err := s.commonRoomMysqlRepository.Update(ctx, tx, commonRoomModel)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRoomMysqlRepository.Update", err)
	}

	if err := s.commonRoomUserMysqlRepository.Delete(ctx, tx, commonRoomUserModel); err != nil {
		return nil, errors.NewMethodError("s.commonRoomUserMysqlRepository.Delete", err)
	}

	return SetRoomCheckOutResponse(newCommonRoomModel, commonRoomUserModel), nil
}

// generateRoomId ルームIdを生成する
func (s *roomService) generateRoomId(ctx context.Context, userId string) (string, error) {
	commonRoomModels, err := s.commonRoomMysqlRepository.FindListByHostUserId(ctx, userId)
	if err != nil {
		return "", errors.NewMethodError("s.commonRoomMysqlRepository.Create", err)
	}

	maxRoomNumber, err := s.configService.GetInt32(ctx, enum.ConfigType_MaxRoomNumber)
	if err != nil {
		return "", errors.NewMethodError("s.configService.GetInt32", err)
	}

	if len(commonRoomModels) > int(maxRoomNumber) {
		return "", errors.NewError("room number exceeded")
	}

	roomId, err := keys.GenerateRoomId()
	if err != nil {
		return "", errors.NewMethodError("keys.GenerateRoomId", err)
	}

	return roomId, nil
}
