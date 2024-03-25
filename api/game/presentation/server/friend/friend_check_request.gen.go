// Package friend フレンド確認リクエスト
package friend

func SetFriendCheckRequest(userId string, friendUserId string) *FriendCheckRequest {
	return &FriendCheckRequest{
		UserId:       userId,
		FriendUserId: friendUserId,
	}
}
