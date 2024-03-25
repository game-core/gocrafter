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
	Search(ctx context.Context, req *roomServer.RoomSearchRequest) (*roomServer.RoomSearchResponse, error)
	Create(ctx context.Context, req *roomServer.RoomCreateRequest) (*roomServer.RoomCreateResponse, error)
	Delete(ctx context.Context, req *roomServer.RoomDeleteRequest) (*roomServer.RoomDeleteResponse, error)
	CheckIn(ctx context.Context, req *roomServer.RoomCheckInRequest) (*roomServer.RoomCheckInResponse, error)
	CheckOut(ctx context.Context, req *roomServer.RoomCheckOutRequest) (*roomServer.RoomCheckOutResponse, error)
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

// Search ルームを検索する
func (s *roomUsecase) Search(ctx context.Context, req *roomServer.RoomSearchRequest) (*roomServer.RoomSearchResponse, error) {
	result, err := s.roomService.Search(ctx, roomService.SetRoomSearchRequest(req.UserId, req.Name))
	if err != nil {
		return nil, errors.NewMethodError("s.roomService.Search", err)
	}

	return roomServer.SetRoomSearchResponse(roomServer.SetCommonRooms(result.CommonRooms)), nil
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

// CheckIn ルームに参加する
func (s *roomUsecase) CheckIn(ctx context.Context, req *roomServer.RoomCheckInRequest) (*roomServer.RoomCheckInResponse, error) {
	// transaction
	tx, err := s.transactionService.CommonMysqlBegin(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.CommonMysqlBegin", err)
	}
	defer func() {
		s.transactionService.CommonMysqlEnd(ctx, tx, err)
	}()

	result, err := s.roomService.CheckIn(ctx, tx, roomService.SetRoomCheckInRequest(req.UserId, req.RoomId))
	if err != nil {
		return nil, errors.NewMethodError("s.roomService.CheckIn", err)
	}

	return roomServer.SetRoomCheckInResponse(
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

// CheckOut ルームを退出する
func (s *roomUsecase) CheckOut(ctx context.Context, req *roomServer.RoomCheckOutRequest) (*roomServer.RoomCheckOutResponse, error) {
	// transaction
	tx, err := s.transactionService.CommonMysqlBegin(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.CommonMysqlBegin", err)
	}
	defer func() {
		s.transactionService.CommonMysqlEnd(ctx, tx, err)
	}()

	result, err := s.roomService.CheckOut(ctx, tx, roomService.SetRoomCheckOutRequest(req.UserId, req.RoomId))
	if err != nil {
		return nil, errors.NewMethodError("s.roomService.CheckOut", err)
	}

	return roomServer.SetRoomCheckOutResponse(
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
