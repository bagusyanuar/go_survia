package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
)

type City struct{}

func (City) All(q string) (b []model.City, err error) {
	var cities []model.City
	if err = database.DB.Unscoped().Model(&model.Province{}).Where("name LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&cities).Error; err != nil {
		return cities, err
	}
	return cities, nil
}
