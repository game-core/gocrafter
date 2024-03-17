// Package profile ユーザープロフィール
package profile

func SetUserProfile(userId string, name string, content string) *UserProfile {
	return &UserProfile{
		UserId:  userId,
		Name:    name,
		Content: content,
	}
}
