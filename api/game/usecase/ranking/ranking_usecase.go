package ranking

import (
	"context"

	rankingServer "github.com/game-core/gocrafter/api/game/presentation/server/ranking"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/times"
	rankingService "github.com/game-core/gocrafter/pkg/domain/model/ranking"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

type RankingUsecase interface {
	GetMaster(ctx context.Context, req *rankingServer.RankingGetMasterRequest) (*rankingServer.RankingGetMasterResponse, error)
	Get(ctx context.Context, req *rankingServer.RankingGetRequest) (*rankingServer.RankingGetResponse, error)
	Update(ctx context.Context, req *rankingServer.RankingUpdateRequest) (*rankingServer.RankingUpdateResponse, error)
}

type rankingUsecase struct {
	rankingService     rankingService.RankingService
	transactionService transactionService.TransactionService
}

func NewRankingUsecase(
	rankingService rankingService.RankingService,
	transactionService transactionService.TransactionService,
) RankingUsecase {
	return &rankingUsecase{
		rankingService:     rankingService,
		transactionService: transactionService,
	}
}

// GetMaster マスターデータを取得する
func (s *rankingUsecase) GetMaster(ctx context.Context, req *rankingServer.RankingGetMasterRequest) (*rankingServer.RankingGetMasterResponse, error) {
	result, err := s.rankingService.GetMaster(ctx, rankingService.SetRankingGetMasterRequest(req.MasterRankingId))
	if err != nil {
		return nil, errors.NewMethodError("s.rankingService.GetMaster", err)
	}

	return rankingServer.SetRankingGetMasterResponse(
		rankingServer.SetMasterRanking(
			result.MasterRanking.Id,
			result.MasterRanking.MasterRankingEventId,
			result.MasterRanking.Name,
			rankingServer.RankingScopeType(result.MasterRanking.RankingScopeType),
			result.MasterRanking.RankingLimit,
		),
		rankingServer.SetMasterRankingEvent(
			result.MasterRankingEvent.Id,
			result.MasterRankingEvent.Name,
			result.MasterRankingEvent.ResetHour,
			result.MasterRankingEvent.IntervalHour,
			result.MasterRankingEvent.RepeatSetting,
			times.TimeToPb(&result.MasterRankingEvent.StartAt),
			times.TimeToPb(result.MasterRankingEvent.EndAt),
		),
		rankingServer.SetMasterRankingScope(
			result.MasterRankingScope.Id,
			result.MasterRankingScope.Name,
			rankingServer.RankingScopeType(result.MasterRanking.RankingScopeType),
		),
	), nil
}

// Get ランキングを取得する
func (s *rankingUsecase) Get(ctx context.Context, req *rankingServer.RankingGetRequest) (*rankingServer.RankingGetResponse, error) {
	result, err := s.rankingService.Get(ctx, times.Now(), rankingService.SetRankingGetRequest(req.UserId, req.MasterRankingId, req.RoomId))
	if err != nil {
		return nil, errors.NewMethodError("s.rankingService.Get", err)
	}

	return rankingServer.SetRankingGetResponse(
		rankingServer.SetCommonRankingRooms(result.CommonRankingRooms),
		rankingServer.SetCommonRankingWorlds(result.CommonRankingWorlds),
	), nil
}

// Update ランキングを更新する
func (s *rankingUsecase) Update(ctx context.Context, req *rankingServer.RankingUpdateRequest) (*rankingServer.RankingUpdateResponse, error) {
	// transaction
	tx, err := s.transactionService.CommonMysqlBegin(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.CommonMysqlBegin", err)
	}
	defer func() {
		s.transactionService.CommonMysqlEnd(ctx, tx, err)
	}()

	result, err := s.rankingService.Update(ctx, tx, times.Now(), rankingService.SetRankingUpdateRequest(req.UserId, req.MasterRankingId, req.RoomId, req.Score))
	if err != nil {
		return nil, errors.NewMethodError("s.rankingService.Update", err)
	}

	return rankingServer.SetRankingUpdateResponse(
		rankingServer.SetCommonRankingRooms(result.CommonRankingRooms),
		rankingServer.SetCommonRankingWorlds(result.CommonRankingWorlds),
	), nil
}
