// Package friend フレンド承認拒否リクエスト
package friend

type FriendDisapproveRequests []*FriendDisapproveRequest

type FriendDisapproveRequest struct {
	UserId       string
	FriendUserId string
}

func NewFriendDisapproveRequest() *FriendDisapproveRequest {
	return &FriendDisapproveRequest{}
}

func NewFriendDisapproveRequests() FriendDisapproveRequests {
	return FriendDisapproveRequests{}
}

func SetFriendDisapproveRequest(userId string, friendUserId string) *FriendDisapproveRequest {
	return &FriendDisapproveRequest{
		UserId:       userId,
		FriendUserId: friendUserId,
	}
}
