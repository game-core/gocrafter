// Package loginBonus ログインボーナスマスター取得レスポンス
package loginBonus

func SetLoginBonusGetMasterResponse(masterLoginBonus *MasterLoginBonus, masterLoginBonusEvent *MasterLoginBonusEvent, masterLoginBonusItems []*MasterLoginBonusItem, masterLoginBonusSchedules []*MasterLoginBonusSchedule) *LoginBonusGetMasterResponse {
	return &LoginBonusGetMasterResponse{
		MasterLoginBonus:          masterLoginBonus,
		MasterLoginBonusEvent:     masterLoginBonusEvent,
		MasterLoginBonusItems:     masterLoginBonusItems,
		MasterLoginBonusSchedules: masterLoginBonusSchedules,
	}
}
