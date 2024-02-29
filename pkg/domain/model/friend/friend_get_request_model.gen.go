// Package friend フレンド取得リクエスト
package friend

type FriendGetRequests []*FriendGetRequest

type FriendGetRequest struct {
	UserId string
}

func NewFriendGetRequest() *FriendGetRequest {
	return &FriendGetRequest{}
}

func NewFriendGetRequests() FriendGetRequests {
	return FriendGetRequests{}
}

func SetFriendGetRequest(userId string) *FriendGetRequest {
	return &FriendGetRequest{
		UserId: userId,
	}
}
