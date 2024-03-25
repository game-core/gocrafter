// Package friend フレンド確認リクエスト
package friend

type FriendCheckRequests []*FriendCheckRequest

type FriendCheckRequest struct {
	UserId       string
	FriendUserId string
}

func NewFriendCheckRequest() *FriendCheckRequest {
	return &FriendCheckRequest{}
}

func NewFriendCheckRequests() FriendCheckRequests {
	return FriendCheckRequests{}
}

func SetFriendCheckRequest(userId string, friendUserId string) *FriendCheckRequest {
	return &FriendCheckRequest{
		UserId:       userId,
		FriendUserId: friendUserId,
	}
}
