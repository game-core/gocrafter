package example_sub

import (
	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/model/user/example_sub"
	example_subRepository "github.com/game-core/gocrafter/domain/repository/user/example_sub"
	"github.com/jinzhu/gorm"
)

type example_subDao struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func NewExampleSubDao(conn *database.SqlHandler) example_subRepository.ExampleSubRepository {
	return &example_subDao{
		Read:  conn.User.ReadConn,
		Write: conn.User.WriteConn,
	}
}

func (e *example_subDao) Create(entity *example_sub.ExampleSub, tx *gorm.DB) (*example_sub.ExampleSub, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = e.Write
	}

	res := conn.Model(&example_sub.ExampleSub{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *example_subDao) Delete(entity *example_sub.ExampleSub, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = e.Write
	}

	res := conn.Model(&example_sub.ExampleSub{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (e *example_subDao) FindByExampleID(ExampleID int64) (*example_sub.ExampleSub, error) {
	entity := &example_sub.ExampleSub{}
	res := e.Read.Where("example_id = ?", ExampleID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *example_subDao) FindByID(ID int64) (*example_sub.ExampleSub, error) {
	entity := &example_sub.ExampleSub{}
	res := e.Read.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *example_subDao) FindByIDAndExampleID(ID int64, ExampleID int64) (*example_sub.ExampleSub, error) {
	entity := &example_sub.ExampleSub{}
	res := e.Read.Where("id = ?", ID).Where("example_id = ?", ExampleID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *example_subDao) FindByIDAndUserID(ID int64, UserID int64) (*example_sub.ExampleSub, error) {
	entity := &example_sub.ExampleSub{}
	res := e.Read.Where("id = ?", ID).Where("user_id = ?", UserID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *example_subDao) FindByIDAndUserIDAndExampleID(ID int64, UserID int64, ExampleID int64) (*example_sub.ExampleSub, error) {
	entity := &example_sub.ExampleSub{}
	res := e.Read.Where("id = ?", ID).Where("user_id = ?", UserID).Where("example_id = ?", ExampleID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *example_subDao) FindByUserID(UserID int64) (*example_sub.ExampleSub, error) {
	entity := &example_sub.ExampleSub{}
	res := e.Read.Where("user_id = ?", UserID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *example_subDao) List(limit int64) (*example_sub.ExampleSubs, error) {
	entity := &example_sub.ExampleSubs{}
	res := e.Read.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *example_subDao) Update(entity *example_sub.ExampleSub, tx *gorm.DB) (*example_sub.ExampleSub, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = e.Write
	}

	res := conn.Model(&example_sub.ExampleSub{}).Where("id = ?", entity.ID).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
