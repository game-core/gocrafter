package ranking

import (
	"context"

	"github.com/game-core/gocrafter/api/game/presentation/server/ranking"
	rankingUsecase "github.com/game-core/gocrafter/api/game/usecase/ranking"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/tokens"
)

type RankingHandler interface {
	ranking.RankingServer
}

type rankingHandler struct {
	ranking.UnimplementedRankingServer
	rankingUsecase rankingUsecase.RankingUsecase
}

func NewRankingHandler(
	rankingUsecase rankingUsecase.RankingUsecase,
) RankingHandler {
	return &rankingHandler{
		rankingUsecase: rankingUsecase,
	}
}

// GetMaster マスターデータを取得する
func (s *rankingHandler) GetMaster(ctx context.Context, req *ranking.RankingGetMasterRequest) (*ranking.RankingGetMasterResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.rankingUsecase.GetMaster(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.rankingUsecase.GetMaster", err)
	}

	return res, nil
}

// Get ランキングを取得する
func (s *rankingHandler) Get(ctx context.Context, req *ranking.RankingGetRequest) (*ranking.RankingGetResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.rankingUsecase.Get(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.rankingUsecase.Get", err)
	}

	return res, nil
}

// Update ランキングを更新する
func (s *rankingHandler) Update(ctx context.Context, req *ranking.RankingUpdateRequest) (*ranking.RankingUpdateResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.rankingUsecase.Update(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.rankingUsecase.Update", err)
	}

	return res, nil
}
