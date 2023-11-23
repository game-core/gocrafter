//go:generate mockgen -source=./loginRewardReceiveStep_repository.gen.go -destination=./loginRewardReceiveStep_repository_mock.gen.go -package=loginReward
package loginReward

import (
	"github.com/game-core/gocrafter/domain/entity/user/loginReward"
	"github.com/jinzhu/gorm"
)

type LoginRewardReceiveStepRepository interface {
	Create(entity *loginReward.LoginRewardReceiveStep, shardKey int, tx *gorm.DB) (*loginReward.LoginRewardReceiveStep, error)

	Delete(entity *loginReward.LoginRewardReceiveStep, shardKey int, tx *gorm.DB) error

	FindOrNilByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardReceiveStep, error)

	FindOrNilByAccountIDAndLoginRewardStatusID(AccountID int64, LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveStep, error)

	FindByID(ID int64, shardKey int) (*loginReward.LoginRewardReceiveStep, error)

	FindOrNilByLoginRewardStatusID(LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveStep, error)

	FindOrNilByID(ID int64, shardKey int) (*loginReward.LoginRewardReceiveStep, error)

	List(limit int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error)

	ListByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error)

	ListByAccountIDAndLoginRewardStatusID(AccountID int64, LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error)

	ListByLoginRewardStatusID(LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error)

	Update(entity *loginReward.LoginRewardReceiveStep, shardKey int, tx *gorm.DB) (*loginReward.LoginRewardReceiveStep, error)
}
