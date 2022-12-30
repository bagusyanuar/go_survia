package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
)

type City struct{}

func (City) All(q string) (d []model.CityWithProvince, err error) {
	var cities []model.CityWithProvince
	if err = database.DB.Debug().Unscoped().
		Model(&model.CityWithProvince{}).
		Joins("INNER JOIN `provinces` `province` ON (`province`.`id` = `cities`.`province_id`)").
		Preload("Province").
		Where(
			database.DB.Where("cities.name LIKE ?", "%"+q+"%").Or("province.name LIKE ?", "%"+q+"%"),
		).
		Order("created_at ASC").Find(&cities).Error; err != nil {
		return cities, err
	}
	return cities, nil
}

func (City) FindByID(id string) (d *model.CityWithProvince, err error) {
	var city *model.CityWithProvince
	if err = database.DB.Preload("Province").Model(&model.CityWithProvince{}).First(&city, "id = ?", id).Error; err != nil {
		return city, err
	}
	return city, nil
}

func (City) Create(m *model.City) error {
	if err := database.DB.Create(&m).Error; err != nil {
		return err
	}
	return nil
}

func (City) Patch(id string, d interface{}) error {
	if err := database.DB.Debug().Model(&model.City{}).Where("id = ?", id).Updates(d).Error; err != nil {
		return err
	}
	return nil
}

func (City) Delete(id string) error {
	if err := database.DB.Where("id = ?", id).Delete(&model.City{}).Error; err != nil {
		return err
	}
	return nil
}
