package userFriend

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

// GetFriends フレンド一覧を取得する
func (s UserFriends) GetFriends() UserFriends {
	var userFriendModels UserFriends
	for _, userFriendModel := range s {
		if userFriendModel.FriendType == enum.FriendType_Approved {
			userFriendModels = append(userFriendModels, userFriendModel)
		}
	}

	return userFriendModels
}
