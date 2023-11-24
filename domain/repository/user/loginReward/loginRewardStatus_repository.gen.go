//go:generate mockgen -source=./loginRewardStatus_repository.gen.go -destination=./loginRewardStatus_repository_mock.gen.go -package=loginReward
package loginReward

import (
	"github.com/game-core/gocrafter/domain/entity/user/loginReward"
	"gorm.io/gorm"
)

type LoginRewardStatusRepository interface {
	Create(entity *loginReward.LoginRewardStatus, shardKey string, tx *gorm.DB) (*loginReward.LoginRewardStatus, error)

	Delete(entity *loginReward.LoginRewardStatus, shardKey string, tx *gorm.DB) error

	FindByAccountID(AccountID int64, shardKey string) (*loginReward.LoginRewardStatus, error)

	FindByAccountIDAndLoginRewardModelName(AccountID int64, LoginRewardModelName string, shardKey string) (*loginReward.LoginRewardStatus, error)

	FindByID(ID int64, shardKey string) (*loginReward.LoginRewardStatus, error)

	FindByLoginRewardModelName(LoginRewardModelName string, shardKey string) (*loginReward.LoginRewardStatus, error)

	FindOrNilByAccountID(AccountID int64, shardKey string) (*loginReward.LoginRewardStatus, error)

	FindOrNilByAccountIDAndLoginRewardModelName(AccountID int64, LoginRewardModelName string, shardKey string) (*loginReward.LoginRewardStatus, error)

	FindOrNilByID(ID int64, shardKey string) (*loginReward.LoginRewardStatus, error)

	FindOrNilByLoginRewardModelName(LoginRewardModelName string, shardKey string) (*loginReward.LoginRewardStatus, error)

	List(limit int, shardKey string) (*loginReward.LoginRewardStatuses, error)

	ListByAccountID(AccountID int64, shardKey string) (*loginReward.LoginRewardStatuses, error)

	ListByAccountIDAndLoginRewardModelName(AccountID int64, LoginRewardModelName string, shardKey string) (*loginReward.LoginRewardStatuses, error)

	ListByLoginRewardModelName(LoginRewardModelName string, shardKey string) (*loginReward.LoginRewardStatuses, error)

	Save(entity *loginReward.LoginRewardStatus, shardKey string, tx *gorm.DB) (*loginReward.LoginRewardStatus, error)
}
