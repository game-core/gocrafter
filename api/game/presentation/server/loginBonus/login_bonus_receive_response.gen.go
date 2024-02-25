// Package loginBonus ログインボーナス受け取りレスポンス
package loginBonus

func SetLoginBonusReceiveResponse(userLoginBonus *UserLoginBonus, masterLoginBonus *MasterLoginBonus, masterLoginBonusEvent *MasterLoginBonusEvent, masterLoginBonusItems []*MasterLoginBonusItem, masterLoginBonusSchedule *MasterLoginBonusSchedule) *LoginBonusReceiveResponse {
	return &LoginBonusReceiveResponse{
		UserLoginBonus:           userLoginBonus,
		MasterLoginBonus:         masterLoginBonus,
		MasterLoginBonusEvent:    masterLoginBonusEvent,
		MasterLoginBonusItems:    masterLoginBonusItems,
		MasterLoginBonusSchedule: masterLoginBonusSchedule,
	}
}
