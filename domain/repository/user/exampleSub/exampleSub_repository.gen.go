package exampleSub

import (
	"github.com/architecture-template/echo-ddd/domain/model/user/exampleSub"
	"github.com/jinzhu/gorm"
)

type ExampleSubRepository interface {
	Create(exampleSub *exampleSub.ExampleSub, tx *gorm.DB) (*exampleSub.ExampleSub, error)

	Delete(exampleSub *exampleSub.ExampleSub, tx *gorm.DB) error

	FindByExampleID(ExampleID int64) (*exampleSub.ExampleSub, error)

	FindByID(ID int64) (*exampleSub.ExampleSub, error)

	FindByIDAndExampleID(ID int64, ExampleID int64) (*exampleSub.ExampleSub, error)

	FindByIDAndUserID(ID int64, UserID int64) (*exampleSub.ExampleSub, error)

	FindByIDAndUserIDAndExampleID(ID int64, UserID int64, ExampleID int64) (*exampleSub.ExampleSub, error)

	FindByUserID(UserID int64) (*exampleSub.ExampleSub, error)

	List(limit int64) (*exampleSub.ExampleSubs, error)

	Update(exampleSub *exampleSub.ExampleSub, tx *gorm.DB) (*exampleSub.ExampleSub, error)
}
