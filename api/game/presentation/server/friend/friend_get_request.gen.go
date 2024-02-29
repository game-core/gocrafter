// Package friend フレンド取得リクエスト
package friend

func SetFriendGetRequest(userId string) *FriendGetRequest {
	return &FriendGetRequest{
		UserId: userId,
	}
}
