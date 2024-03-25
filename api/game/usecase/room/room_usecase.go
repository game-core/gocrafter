package room

import (
	"context"

	"github.com/game-core/gocrafter/pkg/domain/enum"

	roomServer "github.com/game-core/gocrafter/api/game/presentation/server/room"
	"github.com/game-core/gocrafter/internal/errors"
	roomService "github.com/game-core/gocrafter/pkg/domain/model/room"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

type RoomUsecase interface {
	Create(ctx context.Context, req *roomServer.RoomCreateRequest) (*roomServer.RoomCreateResponse, error)
	Delete(ctx context.Context, req *roomServer.RoomDeleteRequest) (*roomServer.RoomDeleteResponse, error)
	Join(ctx context.Context, req *roomServer.RoomJoinRequest) (*roomServer.RoomJoinResponse, error)
}

type roomUsecase struct {
	roomService        roomService.RoomService
	transactionService transactionService.TransactionService
}

func NewRoomUsecase(
	roomService roomService.RoomService,
	transactionService transactionService.TransactionService,
) RoomUsecase {
	return &roomUsecase{
		roomService:        roomService,
		transactionService: transactionService,
	}
}

// Create ルームを作成する
func (s *roomUsecase) Create(ctx context.Context, req *roomServer.RoomCreateRequest) (*roomServer.RoomCreateResponse, error) {
	// transaction
	tx, err := s.transactionService.CommonMysqlBegin(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.CommonMysqlBegin", err)
	}
	defer func() {
		s.transactionService.CommonMysqlEnd(ctx, tx, err)
	}()

	result, err := s.roomService.Create(ctx, tx, roomService.SetRoomCreateRequest(req.UserId, req.Name, enum.RoomReleaseType(req.RoomReleaseType)))
	if err != nil {
		return nil, errors.NewMethodError("s.roomService.Create", err)
	}

	return roomServer.SetRoomCreateResponse(
		roomServer.SetCommonRoom(
			result.CommonRoom.RoomId,
			result.CommonRoom.HostUserId,
			roomServer.RoomReleaseType(result.CommonRoom.RoomReleaseType),
			result.CommonRoom.Name,
			result.CommonRoom.UserCount,
		),
	), nil
}

// Delete ルームを削除する
func (s *roomUsecase) Delete(ctx context.Context, req *roomServer.RoomDeleteRequest) (*roomServer.RoomDeleteResponse, error) {
	// transaction
	tx, err := s.transactionService.CommonMysqlBegin(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.CommonMysqlBegin", err)
	}
	defer func() {
		s.transactionService.CommonMysqlEnd(ctx, tx, err)
	}()

	result, err := s.roomService.Delete(ctx, tx, roomService.SetRoomDeleteRequest(req.UserId, req.RoomId))
	if err != nil {
		return nil, errors.NewMethodError("s.roomService.Delete", err)
	}

	return roomServer.SetRoomDeleteResponse(
		roomServer.SetCommonRoom(
			result.CommonRoom.RoomId,
			result.CommonRoom.HostUserId,
			roomServer.RoomReleaseType(result.CommonRoom.RoomReleaseType),
			result.CommonRoom.Name,
			result.CommonRoom.UserCount,
		),
	), nil
}

// Join ルームに参加する
func (s *roomUsecase) Join(ctx context.Context, req *roomServer.RoomJoinRequest) (*roomServer.RoomJoinResponse, error) {
	// transaction
	tx, err := s.transactionService.CommonMysqlBegin(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.CommonMysqlBegin", err)
	}
	defer func() {
		s.transactionService.CommonMysqlEnd(ctx, tx, err)
	}()

	result, err := s.roomService.Join(ctx, tx, roomService.SetRoomJoinRequest(req.UserId, req.RoomId))
	if err != nil {
		return nil, errors.NewMethodError("s.roomService.Join", err)
	}

	return roomServer.SetRoomJoinResponse(
		roomServer.SetCommonRoom(
			result.CommonRoom.RoomId,
			result.CommonRoom.HostUserId,
			roomServer.RoomReleaseType(result.CommonRoom.RoomReleaseType),
			result.CommonRoom.Name,
			result.CommonRoom.UserCount,
		),
		roomServer.SetCommonRoomUser(
			result.CommonRoomUser.UserId,
			result.CommonRoomUser.RoomId,
			roomServer.RoomUserPositionType(result.CommonRoomUser.RoomUserPositionType),
		),
	), nil
}
