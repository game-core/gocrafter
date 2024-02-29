//go:generate mockgen -source=./friend_service.go -destination=./friend_service_mock.gen.go -package=friend
package friend

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/friend/userFriend"
)

type FriendService interface {
	Send(ctx context.Context, txs map[string]*gorm.DB, req *FriendSendRequest) (*FriendSendResponse, error)
}

type friendService struct {
	userFriendRepository userFriend.UserFriendRepository
}

func NewFriendService(
	userFriendRepository userFriend.UserFriendRepository,
) FriendService {
	return &friendService{
		userFriendRepository: userFriendRepository,
	}
}

// Send フレンド申請を送信する
func (s *friendService) Send(ctx context.Context, txs map[string]*gorm.DB, req *FriendSendRequest) (*FriendSendResponse, error) {
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
