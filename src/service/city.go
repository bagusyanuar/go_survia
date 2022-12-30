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

func (city *City) Patch(id string, r *req.City) error {
	provinceID, e := uuid.Parse(r.ProvinceID)
	if e != nil {
		return e
	}
	data := map[string]interface{}{
		"name":        r.Name,
		"code":        r.Code,
		"province_id": provinceID,
	}
	return city.repository.Patch(id, data)
}

func (city *City) Delete(id string) error {
	return city.repository.Delete(id)
}

func (city *City) FindAll(q string) (d []model.CityWithProvince, err error) {
	return city.repository.All(q)
}

func (city *City) FindByID(id string) (d *model.CityWithProvince, err error) {
	data, e := city.repository.FindByID(id)
	if e != nil {
		return nil, e
	}
	return data, nil
}
