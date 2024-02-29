// Package friend フレンド承認拒否レスポンス
package friend

import (
	"github.com/game-core/gocrafter/pkg/domain/model/friend/userFriend"
)

type FriendDisapproveResponses []*FriendDisapproveResponse

type FriendDisapproveResponse struct {
	UserFriend *userFriend.UserFriend
}

func NewFriendDisapproveResponse() *FriendDisapproveResponse {
	return &FriendDisapproveResponse{}
}

func NewFriendDisapproveResponses() FriendDisapproveResponses {
	return FriendDisapproveResponses{}
}

func SetFriendDisapproveResponse(userFriend *userFriend.UserFriend) *FriendDisapproveResponse {
	return &FriendDisapproveResponse{
		UserFriend: userFriend,
	}
}
