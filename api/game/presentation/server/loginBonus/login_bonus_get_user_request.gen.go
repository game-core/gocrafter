// Package loginBonus ログインボーナスユーザー取得リクエスト
package loginBonus

func SetLoginBonusGetUserRequest(userId string) *LoginBonusGetUserRequest {
	return &LoginBonusGetUserRequest{
		UserId: userId,
	}
}
