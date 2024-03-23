// Package masterLoginBonusSchedule ログインボーナススケジュール
//
//go:generate mockgen -source=./master_login_bonus_schedule_mysql_repository.gen.go -destination=./master_login_bonus_schedule_mysql_repository_mock.gen.go -package=masterLoginBonusSchedule
package masterLoginBonusSchedule

import (
	"context"

	"gorm.io/gorm"
)

type MasterLoginBonusScheduleMysqlRepository interface {
	Find(ctx context.Context, id int64) (*MasterLoginBonusSchedule, error)
	FindOrNil(ctx context.Context, id int64) (*MasterLoginBonusSchedule, error)
	FindByMasterLoginBonusId(ctx context.Context, masterLoginBonusId int64) (*MasterLoginBonusSchedule, error)
	FindByStep(ctx context.Context, step int32) (*MasterLoginBonusSchedule, error)
	FindByMasterLoginBonusIdAndStep(ctx context.Context, masterLoginBonusId int64, step int32) (*MasterLoginBonusSchedule, error)
	FindOrNilByMasterLoginBonusId(ctx context.Context, masterLoginBonusId int64) (*MasterLoginBonusSchedule, error)
	FindOrNilByStep(ctx context.Context, step int32) (*MasterLoginBonusSchedule, error)
	FindOrNilByMasterLoginBonusIdAndStep(ctx context.Context, masterLoginBonusId int64, step int32) (*MasterLoginBonusSchedule, error)
	FindList(ctx context.Context) (MasterLoginBonusSchedules, error)
	FindListByMasterLoginBonusId(ctx context.Context, masterLoginBonusId int64) (MasterLoginBonusSchedules, error)
	FindListByStep(ctx context.Context, step int32) (MasterLoginBonusSchedules, error)
	FindListByMasterLoginBonusIdAndStep(ctx context.Context, masterLoginBonusId int64, step int32) (MasterLoginBonusSchedules, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusSchedule) (*MasterLoginBonusSchedule, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonusSchedules) (MasterLoginBonusSchedules, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusSchedule) (*MasterLoginBonusSchedule, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusSchedule) error
}
