//go:generate mockgen -source=./friend_service.go -destination=./friend_service_mock.gen.go -package=friend
package friend

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/account"
	"github.com/game-core/gocrafter/pkg/domain/model/friend/userFriend"
)

type FriendService interface {
	Send(ctx context.Context, txs map[string]*gorm.DB, req *FriendSendRequest) (*FriendSendResponse, error)
	Approve(ctx context.Context, txs map[string]*gorm.DB, req *FriendApproveRequest) (*FriendApproveResponse, error)
	Disapprove(ctx context.Context, txs map[string]*gorm.DB, req *FriendDisapproveRequest) (*FriendDisapproveResponse, error)
	Delete(ctx context.Context, txs map[string]*gorm.DB, req *FriendDeleteRequest) (*FriendDeleteResponse, error)
}

type friendService struct {
	accountService       account.AccountService
	userFriendRepository userFriend.UserFriendRepository
}

func NewFriendService(
	accountService account.AccountService,
	userFriendRepository userFriend.UserFriendRepository,
) FriendService {
	return &friendService{
		accountService:       accountService,
		userFriendRepository: userFriendRepository,
	}
}

// Send フレンド申請を送信する
func (s *friendService) Send(ctx context.Context, txs map[string]*gorm.DB, req *FriendSendRequest) (*FriendSendResponse, error) {
	if _, err := s.accountService.FindByUserId(ctx, req.FriendUserId); err != nil {
		return nil, errors.NewMethodError("s.accountService.FindByUserId", err)
	}

	userFriendModel, err := s.userFriendRepository.FindOrNil(ctx, req.UserId, req.FriendUserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userFriendRepository.FindOrNil", err)
	}

	if userFriendModel != nil {
		return nil, errors.NewError("already applied")
	}

	if _, err := s.userFriendRepository.Create(ctx, txs[req.FriendUserId], userFriend.SetUserFriend(req.FriendUserId, req.UserId, enum.FriendType_NotApproved)); err != nil {
		return nil, errors.NewMethodError("s.userFriendRepository.Create", err)
	}

	result, err := s.userFriendRepository.Create(ctx, txs[req.UserId], userFriend.SetUserFriend(req.UserId, req.FriendUserId, enum.FriendType_Applying))
	if err != nil {
		return nil, errors.NewMethodError("s.userFriendRepository.Create", err)
	}

	return SetFriendSendResponse(result), nil
}

// Approve フレンド申請を承認する
func (s *friendService) Approve(ctx context.Context, txs map[string]*gorm.DB, req *FriendApproveRequest) (*FriendApproveResponse, error) {
	userFriendModel, err := s.userFriendRepository.FindOrNil(ctx, req.UserId, req.FriendUserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userFriendRepository.FindOrNil", err)
	}

	if userFriendModel == nil || userFriendModel.FriendType != enum.FriendType_NotApproved {
		return nil, errors.NewError("not applied")
	}

	if _, err := s.userFriendRepository.Update(ctx, txs[req.FriendUserId], userFriend.SetUserFriend(req.FriendUserId, req.UserId, enum.FriendType_Approved)); err != nil {
		return nil, errors.NewMethodError("s.userFriendRepository.Update", err)
	}

	result, err := s.userFriendRepository.Update(ctx, txs[req.UserId], userFriend.SetUserFriend(req.UserId, req.FriendUserId, enum.FriendType_Approved))
	if err != nil {
		return nil, errors.NewMethodError("s.userFriendRepository.Update", err)
	}

	return SetFriendApproveResponse(result), nil
}

// Disapprove フレンド申請を拒否する
func (s *friendService) Disapprove(ctx context.Context, txs map[string]*gorm.DB, req *FriendDisapproveRequest) (*FriendDisapproveResponse, error) {
	userFriendModel, err := s.userFriendRepository.FindOrNil(ctx, req.UserId, req.FriendUserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userFriendRepository.FindOrNil", err)
	}

	if userFriendModel == nil || userFriendModel.FriendType != enum.FriendType_NotApproved {
		return nil, errors.NewError("not applied")
	}

	if err := s.userFriendRepository.Delete(ctx, txs[req.FriendUserId], userFriend.SetUserFriend(req.FriendUserId, req.UserId, enum.FriendType_NotApproved)); err != nil {
		return nil, errors.NewMethodError("s.userFriendRepository.Delete", err)
	}

	if err := s.userFriendRepository.Delete(ctx, txs[req.UserId], userFriend.SetUserFriend(req.UserId, req.FriendUserId, enum.FriendType_Applying)); err != nil {
		return nil, errors.NewMethodError("s.userFriendRepository.Delete", err)
	}

	return SetFriendDisapproveResponse(userFriend.SetUserFriend(req.UserId, req.FriendUserId, enum.FriendType_Disapproved)), nil
}

// Delete フレンドを削除する
func (s *friendService) Delete(ctx context.Context, txs map[string]*gorm.DB, req *FriendDeleteRequest) (*FriendDeleteResponse, error) {
	userFriendModel, err := s.userFriendRepository.FindOrNil(ctx, req.UserId, req.FriendUserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userFriendRepository.FindOrNil", err)
	}

	if userFriendModel == nil || userFriendModel.FriendType != enum.FriendType_Approved {
		return nil, errors.NewError("not friend")
	}

	if err := s.userFriendRepository.Delete(ctx, txs[req.FriendUserId], userFriend.SetUserFriend(req.FriendUserId, req.UserId, enum.FriendType_Approved)); err != nil {
		return nil, errors.NewMethodError("s.userFriendRepository.Delete", err)
	}

	if err := s.userFriendRepository.Delete(ctx, txs[req.UserId], userFriend.SetUserFriend(req.UserId, req.FriendUserId, enum.FriendType_Approved)); err != nil {
		return nil, errors.NewMethodError("s.userFriendRepository.Delete", err)
	}

	return SetFriendDeleteResponse(userFriend.SetUserFriend(req.UserId, req.FriendUserId, enum.FriendType_NotFriend)), nil
}
