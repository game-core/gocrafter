package example

import (
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/admin/example"
	exampleRepository "github.com/game-core/gocrafter/domain/repository/admin/example"
)

type exampleDao struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func NewExampleDao(conn *database.SqlHandler) exampleRepository.ExampleRepository {
	return &exampleDao{
		Read:  conn.Auth.ReadConn,
		Write: conn.Auth.WriteConn,
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

func (d *exampleDao) FindByIDAndName(ID int64, Name string) (*example.Example, error) {
	entity := &example.Example{}
	res := d.Read.Where("id = ?", ID).Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *exampleDao) FindByName(Name string) (*example.Example, error) {
	entity := &example.Example{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *exampleDao) FindOrNilByID(ID int64) (*example.Example, error) {
	entity := &example.Example{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *exampleDao) FindOrNilByIDAndName(ID int64, Name string) (*example.Example, error) {
	entity := &example.Example{}
	res := d.Read.Where("id = ?", ID).Where("name = ?", Name).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *exampleDao) FindOrNilByName(Name string) (*example.Example, error) {
	entity := &example.Example{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *exampleDao) List(limit int) (*example.Examples, error) {
	entity := &example.Examples{}
	res := d.Read.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *exampleDao) ListByIDAndName(ID int64, Name string) (*example.Examples, error) {
	entity := &example.Examples{}
	res := d.Read.Where("id = ?", ID).Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *exampleDao) ListByName(Name string) (*example.Examples, error) {
	entity := &example.Examples{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *exampleDao) Save(entity *example.Example, tx *gorm.DB) (*example.Example, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&example.Example{}).Where("id = ?", entity.ID).Save(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
