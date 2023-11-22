//go:generate mockgen -source=./loginReward_repository.gen.go -destination=./loginReward_repository_mock.gen.go -package=loginReward
package loginReward

import (
	"github.com/game-core/gocrafter/domain/entity/user/loginReward"
	"github.com/jinzhu/gorm"
)

type LoginRewardStatusRepository interface {
	Create(entity *loginReward.LoginRewardStatus, shardKey int, tx *gorm.DB) (*loginReward.LoginRewardStatus, error)

	Delete(entity *loginReward.LoginRewardStatus, shardKey int, tx *gorm.DB) error

	FindByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardStatus, error)

	FindByAccountIDAndLoginRewardModelID(AccountID int64, LoginRewardModelID int64, shardKey int) (*loginReward.LoginRewardStatus, error)

	FindByID(ID int64, shardKey int) (*loginReward.LoginRewardStatus, error)

	FindByLoginRewardModelID(LoginRewardModelID int64, shardKey int) (*loginReward.LoginRewardStatus, error)

	List(limit int64, shardKey int) (*loginReward.LoginRewardStatuss, error)

	ListByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardStatuss, error)

	ListByAccountIDAndLoginRewardModelID(AccountID int64, LoginRewardModelID int64, shardKey int) (*loginReward.LoginRewardStatuss, error)

	ListByLoginRewardModelID(LoginRewardModelID int64, shardKey int) (*loginReward.LoginRewardStatuss, error)

	Update(entity *loginReward.LoginRewardStatus, shardKey int, tx *gorm.DB) (*loginReward.LoginRewardStatus, error)
}
