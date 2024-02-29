// Package friend ユーザーフレンド
package friend

import (
	"github.com/game-core/gocrafter/pkg/domain/model/friend/userFriend"
)

func SetUserFriends(userFriendModels userFriend.UserFriends) []*UserFriend {
	var userFriends []*UserFriend
	for _, userFriendModel := range userFriendModels {
		userFriends = append(
			userFriends,
			SetUserFriend(
				userFriendModel.UserId,
				userFriendModel.FriendUserId,
				FriendType(userFriendModel.FriendType),
			),
		)
	}

	return userFriends
}
