// Package friend フレンド承認リクエスト
package friend

type FriendApproveRequests []*FriendApproveRequest

type FriendApproveRequest struct {
	UserId       string
	FriendUserId string
}

func NewFriendApproveRequest() *FriendApproveRequest {
	return &FriendApproveRequest{}
}

func NewFriendApproveRequests() FriendApproveRequests {
	return FriendApproveRequests{}
}

func SetFriendApproveRequest(userId string, friendUserId string) *FriendApproveRequest {
	return &FriendApproveRequest{
		UserId:       userId,
		FriendUserId: friendUserId,
	}
}
