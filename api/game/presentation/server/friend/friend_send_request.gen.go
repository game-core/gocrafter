// Package friend フレンド申請リクエスト
package friend

func SetFriendSendRequest(userId string, friendUserId string) *FriendSendRequest {
	return &FriendSendRequest{
		UserId:       userId,
		FriendUserId: friendUserId,
	}
}
