// Package friend フレンド承認レスポンス
package friend

import (
	"github.com/game-core/gocrafter/pkg/domain/model/friend/userFriend"
)

type FriendApproveResponses []*FriendApproveResponse

type FriendApproveResponse struct {
	UserFriend *userFriend.UserFriend
}

func NewFriendApproveResponse() *FriendApproveResponse {
	return &FriendApproveResponse{}
}

func NewFriendApproveResponses() FriendApproveResponses {
	return FriendApproveResponses{}
}

func SetFriendApproveResponse(userFriend *userFriend.UserFriend) *FriendApproveResponse {
	return &FriendApproveResponse{
		UserFriend: userFriend,
	}
}
