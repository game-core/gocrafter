// Package friend フレンド削除レスポンス
package friend

func SetFriendDeleteResponse(userFriend *UserFriend) *FriendDeleteResponse {
	return &FriendDeleteResponse{
		UserFriend: userFriend,
	}
}
