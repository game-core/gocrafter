// Package friend フレンド確認レスポンス
package friend

import (
	"github.com/game-core/gocrafter/pkg/domain/model/friend/userFriend"
)

type FriendCheckResponses []*FriendCheckResponse

type FriendCheckResponse struct {
	UserFriend *userFriend.UserFriend
}

func NewFriendCheckResponse() *FriendCheckResponse {
	return &FriendCheckResponse{}
}

func NewFriendCheckResponses() FriendCheckResponses {
	return FriendCheckResponses{}
}

func SetFriendCheckResponse(userFriend *userFriend.UserFriend) *FriendCheckResponse {
	return &FriendCheckResponse{
		UserFriend: userFriend,
	}
}
