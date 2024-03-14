// Package userAction ユーザーアクション
package userAction

import (
	"time"
)

type UserActions []*UserAction

type UserAction struct {
	UserId         string
	Name           string
	MasterActionId int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewUserAction() *UserAction {
	return &UserAction{}
}

func NewUserActions() UserActions {
	return UserActions{}
}

func SetUserAction(userId string, name string, masterActionId int64, createdAt time.Time, updatedAt time.Time) *UserAction {
	return &UserAction{
		UserId:         userId,
		Name:           name,
		MasterActionId: masterActionId,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}

func (t *UserAction) TableName() string {
	return "user_action"
}
