// Package loginBonus ログインボーナスユーザー取得リクエスト
package loginBonus

type LoginBonusGetUserRequests []*LoginBonusGetUserRequest

type LoginBonusGetUserRequest struct {
	UserId string
}

func NewLoginBonusGetUserRequest() *LoginBonusGetUserRequest {
	return &LoginBonusGetUserRequest{}
}

func NewLoginBonusGetUserRequests() LoginBonusGetUserRequests {
	return LoginBonusGetUserRequests{}
}

func SetLoginBonusGetUserRequest(userId string) *LoginBonusGetUserRequest {
	return &LoginBonusGetUserRequest{
		UserId: userId,
	}
}
