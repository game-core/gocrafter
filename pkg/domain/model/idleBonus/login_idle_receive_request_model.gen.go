// Package idleBonus 放置ボーナス受け取りリクエスト
package idleBonus

type IdleBonusReceiveRequests []*IdleBonusReceiveRequest

type IdleBonusReceiveRequest struct {
	UserId            string
	MasterIdleBonusId int64
}

func NewIdleBonusReceiveRequest() *IdleBonusReceiveRequest {
	return &IdleBonusReceiveRequest{}
}

func NewIdleBonusReceiveRequests() IdleBonusReceiveRequests {
	return IdleBonusReceiveRequests{}
}

func SetIdleBonusReceiveRequest(userId string, masterIdleBonusId int64) *IdleBonusReceiveRequest {
	return &IdleBonusReceiveRequest{
		UserId:            userId,
		MasterIdleBonusId: masterIdleBonusId,
	}
}
