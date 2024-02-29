// Package friend フレンド削除リクエスト
package friend

type FriendDeleteRequests []*FriendDeleteRequest

type FriendDeleteRequest struct {
	UserId       string
	FriendUserId string
}

func NewFriendDeleteRequest() *FriendDeleteRequest {
	return &FriendDeleteRequest{}
}

func NewFriendDeleteRequests() FriendDeleteRequests {
	return FriendDeleteRequests{}
}

func SetFriendDeleteRequest(userId string, friendUserId string) *FriendDeleteRequest {
	return &FriendDeleteRequest{
		UserId:       userId,
		FriendUserId: friendUserId,
	}
}
