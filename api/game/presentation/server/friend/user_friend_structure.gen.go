// Package friend ユーザーフレンド
package friend

func SetUserFriend(userId string, friendUserId string, friendType FriendType) *UserFriend {
	return &UserFriend{
		UserId:       userId,
		FriendUserId: friendUserId,
		FriendType:   friendType,
	}
}
