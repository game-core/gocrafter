// Package userFriend ユーザーフレンド
package userFriend

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type UserFriends []*UserFriend

type UserFriend struct {
	UserId       string
	FriendUserId string
	FriendType   enum.FriendType
}

func NewUserFriend() *UserFriend {
	return &UserFriend{}
}

func NewUserFriends() UserFriends {
	return UserFriends{}
}

func SetUserFriend(userId string, friendUserId string, friendType enum.FriendType) *UserFriend {
	return &UserFriend{
		UserId:       userId,
		FriendUserId: friendUserId,
		FriendType:   friendType,
	}
}
