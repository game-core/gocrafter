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
