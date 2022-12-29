package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
)

type Province struct{}

func (Province) All(q string) (res []model.Province, err error) {
	var provinces []model.Province
	if err = database.DB.Unscoped().Model(&model.Province{}).Where("name LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&provinces).Error; err != nil {
		return provinces, err
	}
	return provinces, nil
}

func (Province) FindByID(id string) (res *model.Province, err error) {
	var province *model.Province
	if err = database.DB.Model(&model.Province{}).First(&province, "id = ?", id).Error; err != nil {
		return province, err
	}
	return province, nil
}

func (Province) Create(m *model.Province) error {
	if err := database.DB.Create(&m).Error; err != nil {
		return err
	}
	return nil
}

func (Province) Patch(id string, d interface{}) error {
	if err := database.DB.Debug().Model(&model.Province{}).Where("id = ?", id).Updates(d).Error; err != nil {
		return err
	}
	return nil
}

func (Province) Delete(id string) error {
	if err := database.DB.Where("id = ?", id).Delete(&model.Province{}).Error; err != nil {
		return err
	}
	return nil
}
