package room

import (
	"context"

	"github.com/game-core/gocrafter/api/game/presentation/server/room"
	roomUsecase "github.com/game-core/gocrafter/api/game/usecase/room"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/tokens"
)

type RoomHandler interface {
	room.RoomServer
}

type roomHandler struct {
	room.UnimplementedRoomServer
	roomUsecase roomUsecase.RoomUsecase
}

func NewRoomHandler(
	roomUsecase roomUsecase.RoomUsecase,
) RoomHandler {
	return &roomHandler{
		roomUsecase: roomUsecase,
	}
}

// Search ルームを検索する
func (s *roomHandler) Search(ctx context.Context, req *room.RoomSearchRequest) (*room.RoomSearchResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.roomUsecase.Search(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.roomUsecase.Search", err)
	}

	return res, nil
}

// Create ルームを作成する
func (s *roomHandler) Create(ctx context.Context, req *room.RoomCreateRequest) (*room.RoomCreateResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.roomUsecase.Create(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.roomUsecase.Create", err)
	}

	return res, nil
}

// Delete ルームを削除する
func (s *roomHandler) Delete(ctx context.Context, req *room.RoomDeleteRequest) (*room.RoomDeleteResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.roomUsecase.Delete(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.roomUsecase.Delete", err)
	}

	return res, nil
}

// Check ルームを確認する
func (s *roomHandler) Check(ctx context.Context, req *room.RoomCheckRequest) (*room.RoomCheckResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.roomUsecase.Check(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.roomUsecase.Check", err)
	}

	return res, nil
}

// CheckIn ルームに参加する
func (s *roomHandler) CheckIn(ctx context.Context, req *room.RoomCheckInRequest) (*room.RoomCheckInResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.roomUsecase.CheckIn(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.roomUsecase.CheckIn", err)
	}

	return res, nil
}

// CheckOut ルームを退出する
func (s *roomHandler) CheckOut(ctx context.Context, req *room.RoomCheckOutRequest) (*room.RoomCheckOutResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.roomUsecase.CheckOut(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.roomUsecase.CheckOut", err)
	}

	return res, nil
}
