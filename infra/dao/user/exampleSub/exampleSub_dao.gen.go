package exampleSub

import (
	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/model/user/exampleSub"
	exampleSubRepository "github.com/game-core/gocrafter/domain/repository/user/exampleSub"
	"github.com/jinzhu/gorm"
)

type exampleSubDao struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func NewExampleSubDao(conn *database.SqlHandler) exampleSubRepository.ExampleSubRepository {
	return &exampleSubDao{
		Read:  conn.User.ReadConn,
		Write: conn.User.WriteConn,
	}
}

func (e *exampleSubDao) Create(entity *exampleSub.ExampleSub, tx *gorm.DB) (*exampleSub.ExampleSub, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = e.Write
	}

	res := conn.Model(&exampleSub.ExampleSub{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *exampleSubDao) Delete(entity *exampleSub.ExampleSub, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = e.Write
	}

	res := conn.Model(&exampleSub.ExampleSub{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (e *exampleSubDao) FindByExampleID(ExampleID int64) (*exampleSub.ExampleSub, error) {
	entity := &exampleSub.ExampleSub{}
	res := e.Read.Where("example_id = ?", ExampleID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *exampleSubDao) FindByID(ID int64) (*exampleSub.ExampleSub, error) {
	entity := &exampleSub.ExampleSub{}
	res := e.Read.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *exampleSubDao) FindByIDAndExampleID(ID int64, ExampleID int64) (*exampleSub.ExampleSub, error) {
	entity := &exampleSub.ExampleSub{}
	res := e.Read.Where("id = ?", ID).Where("example_id = ?", ExampleID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *exampleSubDao) FindByIDAndUserID(ID int64, UserID int64) (*exampleSub.ExampleSub, error) {
	entity := &exampleSub.ExampleSub{}
	res := e.Read.Where("id = ?", ID).Where("user_id = ?", UserID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *exampleSubDao) FindByIDAndUserIDAndExampleID(ID int64, UserID int64, ExampleID int64) (*exampleSub.ExampleSub, error) {
	entity := &exampleSub.ExampleSub{}
	res := e.Read.Where("id = ?", ID).Where("user_id = ?", UserID).Where("example_id = ?", ExampleID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *exampleSubDao) FindByUserID(UserID int64) (*exampleSub.ExampleSub, error) {
	entity := &exampleSub.ExampleSub{}
	res := e.Read.Where("user_id = ?", UserID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *exampleSubDao) List(limit int64) (*exampleSub.ExampleSubs, error) {
	entity := &exampleSub.ExampleSubs{}
	res := e.Read.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *exampleSubDao) Update(entity *exampleSub.ExampleSub, tx *gorm.DB) (*exampleSub.ExampleSub, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = e.Write
	}

	res := conn.Model(&exampleSub.ExampleSub{}).Where("id = ?", entity.ID).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
