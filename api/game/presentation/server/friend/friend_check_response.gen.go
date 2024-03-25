// Package friend フレンド確認レスポンス
package friend

func SetFriendCheckResponse(userFriend *UserFriend) *FriendCheckResponse {
	return &FriendCheckResponse{
		UserFriend: userFriend,
	}
}
