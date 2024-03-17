// Package profile プロフィール更新レスポンス
package profile

func SetProfileUpdateResponse(userProfile *UserProfile) *ProfileUpdateResponse {
	return &ProfileUpdateResponse{
		UserProfile: userProfile,
	}
}
