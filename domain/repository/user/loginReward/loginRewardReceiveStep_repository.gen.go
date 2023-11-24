//go:generate mockgen -source=./loginRewardReceiveStep_repository.gen.go -destination=./loginRewardReceiveStep_repository_mock.gen.go -package=loginReward
package loginReward

import (
	"github.com/game-core/gocrafter/domain/entity/user/loginReward"
	"gorm.io/gorm"
)

type LoginRewardReceiveStepRepository interface {
	Create(entity *loginReward.LoginRewardReceiveStep, shardKey string, tx *gorm.DB) (*loginReward.LoginRewardReceiveStep, error)

	Delete(entity *loginReward.LoginRewardReceiveStep, shardKey string, tx *gorm.DB) error

	FindByAccountID(AccountID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error)

	FindByAccountIDAndLoginRewardStatusID(AccountID int64, LoginRewardStatusID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error)

	FindByID(ID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error)

	FindByLoginRewardStatusID(LoginRewardStatusID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error)

	FindOrNilByAccountID(AccountID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error)

	FindOrNilByAccountIDAndLoginRewardStatusID(AccountID int64, LoginRewardStatusID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error)

	FindOrNilByID(ID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error)

	FindOrNilByLoginRewardStatusID(LoginRewardStatusID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error)

	List(limit int, shardKey string) (*loginReward.LoginRewardReceiveSteps, error)

	ListByAccountID(AccountID int64, shardKey string) (*loginReward.LoginRewardReceiveSteps, error)

	ListByAccountIDAndLoginRewardStatusID(AccountID int64, LoginRewardStatusID int64, shardKey string) (*loginReward.LoginRewardReceiveSteps, error)

	ListByLoginRewardStatusID(LoginRewardStatusID int64, shardKey string) (*loginReward.LoginRewardReceiveSteps, error)

	Save(entity *loginReward.LoginRewardReceiveStep, shardKey string, tx *gorm.DB) (*loginReward.LoginRewardReceiveStep, error)
}
