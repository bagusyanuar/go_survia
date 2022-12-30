package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
)

type Sec struct{}

func (Sec) All(q string) (d []model.Sec, err error) {
	var secs []model.Sec
	if err = database.DB.Unscoped().Model(&model.Sec{}).Where("name LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&secs).Error; err != nil {
		return secs, err
	}
	return secs, nil
}

func (Sec) FindByID(id string) (d *model.Sec, err error) {
	var sec *model.Sec
	if err = database.DB.Model(&model.Sec{}).First(&sec, "id = ?", id).Error; err != nil {
		return sec, err
	}
	return sec, nil
}

func (Sec) Create(entity *model.Sec) error {
	if err := database.DB.Create(&entity).Error; err != nil {
		return err
	}
	return nil
}

func (Sec) Patch(id string, d interface{}) error {
	if err := database.DB.Debug().Model(&model.Sec{}).Where("id = ?", id).Updates(d).Error; err != nil {
		return err
	}
	return nil
}

func (Sec) Delete(id string) (err error) {
	if err := database.DB.Where("id = ?", id).Delete(&model.Sec{}).Error; err != nil {
		return err
	}
	return nil
}
