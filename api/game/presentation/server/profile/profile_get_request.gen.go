// Package profile プロフィール取得リクエスト
package profile

func SetProfileGetRequest(userId string) *ProfileGetRequest {
	return &ProfileGetRequest{
		UserId: userId,
	}
}
