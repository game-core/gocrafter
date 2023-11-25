//go:generate mockgen -source=./loginRewardItem_repository.gen.go -destination=./loginRewardItem_repository_mock.gen.go -package=loginReward
package loginReward

import (
	"github.com/game-core/gocrafter/domain/entity/master/loginReward"
	"gorm.io/gorm"
)

type LoginRewardItemRepository interface {
	Create(entity *loginReward.LoginRewardItem, tx *gorm.DB) (*loginReward.LoginRewardItem, error)

	Delete(entity *loginReward.LoginRewardItem, tx *gorm.DB) error

	FindByID(ID int64) (*loginReward.LoginRewardItem, error)

	FindByName(Name string) (*loginReward.LoginRewardItem, error)

	FindOrNilByID(ID int64) (*loginReward.LoginRewardItem, error)

	FindOrNilByName(Name string) (*loginReward.LoginRewardItem, error)

	List(limit int) (*loginReward.LoginRewardItems, error)

	ListByName(Name string) (*loginReward.LoginRewardItems, error)

	Save(entity *loginReward.LoginRewardItem, tx *gorm.DB) (*loginReward.LoginRewardItem, error)
}
