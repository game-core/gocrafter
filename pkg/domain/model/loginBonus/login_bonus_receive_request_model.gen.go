// Package loginBonus ログインボーナス受け取りリクエスト
package loginBonus

type LoginBonusReceiveRequests []*LoginBonusReceiveRequest

type LoginBonusReceiveRequest struct {
	UserId             string
	MasterLoginBonusId int64
}

func NewLoginBonusReceiveRequest() *LoginBonusReceiveRequest {
	return &LoginBonusReceiveRequest{}
}

func NewLoginBonusReceiveRequests() LoginBonusReceiveRequests {
	return LoginBonusReceiveRequests{}
}

func SetLoginBonusReceiveRequest(userId string, masterLoginBonusId int64) *LoginBonusReceiveRequest {
	return &LoginBonusReceiveRequest{
		UserId:             userId,
		MasterLoginBonusId: masterLoginBonusId,
	}
}
