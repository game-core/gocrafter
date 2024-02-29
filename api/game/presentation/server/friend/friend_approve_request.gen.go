// Package friend フレンド承認リクエスト
package friend

func SetFriendApproveRequest(userId string, friendUserId string) *FriendApproveRequest {
	return &FriendApproveRequest{
		UserId:       userId,
		FriendUserId: friendUserId,
	}
}
