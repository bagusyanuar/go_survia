package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
)

type Province struct{}

func (Province) All(q string) (b []model.Province, err error) {
	var provinces []model.Province
	if err = database.DB.Unscoped().Model(&model.Province{}).Where("name LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&provinces).Error; err != nil {
		return provinces, err
	}
	return provinces, nil
}

func (Province) FindByID(id string) (r *model.Province, err error) {
	var province *model.Province
	if err = database.DB.Model(&model.Province{}).First(&province, "id = ?", id).Error; err != nil {
		return province, err
	}
	return province, nil
}

func (Province) Create(m *model.Province) (r *model.Province, err error) {
	if err := database.DB.Create(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (Province) Patch(m *model.Province, d interface{}) (r *model.Province, err error) {
	if err = database.DB.Model(&m).Updates(d).Error; err != nil {
		return m, err
	}
	return m, nil
}

func (Province) Delete(m *model.Province) (err error) {
	if err = database.DB.Delete(&m).Error; err != nil {
		return err
	}
	return nil
}
