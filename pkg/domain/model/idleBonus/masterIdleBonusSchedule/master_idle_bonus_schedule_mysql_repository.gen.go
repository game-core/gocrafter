// Package masterIdleBonusSchedule 放置ボーナススケジュール
//
//go:generate mockgen -source=./master_idle_bonus_schedule_mysql_repository.gen.go -destination=./master_idle_bonus_schedule_mysql_repository_mock.gen.go -package=masterIdleBonusSchedule
package masterIdleBonusSchedule

import (
	"context"

	"gorm.io/gorm"
)

type MasterIdleBonusScheduleMysqlRepository interface {
	Find(ctx context.Context, id int64) (*MasterIdleBonusSchedule, error)
	FindOrNil(ctx context.Context, id int64) (*MasterIdleBonusSchedule, error)
	FindByMasterIdleBonusId(ctx context.Context, masterIdleBonusId int64) (*MasterIdleBonusSchedule, error)
	FindByStep(ctx context.Context, step int32) (*MasterIdleBonusSchedule, error)
	FindByMasterIdleBonusIdAndStep(ctx context.Context, masterIdleBonusId int64, step int32) (*MasterIdleBonusSchedule, error)
	FindOrNilByMasterIdleBonusId(ctx context.Context, masterIdleBonusId int64) (*MasterIdleBonusSchedule, error)
	FindOrNilByStep(ctx context.Context, step int32) (*MasterIdleBonusSchedule, error)
	FindOrNilByMasterIdleBonusIdAndStep(ctx context.Context, masterIdleBonusId int64, step int32) (*MasterIdleBonusSchedule, error)
	FindList(ctx context.Context) (MasterIdleBonusSchedules, error)
	FindListByMasterIdleBonusId(ctx context.Context, masterIdleBonusId int64) (MasterIdleBonusSchedules, error)
	FindListByStep(ctx context.Context, step int32) (MasterIdleBonusSchedules, error)
	FindListByMasterIdleBonusIdAndStep(ctx context.Context, masterIdleBonusId int64, step int32) (MasterIdleBonusSchedules, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusSchedule) (*MasterIdleBonusSchedule, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterIdleBonusSchedules) (MasterIdleBonusSchedules, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusSchedule) (*MasterIdleBonusSchedule, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusSchedule) error
}
