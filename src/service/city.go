package service

import (
	"go-survia/src/model"
	"go-survia/src/repositories"
	req "go-survia/src/request"

	"github.com/google/uuid"
)

type City struct {
	repository repositories.City
}

func (city *City) Create(r *req.City) error {

	provinceID, e := uuid.Parse(r.ProvinceID)
	if e != nil {
		return e
	}
	entity := model.City{
		ProvinceID: provinceID,
		Code:       r.Code,
		Name:       r.Name,
	}
	if e := city.repository.Create(&entity); e != nil {
		return e
	}
	return nil
}

func (city *City) FindAll(q string) (d []model.CityWithProvince, err error) {
	return city.repository.All(q)
}
