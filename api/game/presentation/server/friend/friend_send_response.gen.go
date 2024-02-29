// Package friend フレンド申請レスポンス
package friend

func SetFriendSendResponse(userFriend *UserFriend) *FriendSendResponse {
	return &FriendSendResponse{
		UserFriend: userFriend,
	}
}
