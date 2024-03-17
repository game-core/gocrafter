// Package idleBonus 放置ボーナスマスター取得リクエスト
package idleBonus

type IdleBonusGetMasterRequests []*IdleBonusGetMasterRequest

type IdleBonusGetMasterRequest struct {
	MasterIdleBonusId int64
}

func NewIdleBonusGetMasterRequest() *IdleBonusGetMasterRequest {
	return &IdleBonusGetMasterRequest{}
}

func NewIdleBonusGetMasterRequests() IdleBonusGetMasterRequests {
	return IdleBonusGetMasterRequests{}
}

func SetIdleBonusGetMasterRequest(masterIdleBonusId int64) *IdleBonusGetMasterRequest {
	return &IdleBonusGetMasterRequest{
		MasterIdleBonusId: masterIdleBonusId,
	}
}
