// Package friend フレンド削除リクエスト
package friend

func SetFriendDeleteRequest(userId string, friendUserId string) *FriendDeleteRequest {
	return &FriendDeleteRequest{
		UserId:       userId,
		FriendUserId: friendUserId,
	}
}
