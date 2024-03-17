// Package userAction ユーザーアクション
package userAction

import (
	"time"
)

type UserActions []*UserAction

type UserAction struct {
	UserId         string
	MasterActionId int64
	StartedAt      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewUserAction() *UserAction {
	return &UserAction{}
}

func NewUserActions() UserActions {
	return UserActions{}
}

func SetUserAction(userId string, masterActionId int64, startedAt time.Time, createdAt time.Time, updatedAt time.Time) *UserAction {
	return &UserAction{
		UserId:         userId,
		MasterActionId: masterActionId,
		StartedAt:      startedAt,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}

func (t *UserAction) TableName() string {
	return "user_action"
}
