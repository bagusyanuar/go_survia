package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
)

type City struct{}

func (City) All(q string) (b []model.City, err error) {
	var cities []model.City
	if err = database.DB.Unscoped().Model(&model.City{}).Where("name LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&cities).Error; err != nil {
		return cities, err
	}
	return cities, nil
}

func (City) FindByID(id string) (r *model.City, err error) {
	var city *model.City
	if err = database.DB.Model(&model.City{}).First(&city, "id = ?", id).Error; err != nil {
		return city, err
	}
	return city, nil
}

func (City) Create(m *model.City) (r *model.City, err error) {
	if err := database.DB.Create(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (City) Patch(m *model.City, d interface{}) (r *model.City, err error) {
	if err = database.DB.Model(&m).Updates(d).Error; err != nil {
		return m, err
	}
	return m, nil
}

func (City) Delete(m *model.City) (err error) {
	if err = database.DB.Delete(&m).Error; err != nil {
		return err
	}
	return nil
}
