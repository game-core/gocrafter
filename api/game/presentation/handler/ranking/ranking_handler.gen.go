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
