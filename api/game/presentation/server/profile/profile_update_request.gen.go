// Package profile プロフィール更新リクエスト
package profile

func SetProfileUpdateRequest(userId string, name string, content string) *ProfileUpdateRequest {
	return &ProfileUpdateRequest{
		UserId:  userId,
		Name:    name,
		Content: content,
	}
}
