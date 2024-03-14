// Package profile プロフィール作成レスポンス
package profile

func SetProfileCreateResponse(userProfile *UserProfile) *ProfileCreateResponse {
	return &ProfileCreateResponse{
		UserProfile: userProfile,
	}
}
