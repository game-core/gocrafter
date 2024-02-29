// Package friend フレンド取得レスポンス
package friend

func SetFriendGetResponse(userFriends []*UserFriend) *FriendGetResponse {
	return &FriendGetResponse{
		UserFriends: userFriends,
	}
}
