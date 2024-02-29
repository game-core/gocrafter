// Package friend フレンド承認レスポンス
package friend

func SetFriendApproveResponse(userFriend *UserFriend) *FriendApproveResponse {
	return &FriendApproveResponse{
		UserFriend: userFriend,
	}
}
