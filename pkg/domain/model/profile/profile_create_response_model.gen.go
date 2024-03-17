// Package profile プロフィール作成レスポンス
package profile

import (
	"github.com/game-core/gocrafter/pkg/domain/model/profile/userProfile"
)

type ProfileCreateResponses []*ProfileCreateResponse

type ProfileCreateResponse struct {
	UserProfile *userProfile.UserProfile
}

func NewProfileCreateResponse() *ProfileCreateResponse {
	return &ProfileCreateResponse{}
}

func NewProfileCreateResponses() ProfileCreateResponses {
	return ProfileCreateResponses{}
}

func SetProfileCreateResponse(userProfile *userProfile.UserProfile) *ProfileCreateResponse {
	return &ProfileCreateResponse{
		UserProfile: userProfile,
	}
}
