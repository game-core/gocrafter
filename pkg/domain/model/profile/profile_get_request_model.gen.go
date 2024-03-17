// Package profile プロフィール取得リクエスト
package profile

type ProfileGetRequests []*ProfileGetRequest

type ProfileGetRequest struct {
	UserId string
}

func NewProfileGetRequest() *ProfileGetRequest {
	return &ProfileGetRequest{}
}

func NewProfileGetRequests() ProfileGetRequests {
	return ProfileGetRequests{}
}

func SetProfileGetRequest(userId string) *ProfileGetRequest {
	return &ProfileGetRequest{
		UserId: userId,
	}
}
