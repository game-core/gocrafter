// Package loginBonus ログインボーナス受け取りリクエスト
package loginBonus

func SetLoginBonusReceiveRequest(userId string, masterLoginBonusId int64) *LoginBonusReceiveRequest {
	return &LoginBonusReceiveRequest{
		UserId:             userId,
		MasterLoginBonusId: masterLoginBonusId,
	}
}
