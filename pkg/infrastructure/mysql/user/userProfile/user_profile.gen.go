// Package userProfile ユーザープロフィール
package userProfile

import (
	"time"
)

type UserProfiles []*UserProfile

type UserProfile struct {
	UserId    string
	Name      string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserProfile() *UserProfile {
	return &UserProfile{}
}

func NewUserProfiles() UserProfiles {
	return UserProfiles{}
}

func SetUserProfile(userId string, name string, content string, createdAt time.Time, updatedAt time.Time) *UserProfile {
	return &UserProfile{
		UserId:    userId,
		Name:      name,
		Content:   content,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (t *UserProfile) TableName() string {
	return "user_profile"
}
