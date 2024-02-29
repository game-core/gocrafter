// Package friend フレンド送信リクエスト
package friend

type FriendSendRequests []*FriendSendRequest

type FriendSendRequest struct {
	UserId       string
	FriendUserId string
}

func NewFriendSendRequest() *FriendSendRequest {
	return &FriendSendRequest{}
}

func NewFriendSendRequests() FriendSendRequests {
	return FriendSendRequests{}
}

func SetFriendSendRequest(userId string, friendUserId string) *FriendSendRequest {
	return &FriendSendRequest{
		UserId:       userId,
		FriendUserId: friendUserId,
	}
}
