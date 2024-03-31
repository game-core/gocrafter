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
)

type RankingService interface {
	GetMaster(ctx context.Context, req *RankingGetMasterRequest) (*RankingGetMasterResponse, error)
	Get(ctx context.Context, now time.Time, req *RankingGetRequest) (*RankingGetResponse, error)
	Update(ctx context.Context, tx *gorm.DB, now time.Time, req *RankingUpdateRequest) (*RankingUpdateResponse, error)
}

type rankingService struct {
	commonRankingRoomMysqlRepository  commonRankingRoom.CommonRankingRoomMysqlRepository
	commonRankingWorldMysqlRepository commonRankingWorld.CommonRankingWorldMysqlRepository
	masterRankingMysqlRepository      masterRanking.MasterRankingMysqlRepository
	masterRankingEventMysqlRepository masterRankingEvent.MasterRankingEventMysqlRepository
	masterRankingScopeMysqlRepository masterRankingScope.MasterRankingScopeMysqlRepository
}

func NewRankingService(
	commonRankingRoomMysqlRepository commonRankingRoom.CommonRankingRoomMysqlRepository,
	commonRankingWorldMysqlRepository commonRankingWorld.CommonRankingWorldMysqlRepository,
	masterRankingMysqlRepository masterRanking.MasterRankingMysqlRepository,
	masterRankingEventMysqlRepository masterRankingEvent.MasterRankingEventMysqlRepository,
	masterRankingScopeMysqlRepository masterRankingScope.MasterRankingScopeMysqlRepository,
) RankingService {
	return &rankingService{
		commonRankingRoomMysqlRepository:  commonRankingRoomMysqlRepository,
		commonRankingWorldMysqlRepository: commonRankingWorldMysqlRepository,
		masterRankingMysqlRepository:      masterRankingMysqlRepository,
		masterRankingEventMysqlRepository: masterRankingEventMysqlRepository,
		masterRankingScopeMysqlRepository: masterRankingScopeMysqlRepository,
	}
}

// GetMaster マスターデータを取得する
func (s *rankingService) GetMaster(ctx context.Context, req *RankingGetMasterRequest) (*RankingGetMasterResponse, error) {
	masterRankingModel, err := s.masterRankingMysqlRepository.Find(ctx, req.MasterRankingEventId)
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
	masterRankingModel, err := s.masterRankingMysqlRepository.Find(ctx, req.MasterRankingEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterRankingMysqlRepository.Find", err)
	}

	masterRankingEventModel, err := s.getEvent(ctx, now, masterRankingModel.MasterRankingEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.getEvent", err)
	}

	switch masterRankingModel.RankingScopeType {
	case enum.RankingScopeType_Room:
		result, err := s.getRoomRanking(ctx, masterRankingEventModel.GetLastEventAt(now), masterRankingModel.Id, req.RoomId)
		if err != nil {
			return nil, errors.NewMethodError("s.getRoomRanking", err)
		}
		return SetRankingGetResponse(result, commonRankingWorld.NewCommonRankingWorlds()), nil
	case enum.RankingScopeType_World:
		result, err := s.getWorldRanking(ctx, masterRankingEventModel.GetLastEventAt(now), masterRankingModel.Id)
		if err != nil {
			return nil, errors.NewMethodError("s.getWorldRanking", err)
		}
		return SetRankingGetResponse(commonRankingRoom.NewCommonRankingRooms(), result), nil
	default:
		return nil, nil
	}
}

// Update ランキングを更新する
func (s *rankingService) Update(ctx context.Context, tx *gorm.DB, now time.Time, req *RankingUpdateRequest) (*RankingUpdateResponse, error) {
	masterRankingModel, err := s.masterRankingMysqlRepository.Find(ctx, req.MasterRankingEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterRankingMysqlRepository.Find", err)
	}

	masterRankingEventModel, err := s.getEvent(ctx, now, masterRankingModel.MasterRankingEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.getEvent", err)
	}

	switch masterRankingModel.RankingScopeType {
	case enum.RankingScopeType_Room:
		result, err := s.updateRoomRanking(ctx, tx, now, masterRankingEventModel.GetLastEventAt(now), masterRankingModel.Id, req.RoomId, masterRankingModel.Limit, req.UserId, req.Score)
		if err != nil {
			return nil, errors.NewMethodError("s.updateRoomRanking", err)
		}
		return SetRankingUpdateResponse(result, commonRankingWorld.NewCommonRankingWorlds()), nil
	case enum.RankingScopeType_World:
		result, err := s.updateWorldRanking(ctx, tx, now, masterRankingEventModel.GetLastEventAt(now), masterRankingModel.Id, masterRankingModel.Limit, req.UserId, req.Score)
		if err != nil {
			return nil, errors.NewMethodError("s.updateWorldRanking", err)
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

// getRoomRanking ルームランキングを取得する
func (s *rankingService) getRoomRanking(ctx context.Context, lastEventAt time.Time, masterRankingId int64, roomId string) (commonRankingRoom.CommonRankingRooms, error) {
	commonRankingRoomModels, err := s.commonRankingRoomMysqlRepository.FindListByMasterRankingIdAndRoomId(ctx, masterRankingId, roomId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRankingRoomMysqlRepository.FindListByMasterRankingIdAndRoomId", err)
	}

	return commonRankingRoomModels.SortRanking(lastEventAt), nil
}

// getWorldRanking ワールドランキングを取得する
func (s *rankingService) getWorldRanking(ctx context.Context, lastEventAt time.Time, masterRankingId int64) (commonRankingWorld.CommonRankingWorlds, error) {
	commonRankingWorldModels, err := s.commonRankingWorldMysqlRepository.FindListByMasterRankingId(ctx, masterRankingId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRankingWorldMysqlRepository.FindListByMasterRankingId", err)
	}

	return commonRankingWorldModels.SortRanking(lastEventAt), nil
}

// updateRoomRanking ルームランキングを更新する
func (s *rankingService) updateRoomRanking(ctx context.Context, tx *gorm.DB, now, lastEventAt time.Time, masterRankingId int64, roomId string, limit int32, userId string, score int32) (commonRankingRoom.CommonRankingRooms, error) {
	commonRankingRoomModels, err := s.commonRankingRoomMysqlRepository.FindListByMasterRankingIdAndRoomId(ctx, masterRankingId, roomId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRankingRoomMysqlRepository.FindListByMasterRankingIdAndRoomId", err)
	}

	if len(commonRankingRoomModels) < int(limit) {
		if _, err := s.commonRankingRoomMysqlRepository.Create(ctx, tx, commonRankingRoom.SetCommonRankingRoom(masterRankingId, roomId, userId, score, now)); err != nil {
			return nil, errors.NewMethodError("s.commonRankingRoomMysqlRepository.Create", err)
		}

		result, err := s.getRoomRanking(ctx, lastEventAt, masterRankingId, roomId)
		if err != nil {
			return nil, errors.NewMethodError("s.getRoomRanking", err)
		}

		return result, nil
	}

	if _, err := s.commonRankingRoomMysqlRepository.Update(ctx, tx, commonRankingRoomModels.ExcludeRanking(masterRankingId, roomId, userId, score, now, lastEventAt)); err != nil {
		return nil, errors.NewMethodError("s.commonRankingRoomMysqlRepository.Update", err)
	}

	result, err := s.getRoomRanking(ctx, lastEventAt, masterRankingId, roomId)
	if err != nil {
		return nil, errors.NewMethodError("s.getRoomRanking", err)
	}

	return result, nil
}

// updateWorldRanking ワールドランキングを更新する
func (s *rankingService) updateWorldRanking(ctx context.Context, tx *gorm.DB, now time.Time, lastEventAt time.Time, masterRankingId int64, limit int32, userId string, score int32) (commonRankingWorld.CommonRankingWorlds, error) {
	commonRankingWorldModels, err := s.commonRankingWorldMysqlRepository.FindListByMasterRankingId(ctx, masterRankingId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRankingWorldMysqlRepository.FindListByMasterRankingId", err)
	}

	if len(commonRankingWorldModels) < int(limit) {
		if _, err := s.commonRankingWorldMysqlRepository.Create(ctx, tx, commonRankingWorld.SetCommonRankingWorld(masterRankingId, userId, score, now)); err != nil {
			return nil, errors.NewMethodError("s.commonRankingWorldMysqlRepository.Create", err)
		}

		result, err := s.getWorldRanking(ctx, lastEventAt, masterRankingId)
		if err != nil {
			return nil, errors.NewMethodError("s.getWorldRanking", err)
		}

		return result, nil
	}

	if _, err := s.commonRankingWorldMysqlRepository.Update(ctx, tx, commonRankingWorldModels.ExcludeRanking(masterRankingId, userId, score, now, lastEventAt)); err != nil {
		return nil, errors.NewMethodError("s.commonRankingWorldMysqlRepository.Update", err)
	}

	result, err := s.getWorldRanking(ctx, lastEventAt, masterRankingId)
	if err != nil {
		return nil, errors.NewMethodError("s.getWorldRanking", err)
	}

	return result, nil
}
