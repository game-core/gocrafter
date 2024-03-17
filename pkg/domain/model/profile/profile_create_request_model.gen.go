// Package profile プロフィール作成リクエスト
package profile

type ProfileCreateRequests []*ProfileCreateRequest

type ProfileCreateRequest struct {
	UserId  string
	Name    string
	Content string
}

func NewProfileCreateRequest() *ProfileCreateRequest {
	return &ProfileCreateRequest{}
}

func NewProfileCreateRequests() ProfileCreateRequests {
	return ProfileCreateRequests{}
}

func SetProfileCreateRequest(userId string, name string, content string) *ProfileCreateRequest {
	return &ProfileCreateRequest{
		UserId:  userId,
		Name:    name,
		Content: content,
	}
}
