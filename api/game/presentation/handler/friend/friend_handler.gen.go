package friend

import (
	"context"

	"github.com/game-core/gocrafter/api/game/presentation/server/friend"
	friendUsecase "github.com/game-core/gocrafter/api/game/usecase/friend"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/tokens"
)

type FriendHandler interface {
	friend.FriendServer
}

type friendHandler struct {
	friend.UnimplementedFriendServer
	friendUsecase friendUsecase.FriendUsecase
}

func NewFriendHandler(
	friendUsecase friendUsecase.FriendUsecase,
) FriendHandler {
	return &friendHandler{
		friendUsecase: friendUsecase,
	}
}

// Send フレンド申請を送信する
func (s *friendHandler) Send(ctx context.Context, req *friend.FriendSendRequest) (*friend.FriendSendResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.friendUsecase.Send(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.friendUsecase.Send", err)
	}

	return res, nil
}

// Approve フレンド申請を承諾する
func (s *friendHandler) Approve(ctx context.Context, req *friend.FriendApproveRequest) (*friend.FriendApproveResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.friendUsecase.Approve(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.friendUsecase.Approve", err)
	}

	return res, nil
}

// Disapprove フレンド申請を拒否する
func (s *friendHandler) Disapprove(ctx context.Context, req *friend.FriendDisapproveRequest) (*friend.FriendDisapproveResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.friendUsecase.Disapprove(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.friendUsecase.Disapprove", err)
	}

	return res, nil
}

// Delete フレンドを削除する
func (s *friendHandler) Delete(ctx context.Context, req *friend.FriendDeleteRequest) (*friend.FriendDeleteResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.friendUsecase.Delete(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.friendUsecase.Delete", err)
	}

	return res, nil
}
