// Package profile プロフィール作成レスポンス
package profile

func SetProfileGetResponse(userProfile *UserProfile) *ProfileGetResponse {
	return &ProfileGetResponse{
		UserProfile: userProfile,
	}
}
