// Package loginBonus ログインボーナスマスター取得リクエスト
package loginBonus

func SetLoginBonusGetMasterRequest(userId string, masterLoginBonusId int64) *LoginBonusGetMasterRequest {
	return &LoginBonusGetMasterRequest{
		UserId:             userId,
		MasterLoginBonusId: masterLoginBonusId,
	}
}
