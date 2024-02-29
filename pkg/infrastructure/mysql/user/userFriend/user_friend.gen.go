// Package userFriend ユーザーフレンド
package userFriend

import (
	"time"

	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type UserFriends []*UserFriend

type UserFriend struct {
	UserId       string
	FriendUserId string
	FriendType   enum.FriendType
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUserFriend() *UserFriend {
	return &UserFriend{}
}

func NewUserFriends() UserFriends {
	return UserFriends{}
}

func SetUserFriend(userId string, friendUserId string, friendType enum.FriendType, createdAt time.Time, updatedAt time.Time) *UserFriend {
	return &UserFriend{
		UserId:       userId,
		FriendUserId: friendUserId,
		FriendType:   friendType,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}

func (t *UserFriend) TableName() string {
	return "user_friend"
}
