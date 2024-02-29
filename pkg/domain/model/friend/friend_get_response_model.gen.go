// Package friend フレンド取得レスポンス
package friend

import (
	"github.com/game-core/gocrafter/pkg/domain/model/friend/userFriend"
)

type FriendGetResponses []*FriendGetResponse

type FriendGetResponse struct {
	UserFriends userFriend.UserFriends
}

func NewFriendGetResponse() *FriendGetResponse {
	return &FriendGetResponse{}
}

func NewFriendGetResponses() FriendGetResponses {
	return FriendGetResponses{}
}

func SetFriendGetResponse(userFriends userFriend.UserFriends) *FriendGetResponse {
	return &FriendGetResponse{
		UserFriends: userFriends,
	}
}
