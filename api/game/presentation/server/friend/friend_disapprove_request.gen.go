// Package friend フレンド承認拒否リクエスト
package friend

func SetFriendDisapproveRequest(userId string, friendUserId string) *FriendDisapproveRequest {
	return &FriendDisapproveRequest{
		UserId:       userId,
		FriendUserId: friendUserId,
	}
}
