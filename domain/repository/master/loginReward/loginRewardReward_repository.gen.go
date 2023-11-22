//go:generate mockgen -source=./loginRewardReward_repository.gen.go -destination=./loginRewardReward_repository_mock.gen.go -package=loginReward
package loginReward

import (
	"github.com/game-core/gocrafter/domain/entity/master/loginReward"
	"github.com/jinzhu/gorm"
)

type LoginRewardRewardRepository interface {
	Create(entity *loginReward.LoginRewardReward, tx *gorm.DB) (*loginReward.LoginRewardReward, error)

	Delete(entity *loginReward.LoginRewardReward, tx *gorm.DB) error

	FindByID(ID int64) (*loginReward.LoginRewardReward, error)

	FindByItemID(ItemID int64) (*loginReward.LoginRewardReward, error)

	FindByLoginRewardID(LoginRewardID int64) (*loginReward.LoginRewardReward, error)

	FindByLoginRewardIDAndItemID(LoginRewardID int64, ItemID int64) (*loginReward.LoginRewardReward, error)

	FindByName(Name string) (*loginReward.LoginRewardReward, error)

	List(limit int64) (*loginReward.LoginRewardRewards, error)

	ListByItemID(ItemID int64) (*loginReward.LoginRewardRewards, error)

	ListByLoginRewardID(LoginRewardID int64) (*loginReward.LoginRewardRewards, error)

	ListByLoginRewardIDAndItemID(LoginRewardID int64, ItemID int64) (*loginReward.LoginRewardRewards, error)

	ListByName(Name string) (*loginReward.LoginRewardRewards, error)

	Update(entity *loginReward.LoginRewardReward, tx *gorm.DB) (*loginReward.LoginRewardReward, error)
}