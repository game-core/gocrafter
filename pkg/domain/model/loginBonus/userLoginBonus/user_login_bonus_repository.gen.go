// Package userLoginBonus ユーザーログインボーナス
//
//go:generate mockgen -source=./user_login_bonus_repository.gen.go -destination=./user_login_bonus_repository_mock.gen.go -package=userLoginBonus
package userLoginBonus

import (
	"context"

	"gorm.io/gorm"
)

type UserLoginBonusRepository interface {
	Find(ctx context.Context, userId string, masterLoginBonusId int64) (*UserLoginBonus, error)
	FindOrNil(ctx context.Context, userId string, masterLoginBonusId int64) (*UserLoginBonus, error)
	FindByUserId(ctx context.Context, userId string) (*UserLoginBonus, error)
	FindByUserIdAndMasterLoginBonusId(ctx context.Context, userId string, masterLoginBonusId int64) (*UserLoginBonus, error)
	FindOrNilByUserId(ctx context.Context, userId string) (*UserLoginBonus, error)
	FindOrNilByUserIdAndMasterLoginBonusId(ctx context.Context, userId string, masterLoginBonusId int64) (*UserLoginBonus, error)
	FindList(ctx context.Context, userId string) (UserLoginBonuses, error)
	FindListByUserId(ctx context.Context, userId string) (UserLoginBonuses, error)
	FindListByUserIdAndMasterLoginBonusId(ctx context.Context, userId string, masterLoginBonusId int64) (UserLoginBonuses, error)
	Create(ctx context.Context, tx *gorm.DB, m *UserLoginBonus) (*UserLoginBonus, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms UserLoginBonuses) (UserLoginBonuses, error)
	Update(ctx context.Context, tx *gorm.DB, m *UserLoginBonus) (*UserLoginBonus, error)
	Delete(ctx context.Context, tx *gorm.DB, m *UserLoginBonus) error
}
