// Package friend フレンド削除レスポンス
package friend

import (
	"github.com/game-core/gocrafter/pkg/domain/model/friend/userFriend"
)

type FriendDeleteResponses []*FriendDeleteResponse

type FriendDeleteResponse struct {
	UserFriend *userFriend.UserFriend
}

func NewFriendDeleteResponse() *FriendDeleteResponse {
	return &FriendDeleteResponse{}
}

func NewFriendDeleteResponses() FriendDeleteResponses {
	return FriendDeleteResponses{}
}

func SetFriendDeleteResponse(userFriend *userFriend.UserFriend) *FriendDeleteResponse {
	return &FriendDeleteResponse{
		UserFriend: userFriend,
	}
}
