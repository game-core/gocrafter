// Package masterAction アクション
package masterAction

type MasterActions []*MasterAction

type MasterAction struct {
	Id              int64
	Name            string
	ActionType      string
	AnyId           int64
	TriggerActionId *int64
	NextActionId    *int64
}

func NewMasterAction() *MasterAction {
	return &MasterAction{}
}

func NewMasterActions() MasterActions {
	return MasterActions{}
}

func SetMasterAction(id int64, name string, actionType string, anyId int64, triggerActionId *int64, nextActionId *int64) *MasterAction {
	return &MasterAction{
		Id:              id,
		Name:            name,
		ActionType:      actionType,
		AnyId:           anyId,
		TriggerActionId: triggerActionId,
		NextActionId:    nextActionId,
	}
}

func (t *MasterAction) TableName() string {
	return "master_action"
}
