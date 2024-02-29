// Package friend フレンド承認拒否レスポンス
package friend

func SetFriendDisapproveResponse(userFriend *UserFriend) *FriendDisapproveResponse {
	return &FriendDisapproveResponse{
		UserFriend: userFriend,
	}
}
