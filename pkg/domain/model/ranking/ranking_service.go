//go:generate mockgen -source=./ranking_service.go -destination=./ranking_service_mock.gen.go -package=ranking
package ranking

import (
	"context"

	"github.com/game-core/gocrafter/internal/errors"

	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingWorld"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRanking"
)

type RankingService interface {
	Get(ctx context.Context, req *RankingGetRequest) (*RankingGetResponse, error)
}

type rankingService struct {
	commonRankingRoomMysqlRepository  commonRankingRoom.CommonRankingRoomMysqlRepository
	commonRankingWorldMysqlRepository commonRankingWorld.CommonRankingWorldMysqlRepository
	masterRankingMysqlRepository      masterRanking.MasterRankingMysqlRepository
}

func NewRankingService(
	commonRankingRoomMysqlRepository commonRankingRoom.CommonRankingRoomMysqlRepository,
	commonRankingWorldMysqlRepository commonRankingWorld.CommonRankingWorldMysqlRepository,
	masterRankingMysqlRepository masterRanking.MasterRankingMysqlRepository,
) RankingService {
	return &rankingService{
		commonRankingRoomMysqlRepository:  commonRankingRoomMysqlRepository,
		commonRankingWorldMysqlRepository: commonRankingWorldMysqlRepository,
		masterRankingMysqlRepository:      masterRankingMysqlRepository,
	}
}

// Get ランキングを取得する
func (s *rankingService) Get(ctx context.Context, req *RankingGetRequest) (*RankingGetResponse, error) {
	masterRankingModel, err := s.masterRankingMysqlRepository.Find(ctx, req.MasterRankingEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterRankingMysqlRepository.Find", err)
	}

	switch masterRankingModel.RankingScopeType {
	case enum.RankingScopeType_Room:
		result, err := s.getRoomRanking(ctx, masterRankingModel.Id, req.RoomId)
		if err != nil {
			return nil, errors.NewMethodError("s.getRoomRanking", err)
		}
		return SetRankingGetResponse(result, commonRankingWorld.NewCommonRankingWorlds()), nil
	case enum.RankingScopeType_World:
		result, err := s.getWorldRanking(ctx, masterRankingModel.Id)
		if err != nil {
			return nil, errors.NewMethodError("s.getWorldRanking", err)
		}
		return SetRankingGetResponse(commonRankingRoom.NewCommonRankingRooms(), result), nil
	}

	return nil, nil
}

// getRoomRanking ルームランキングを取得する
func (s *rankingService) getRoomRanking(ctx context.Context, masterRankingId int64, roomId string) (commonRankingRoom.CommonRankingRooms, error) {
	commonRankingRoomModels, err := s.commonRankingRoomMysqlRepository.FindListByMasterRankingIdAndRoomId(ctx, masterRankingId, roomId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRankingRoomMysqlRepository.FindListByMasterRankingIdAndRoomId", err)
	}

	return commonRankingRoomModels, nil
}

// getWorldRanking ワールドランキングを取得する
func (s *rankingService) getWorldRanking(ctx context.Context, masterRankingId int64) (commonRankingWorld.CommonRankingWorlds, error) {
	commonRankingWorldModels, err := s.commonRankingWorldMysqlRepository.FindListByMasterRankingId(ctx, masterRankingId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonRankingWorldMysqlRepository.FindListByMasterRankingId", err)
	}

	return commonRankingWorldModels, nil
}
