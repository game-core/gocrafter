// Package profile プロフィール作成レスポンス
package profile

import (
	"github.com/game-core/gocrafter/pkg/domain/model/profile/userProfile"
)

type ProfileGetResponses []*ProfileGetResponse

type ProfileGetResponse struct {
	UserProfile *userProfile.UserProfile
}

func NewProfileGetResponse() *ProfileGetResponse {
	return &ProfileGetResponse{}
}

func NewProfileGetResponses() ProfileGetResponses {
	return ProfileGetResponses{}
}

func SetProfileGetResponse(userProfile *userProfile.UserProfile) *ProfileGetResponse {
	return &ProfileGetResponse{
		UserProfile: userProfile,
	}
}
