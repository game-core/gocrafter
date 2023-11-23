//go:generate mockgen -source=./loginRewardModel_repository.gen.go -destination=./loginRewardModel_repository_mock.gen.go -package=loginReward
package loginReward

import (
	"github.com/game-core/gocrafter/domain/entity/master/loginReward"
	"github.com/jinzhu/gorm"
)

type LoginRewardModelRepository interface {
	Create(entity *loginReward.LoginRewardModel, tx *gorm.DB) (*loginReward.LoginRewardModel, error)

	Delete(entity *loginReward.LoginRewardModel, tx *gorm.DB) error

	FindOrNilByEventID(EventID int64) (*loginReward.LoginRewardModel, error)

	FindByID(ID int64) (*loginReward.LoginRewardModel, error)

	FindOrNilByName(Name string) (*loginReward.LoginRewardModel, error)

	FindOrNilByNameAndEventID(Name string, EventID int64) (*loginReward.LoginRewardModel, error)

	FindOrNilByID(ID int64) (*loginReward.LoginRewardModel, error)

	List(limit int64) (*loginReward.LoginRewardModels, error)

	ListByEventID(EventID int64) (*loginReward.LoginRewardModels, error)

	ListByName(Name string) (*loginReward.LoginRewardModels, error)

	ListByNameAndEventID(Name string, EventID int64) (*loginReward.LoginRewardModels, error)

	Update(entity *loginReward.LoginRewardModel, tx *gorm.DB) (*loginReward.LoginRewardModel, error)
}
