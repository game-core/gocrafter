// Package loginBonus ログインボーナスユーザー取得レスポンス
package loginBonus

func SetLoginBonusGetUserResponse(userLoginBonuses []*UserLoginBonus) *LoginBonusGetUserResponse {
	return &LoginBonusGetUserResponse{
		UserLoginBonuses: userLoginBonuses,
	}
}
