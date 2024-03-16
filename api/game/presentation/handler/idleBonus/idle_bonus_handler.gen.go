package idleBonus

import (
	"context"

	"github.com/game-core/gocrafter/api/game/presentation/server/idleBonus"
	idleBonusUsecase "github.com/game-core/gocrafter/api/game/usecase/idleBonus"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/tokens"
)

type IdleBonusHandler interface {
	idleBonus.IdleBonusServer
}

type idleBonusHandler struct {
	idleBonus.UnimplementedIdleBonusServer
	idleBonusUsecase idleBonusUsecase.IdleBonusUsecase
}

func NewIdleBonusHandler(
	idleBonusUsecase idleBonusUsecase.IdleBonusUsecase,
) IdleBonusHandler {
	return &idleBonusHandler{
		idleBonusUsecase: idleBonusUsecase,
	}
}

// GetUser 放置ボーナスユーザーを取得する
func (s *idleBonusHandler) GetUser(ctx context.Context, req *idleBonus.IdleBonusGetUserRequest) (*idleBonus.IdleBonusGetUserResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.idleBonusUsecase.GetUser(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.idleBonusUsecase.GetUser", err)
	}

	return res, nil
}

// GetMaster 放置ボーナスマスターを取得する
func (s *idleBonusHandler) GetMaster(ctx context.Context, req *idleBonus.IdleBonusGetMasterRequest) (*idleBonus.IdleBonusGetMasterResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.idleBonusUsecase.GetMaster(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.idleBonusUsecase.GetMaster", err)
	}

	return res, nil
}

// Receive 放置ボーナスを受け取る
func (s *idleBonusHandler) Receive(ctx context.Context, req *idleBonus.IdleBonusReceiveRequest) (*idleBonus.IdleBonusReceiveResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.idleBonusUsecase.Receive(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.idleBonusUsecase.Receive", err)
	}

	return res, nil
}
