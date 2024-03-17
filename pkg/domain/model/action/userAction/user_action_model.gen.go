// Package userAction ユーザーアクション
package userAction

type UserActions []*UserAction

type UserAction struct {
	UserId         string
	Name           string
	MasterActionId int64
}

func NewUserAction() *UserAction {
	return &UserAction{}
}

func NewUserActions() UserActions {
	return UserActions{}
}

func SetUserAction(userId string, name string, masterActionId int64) *UserAction {
	return &UserAction{
		UserId:         userId,
		Name:           name,
		MasterActionId: masterActionId,
	}
}
