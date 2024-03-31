//go:generate mockgen -source=./ranking_service.go -destination=./ranking_service_mock.gen.go -package=ranking
package ranking

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingWorld"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRanking"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRankingEvent"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRankingScope"
	roomService "github.com/game-core/gocrafter/pkg/domain/model/room"
)

type RankingService interface {
	GetMaster(ctx context.Context, req *RankingGetMasterRequest) (*RankingGetMasterResponse, error)
	Get(ctx context.Context, now time.Time, req *RankingGetRequest) (*RankingGetResponse, error)
	Update(ctx context.Context, tx *gorm.DB, now time.Time, req *RankingUpdateRequest) (*RankingUpdateResponse, error)
}

type rankingService struct {
	roomService                       roomService.RoomService
	commonRankingRoomMysqlRepository  commonRankingRoom.CommonRankingRoomMysqlRepository
	commonRankingWorldMysqlRepository commonRankingWorld.CommonRankingWorldMysqlRepository
	masterRankingMysqlRepository      masterRanking.MasterRankingMysqlRepository
	masterRankingEventMysqlRepository masterRankingEvent.MasterRankingEventMysqlRepository
	masterRankingScopeMysqlRepository masterRankingScope.MasterRankingScopeMysqlRepository
}

func NewRankingService(
	roomService roomService.RoomService,
	commonRankingRoomMysqlRepository commonRankingRoom.CommonRankingRoomMysqlRepository,
	commonRankingWorldMysqlRepository commonRankingWorld.CommonRankingWorldMysqlRepository,
	masterRankingMysqlRepository masterRanking.MasterRankingMysqlRepository,
	masterRankingEventMysqlRepository masterRankingEvent.MasterRankingEventMysqlRepository,
	masterRankingScopeMysqlRepository masterRankingScope.MasterRankingScopeMysqlRepository,
) RankingService {
	return &rankingService{
		roomService:                       roomService,
		commonRankingRoomMysqlRepository:  commonRankingRoomMysqlRepository,
		commonRankingWorldMysqlRepository: commonRankingWorldMysqlRepository,
		masterRankingMysqlRepository:      masterRankingMysqlRepository,
		masterRankingEventMysqlRepository: masterRankingEventMysqlRepository,
		masterRankingScopeMysqlRepository: masterRankingScopeMysqlRepository,
	}
}

// GetMaster マスターデータを取得する
func (s *rankingService) GetMaster(ctx context.Context, req *RankingGetMasterRequest) (*RankingGetMasterResponse, error) {
	masterRankingModel, err := s.masterRankingMysqlRepository.Find(ctx, req.MasterRankingId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterRankingMysqlRepository.Find", err)
	}

	masterRankingEventModel, err := s.masterRankingEventMysqlRepository.Find(ctx, masterRankingModel.MasterRankingEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterRankingEventMysqlRepository.Find", err)
	}

	masterRankingScopeModel, err := s.masterRankingScopeMysqlRepository.FindByRankingScopeType(ctx, masterRankingModel.RankingScopeType)
	if err != nil {
		return nil, errors.NewMethodError("s.masterRankingScopeMysqlRepository.FindByRankingScopeType", err)
	}

	return SetRankingGetMasterResponse(masterRankingModel, masterRankingEventModel, masterRankingScopeModel), nil
}

// Get ランキングを取得する
func (s *rankingService) Get(ctx context.Context, now time.Time, req *RankingGetRequest) (*RankingGetResponse, error) {
	masterRankingModel, err := s.masterRankingMysqlRepository.Find(ctx, req.MasterRankingId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterRankingMysqlRepository.Find", err)
	}

	masterRankingEventModel, err := s.getEvent(ctx, now, masterRankingModel.MasterRankingEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.getEvent", err)
	}

	switch masterRankingModel.RankingScopeType {
	case enum.RankingScopeType_Room:
		result, err := s.getRoomRankings(ctx, masterRankingEventModel.GetLastEventAt(now), req.UserId, req.RoomId, masterRankingModel.Id)
		if err != nil {
			return nil, errors.NewMethodError("s.getRoomRankings", err)
		}
		return SetRankingGetResponse(result, commonRankingWorld.NewCommonRankingWorlds()), nil
	case enum.RankingScopeType_World:
		result, err := s.getWorldRankings(ctx, masterRankingEventModel.GetLastEventAt(now), masterRankingModel.Id)
		if err != nil {
			return nil, errors.NewMethodError("s.getWorldRankings", err)
		}
		return SetRankingGetResponse(commonRankingRoom.NewCommonRankingRooms(), result), nil
	default:
		return nil, nil
	}
}

// Update ランキングを更新する
func (s *rankingService) Update(ctx context.Context, tx *gorm.DB, now time.Time, req *RankingUpdateRequest) (*RankingUpdateResponse, error) {
	masterRankingModel, err := s.masterRankingMysqlRepository.Find(ctx, req.MasterRankingId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterRankingMysqlRepository.Find", err)
	}

	masterRankingEventModel, err := s.getEvent(ctx, now, masterRankingModel.MasterRankingEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.getEvent", err)
	}

	switch masterRankingModel.RankingScopeType {
	case enum.RankingScopeType_Room:
		result, err := s.updateRoomRankings(ctx, tx, now, masterRankingEventModel.GetLastEventAt(now), req.UserId, req.RoomId, masterRankingModel.Id, masterRankingModel.RankingLimit, req.Score)
		if err != nil {
			return nil, errors.NewMethodError("s.updateRoomRankings", err)
		}
		return SetRankingUpdateResponse(result, commonRankingWorld.NewCommonRankingWorlds()), nil
	case enum.RankingScopeType_World:
		result, err := s.updateWorldRankings(ctx, tx, now, masterRankingEventModel.GetLastEventAt(now), masterRankingModel.Id, masterRankingModel.RankingLimit, req.UserId, req.Score)
		if err != nil {
			return nil, errors.NewMethodError("s.updateWorldRankings", err)
		}
		return SetRankingUpdateResponse(commonRankingRoom.NewCommonRankingRooms(), result), nil
	default:
		return nil, nil
	}
}

// getEvent イベントを取得する
func (s *rankingService) getEvent(ctx context.Context, now time.Time, masterRankingEventId int64) (*masterRankingEvent.MasterRankingEvent, error) {
	masterRankingEventModel, err := s.masterRankingEventMysqlRepository.Find(ctx, masterRankingEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterRankingEventMysqlRepository.Find", err)
	}

	if !masterRankingEventModel.CheckEventPeriod(now) {
		return nil, errors.NewError("outside the event period")
	}

	return masterRankingEventModel, nil
}

// getRoomRankings ルームランキングを取得する
func (s *rankingService) getRoomRankings(ctx context.Context, lastEventAt time.Time, userId, roomId string, masterRankingId int64) (commonRankingRoom.CommonRankingRooms, error) {
	if _, err := s.roomService.Check(ctx, roomService.SetRoomCheckRequest(userId, roomId)); err != nil {
		return nil, errors.NewMethodError("s.roomService.Check", err)
	}

	commonRankingRoomModels, err := s.commonRankingRoomMysqlRepository.FindListByMasterRankingIdAndRoomId(ctx, masterRankingId, roomId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRankingRoomMysqlRepository.FindListByMasterRankingIdAndRoomId", err)
	}

	return commonRankingRoomModels.SortRanking(lastEventAt), nil
}

// getWorldRankings ワールドランキングを取得する
func (s *rankingService) getWorldRankings(ctx context.Context, lastEventAt time.Time, masterRankingId int64) (commonRankingWorld.CommonRankingWorlds, error) {
	commonRankingWorldModels, err := s.commonRankingWorldMysqlRepository.FindListByMasterRankingId(ctx, masterRankingId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRankingWorldMysqlRepository.FindListByMasterRankingId", err)
	}

	return commonRankingWorldModels.SortRanking(lastEventAt), nil
}

// updateRoomRankings ルームランキングを更新する
func (s *rankingService) updateRoomRankings(ctx context.Context, tx *gorm.DB, now, lastEventAt time.Time, userId, roomId string, masterRankingId int64, limit int32, score int32) (commonRankingRoom.CommonRankingRooms, error) {
	if _, err := s.roomService.Check(ctx, roomService.SetRoomCheckRequest(userId, roomId)); err != nil {
		return nil, errors.NewMethodError("s.roomService.Check", err)
	}

	commonRankingRoomModels, err := s.commonRankingRoomMysqlRepository.FindListByMasterRankingIdAndRoomId(ctx, masterRankingId, roomId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRankingRoomMysqlRepository.FindListByMasterRankingIdAndRoomId", err)
	}

	ranking, updateModel, deleteModel := commonRankingRoomModels.AggregateRanking(masterRankingId, roomId, userId, score, now, lastEventAt, limit)
	if deleteModel != nil {
		if err := s.commonRankingRoomMysqlRepository.Delete(ctx, tx, deleteModel); err != nil {
			return nil, errors.NewMethodError("s.commonRankingRoomMysqlRepository.Delete", err)
		}
	}

	if updateModel != nil {
		if err := s.updateRoomRanking(ctx, tx, commonRankingRoomModels, updateModel); err != nil {
			return nil, errors.NewMethodError("s.updateRoomRanking", err)
		}
	}

	return ranking, nil
}

// updateRoomRanking ルームランキングを更新する
func (s *rankingService) updateRoomRanking(ctx context.Context, tx *gorm.DB, commonRankingRoomModels commonRankingRoom.CommonRankingRooms, commonRankingRoomModel *commonRankingRoom.CommonRankingRoom) error {
	if commonRankingRoomModels.CheckRankingByUserId(commonRankingRoomModel.UserId) {
		if _, err := s.commonRankingRoomMysqlRepository.Update(ctx, tx, commonRankingRoomModel); err != nil {
			return errors.NewMethodError("s.commonRankingRoomMysqlRepository.Update", err)
		}

		return nil
	}

	if _, err := s.commonRankingRoomMysqlRepository.Create(ctx, tx, commonRankingRoomModel); err != nil {
		return errors.NewMethodError("s.commonRankingRoomMysqlRepository.Create", err)
	}

	return nil
}

// updateWorldRankings ルームランキングを更新する
func (s *rankingService) updateWorldRankings(ctx context.Context, tx *gorm.DB, now, lastEventAt time.Time, masterRankingId int64, limit int32, userId string, score int32) (commonRankingWorld.CommonRankingWorlds, error) {
	commonRankingWorldModels, err := s.commonRankingWorldMysqlRepository.FindListByMasterRankingId(ctx, masterRankingId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRankingWorldMysqlRepository.FindListByMasterRankingId", err)
	}

	ranking, updateModel, deleteModel := commonRankingWorldModels.AggregateRanking(masterRankingId, userId, score, now, lastEventAt, limit)
	if deleteModel != nil {
		if err := s.commonRankingWorldMysqlRepository.Delete(ctx, tx, deleteModel); err != nil {
			return nil, errors.NewMethodError("s.commonRankingWorldMysqlRepository.Delete", err)
		}
	}

	if updateModel != nil {
		if err := s.updateWorldRanking(ctx, tx, commonRankingWorldModels, updateModel); err != nil {
			return nil, errors.NewMethodError("s.updateWorldRanking", err)
		}
	}

	return ranking, nil
}

// updateWorldRanking ルームランキングを更新する
func (s *rankingService) updateWorldRanking(ctx context.Context, tx *gorm.DB, commonRankingWorldModels commonRankingWorld.CommonRankingWorlds, commonRankingWorldModel *commonRankingWorld.CommonRankingWorld) error {
	if commonRankingWorldModels.CheckRankingByUserId(commonRankingWorldModel.UserId) {
		if _, err := s.commonRankingWorldMysqlRepository.Update(ctx, tx, commonRankingWorldModel); err != nil {
			return errors.NewMethodError("s.commonRankingWorldMysqlRepository.Update", err)
		}

		return nil
	}

	if _, err := s.commonRankingWorldMysqlRepository.Create(ctx, tx, commonRankingWorldModel); err != nil {
		return errors.NewMethodError("s.commonRankingWorldMysqlRepository.Create", err)
	}

	return nil
}
