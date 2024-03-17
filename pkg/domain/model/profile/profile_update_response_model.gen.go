// Package profile プロフィール更新レスポンス
package profile

import (
	"github.com/game-core/gocrafter/pkg/domain/model/profile/userProfile"
)

type ProfileUpdateResponses []*ProfileUpdateResponse

type ProfileUpdateResponse struct {
	UserProfile *userProfile.UserProfile
}

func NewProfileUpdateResponse() *ProfileUpdateResponse {
	return &ProfileUpdateResponse{}
}

func NewProfileUpdateResponses() ProfileUpdateResponses {
	return ProfileUpdateResponses{}
}

func SetProfileUpdateResponse(userProfile *userProfile.UserProfile) *ProfileUpdateResponse {
	return &ProfileUpdateResponse{
		UserProfile: userProfile,
	}
}
