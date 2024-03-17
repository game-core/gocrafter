// Package masterActionRun 実行されるアクション
package masterActionRun

type MasterActionRuns []*MasterActionRun

type MasterActionRun struct {
	Id       int64
	Name     string
	ActionId int64
}

func NewMasterActionRun() *MasterActionRun {
	return &MasterActionRun{}
}

func NewMasterActionRuns() MasterActionRuns {
	return MasterActionRuns{}
}

func SetMasterActionRun(id int64, name string, actionId int64) *MasterActionRun {
	return &MasterActionRun{
		Id:       id,
		Name:     name,
		ActionId: actionId,
	}
}

func (t *MasterActionRun) TableName() string {
	return "master_action_run"
}
