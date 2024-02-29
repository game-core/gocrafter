// Package friend フレンド送信レスポンス
package friend

import (
	"github.com/game-core/gocrafter/pkg/domain/model/friend/userFriend"
)

type FriendSendResponses []*FriendSendResponse

type FriendSendResponse struct {
	UserFriend *userFriend.UserFriend
}

func NewFriendSendResponse() *FriendSendResponse {
	return &FriendSendResponse{}
}

func NewFriendSendResponses() FriendSendResponses {
	return FriendSendResponses{}
}

func SetFriendSendResponse(userFriend *userFriend.UserFriend) *FriendSendResponse {
	return &FriendSendResponse{
		UserFriend: userFriend,
	}
}
