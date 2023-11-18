package example

import (
	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/master/example"
	exampleRepository "github.com/game-core/gocrafter/domain/repository/master/example"
	"github.com/jinzhu/gorm"
)

type exampleDao struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func NewExampleDao(conn *database.SqlHandler) exampleRepository.ExampleRepository {
	return &exampleDao{
		Read:  conn.Master.ReadConn,
		Write: conn.Master.WriteConn,
	}
}

func (d *exampleDao) Create(entity *example.Example, tx *gorm.DB) (*example.Example, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&example.Example{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *exampleDao) Delete(entity *example.Example, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&example.Example{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (d *exampleDao) FindByID(ID int64) (*example.Example, error) {
	entity := &example.Example{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *exampleDao) FindByIDAndUserID(ID int64, UserID int64) (*example.Example, error) {
	entity := &example.Example{}
	res := d.Read.Where("id = ?", ID).Where("user_id = ?", UserID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *exampleDao) FindByUserID(UserID int64) (*example.Example, error) {
	entity := &example.Example{}
	res := d.Read.Where("user_id = ?", UserID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *exampleDao) List(limit int64) (*example.Examples, error) {
	entity := &example.Examples{}
	res := d.Read.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *exampleDao) Update(entity *example.Example, tx *gorm.DB) (*example.Example, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&example.Example{}).Where("id = ?", entity.ID).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
