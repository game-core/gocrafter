// Package userIdleBonus ユーザー放置ボーナス
//
//go:generate mockgen -source=./user_idle_bonus_repository.gen.go -destination=./user_idle_bonus_repository_mock.gen.go -package=userIdleBonus
package userIdleBonus

import (
	"context"

	"gorm.io/gorm"
)

type UserIdleBonusMysqlRepository interface {
	Find(ctx context.Context, userId string, masterIdleBonusId int64) (*UserIdleBonus, error)
	FindOrNil(ctx context.Context, userId string, masterIdleBonusId int64) (*UserIdleBonus, error)
	FindByUserId(ctx context.Context, userId string) (*UserIdleBonus, error)
	FindByUserIdAndMasterIdleBonusId(ctx context.Context, userId string, masterIdleBonusId int64) (*UserIdleBonus, error)
	FindOrNilByUserId(ctx context.Context, userId string) (*UserIdleBonus, error)
	FindOrNilByUserIdAndMasterIdleBonusId(ctx context.Context, userId string, masterIdleBonusId int64) (*UserIdleBonus, error)
	FindList(ctx context.Context, userId string) (UserIdleBonuses, error)
	FindListByUserId(ctx context.Context, userId string) (UserIdleBonuses, error)
	FindListByUserIdAndMasterIdleBonusId(ctx context.Context, userId string, masterIdleBonusId int64) (UserIdleBonuses, error)
	Create(ctx context.Context, tx *gorm.DB, m *UserIdleBonus) (*UserIdleBonus, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms UserIdleBonuses) (UserIdleBonuses, error)
	Update(ctx context.Context, tx *gorm.DB, m *UserIdleBonus) (*UserIdleBonus, error)
	Delete(ctx context.Context, tx *gorm.DB, m *UserIdleBonus) error
}
