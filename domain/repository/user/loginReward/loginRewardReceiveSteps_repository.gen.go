//go:generate mockgen -source=./loginRewardReceiveSteps_repository.gen.go -destination=./loginRewardReceiveSteps_repository_mock.gen.go -package=loginReward
package loginReward

import (
	"github.com/game-core/gocrafter/domain/entity/user/loginReward"
	"github.com/jinzhu/gorm"
)

type LoginRewardReceiveStepsRepository interface {
	Create(entity *loginReward.LoginRewardReceiveSteps, shardKey int, tx *gorm.DB) (*loginReward.LoginRewardReceiveSteps, error)

	Delete(entity *loginReward.LoginRewardReceiveSteps, shardKey int, tx *gorm.DB) error

	FindByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error)

	FindByAccountIDAndLoginRewardStatusID(AccountID int64, LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error)

	FindByID(ID int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error)

	FindByLoginRewardStatusID(LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error)

	List(limit int64, shardKey int) (*loginReward.LoginRewardReceiveStepses, error)

	ListByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardReceiveStepses, error)

	ListByAccountIDAndLoginRewardStatusID(AccountID int64, LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveStepses, error)

	ListByLoginRewardStatusID(LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveStepses, error)

	Update(entity *loginReward.LoginRewardReceiveSteps, shardKey int, tx *gorm.DB) (*loginReward.LoginRewardReceiveSteps, error)
}
