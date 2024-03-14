// Package profile プロフィール作成リクエスト
package profile

func SetProfileCreateRequest(userId string, name string, content string) *ProfileCreateRequest {
	return &ProfileCreateRequest{
		UserId:  userId,
		Name:    name,
		Content: content,
	}
}
