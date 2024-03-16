// Package loginBonus ログインボーナスマスター取得リクエスト
package loginBonus

type LoginBonusGetMasterRequests []*LoginBonusGetMasterRequest

type LoginBonusGetMasterRequest struct {
	MasterLoginBonusId int64
}

func NewLoginBonusGetMasterRequest() *LoginBonusGetMasterRequest {
	return &LoginBonusGetMasterRequest{}
}

func NewLoginBonusGetMasterRequests() LoginBonusGetMasterRequests {
	return LoginBonusGetMasterRequests{}
}

func SetLoginBonusGetMasterRequest(masterLoginBonusId int64) *LoginBonusGetMasterRequest {
	return &LoginBonusGetMasterRequest{
		MasterLoginBonusId: masterLoginBonusId,
	}
}
