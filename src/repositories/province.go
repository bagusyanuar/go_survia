package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
	adminResponse "go-survia/src/response/admin"
)

type Province struct{}

func (Province) All(q string) (r []adminResponse.APIProvince, err error) {
	var provinces []adminResponse.APIProvince
	if err = database.DB.Unscoped().Model(&model.Province{}).Where("name LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&provinces).Error; err != nil {
		return provinces, err
	}
	return provinces, nil
}

func (Province) FindByID(id string) (r *adminResponse.APIProvince, err error) {
	var province *adminResponse.APIProvince
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
