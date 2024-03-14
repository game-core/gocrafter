// Package userAction ユーザーアクション
//
//go:generate mockgen -source=./user_action_repository.gen.go -destination=./user_action_repository_mock.gen.go -package=userAction
package userAction

import (
	"context"

	"gorm.io/gorm"
)

type UserActionRepository interface {
	Find(ctx context.Context, userId string, masterActionId int64) (*UserAction, error)
	FindOrNil(ctx context.Context, userId string, masterActionId int64) (*UserAction, error)
	FindList(ctx context.Context, userId string) (UserActions, error)
	Create(ctx context.Context, tx *gorm.DB, m *UserAction) (*UserAction, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms UserActions) (UserActions, error)
	Update(ctx context.Context, tx *gorm.DB, m *UserAction) (*UserAction, error)
	Delete(ctx context.Context, tx *gorm.DB, m *UserAction) error
}
