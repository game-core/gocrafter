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
	Get(ctx context.Context, req *FriendGetRequest) (*FriendGetResponse, error)
	Send(ctx context.Context, txs map[string]*gorm.DB, req *FriendSendRequest) (*FriendSendResponse, error)
	Approve(ctx context.Context, txs map[string]*gorm.DB, req *FriendApproveRequest) (*FriendApproveResponse, error)
	Disapprove(ctx context.Context, txs map[string]*gorm.DB, req *FriendDisapproveRequest) (*FriendDisapproveResponse, error)
	Delete(ctx context.Context, txs map[string]*gorm.DB, req *FriendDeleteRequest) (*FriendDeleteResponse, error)
	Check(ctx context.Context, req *FriendCheckRequest) (*FriendCheckResponse, error)
}

type friendService struct {
	accountService            account.AccountService
	userFriendMysqlRepository userFriend.UserFriendMysqlRepository
}

func NewFriendService(
	accountService account.AccountService,
	userFriendMysqlRepository userFriend.UserFriendMysqlRepository,
) FriendService {
	return &friendService{
		accountService:            accountService,
		userFriendMysqlRepository: userFriendMysqlRepository,
	}
}

// Get フレンドを取得する
func (s *friendService) Get(ctx context.Context, req *FriendGetRequest) (*FriendGetResponse, error) {
	userFriendModels, err := s.userFriendMysqlRepository.FindList(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userFriendMysqlRepository.FindList", err)
	}

	return SetFriendGetResponse(userFriendModels), nil
}

// Send フレンド申請を送信する
func (s *friendService) Send(ctx context.Context, txs map[string]*gorm.DB, req *FriendSendRequest) (*FriendSendResponse, error) {
	if _, err := s.accountService.Check(ctx, account.SetAccountCheckRequest(req.FriendUserId)); err != nil {
		return nil, errors.NewMethodError("s.accountService.Check", err)
	}

	userFriendModel, err := s.userFriendMysqlRepository.FindOrNil(ctx, req.UserId, req.FriendUserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userFriendMysqlRepository.FindOrNil", err)
	}

	if userFriendModel != nil {
		return nil, errors.NewError("already applied")
	}

	if _, err := s.userFriendMysqlRepository.Create(ctx, txs[req.FriendUserId], userFriend.SetUserFriend(req.FriendUserId, req.UserId, enum.FriendType_NotApproved)); err != nil {
		return nil, errors.NewMethodError("s.userFriendMysqlRepository.Create", err)
	}

	result, err := s.userFriendMysqlRepository.Create(ctx, txs[req.UserId], userFriend.SetUserFriend(req.UserId, req.FriendUserId, enum.FriendType_Applying))
	if err != nil {
		return nil, errors.NewMethodError("s.userFriendMysqlRepository.Create", err)
	}

	return SetFriendSendResponse(result), nil
}

// Approve フレンド申請を承認する
func (s *friendService) Approve(ctx context.Context, txs map[string]*gorm.DB, req *FriendApproveRequest) (*FriendApproveResponse, error) {
	userFriendModel, err := s.userFriendMysqlRepository.FindOrNil(ctx, req.UserId, req.FriendUserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userFriendMysqlRepository.FindOrNil", err)
	}

	if userFriendModel == nil || userFriendModel.FriendType != enum.FriendType_NotApproved {
		return nil, errors.NewError("not applied")
	}

	if _, err := s.userFriendMysqlRepository.Update(ctx, txs[req.FriendUserId], userFriend.SetUserFriend(req.FriendUserId, req.UserId, enum.FriendType_Approved)); err != nil {
		return nil, errors.NewMethodError("s.userFriendMysqlRepository.Update", err)
	}

	result, err := s.userFriendMysqlRepository.Update(ctx, txs[req.UserId], userFriend.SetUserFriend(req.UserId, req.FriendUserId, enum.FriendType_Approved))
	if err != nil {
		return nil, errors.NewMethodError("s.userFriendMysqlRepository.Update", err)
	}

	return SetFriendApproveResponse(result), nil
}

// Disapprove フレンド申請を拒否する
func (s *friendService) Disapprove(ctx context.Context, txs map[string]*gorm.DB, req *FriendDisapproveRequest) (*FriendDisapproveResponse, error) {
	userFriendModel, err := s.userFriendMysqlRepository.FindOrNil(ctx, req.UserId, req.FriendUserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userFriendMysqlRepository.FindOrNil", err)
	}

	if userFriendModel == nil || userFriendModel.FriendType != enum.FriendType_NotApproved {
		return nil, errors.NewError("not applied")
	}

	if err := s.userFriendMysqlRepository.Delete(ctx, txs[req.FriendUserId], userFriend.SetUserFriend(req.FriendUserId, req.UserId, enum.FriendType_NotApproved)); err != nil {
		return nil, errors.NewMethodError("s.userFriendMysqlRepository.Delete", err)
	}

	if err := s.userFriendMysqlRepository.Delete(ctx, txs[req.UserId], userFriend.SetUserFriend(req.UserId, req.FriendUserId, enum.FriendType_NotApproved)); err != nil {
		return nil, errors.NewMethodError("s.userFriendMysqlRepository.Delete", err)
	}

	return SetFriendDisapproveResponse(userFriend.SetUserFriend(req.UserId, req.FriendUserId, enum.FriendType_Disapproved)), nil
}

// Delete フレンドを削除する
func (s *friendService) Delete(ctx context.Context, txs map[string]*gorm.DB, req *FriendDeleteRequest) (*FriendDeleteResponse, error) {
	userFriendModel, err := s.userFriendMysqlRepository.FindOrNil(ctx, req.UserId, req.FriendUserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userFriendMysqlRepository.FindOrNil", err)
	}

	if userFriendModel == nil || userFriendModel.FriendType != enum.FriendType_Approved {
		return nil, errors.NewError("not friend")
	}

	if err := s.userFriendMysqlRepository.Delete(ctx, txs[req.FriendUserId], userFriend.SetUserFriend(req.FriendUserId, req.UserId, enum.FriendType_Approved)); err != nil {
		return nil, errors.NewMethodError("s.userFriendMysqlRepository.Delete", err)
	}

	if err := s.userFriendMysqlRepository.Delete(ctx, txs[req.UserId], userFriend.SetUserFriend(req.UserId, req.FriendUserId, enum.FriendType_Approved)); err != nil {
		return nil, errors.NewMethodError("s.userFriendMysqlRepository.Delete", err)
	}

	return SetFriendDeleteResponse(userFriend.SetUserFriend(req.UserId, req.FriendUserId, enum.FriendType_NotFriend)), nil
}

// Check フレンドを確認する
func (s *friendService) Check(ctx context.Context, req *FriendCheckRequest) (*FriendCheckResponse, error) {
	userFriendModel, err := s.userFriendMysqlRepository.Find(ctx, req.UserId, req.FriendUserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userFriendMysqlRepository.Find", err)
	}

	if userFriendModel.FriendType != enum.FriendType_Approved {
		return nil, errors.NewError("not a friend")
	}

	return SetFriendCheckResponse(userFriendModel), nil
}
