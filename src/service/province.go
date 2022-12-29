package service

import (
	"go-survia/src/model"
	"go-survia/src/repositories"
	req "go-survia/src/request"
)

type Province struct {
	repository repositories.Province
}

func (province *Province) Create(r *req.Province) error {
	entity := model.Province{
		Code: r.Code,
		Name: r.Name,
	}
	if e := province.repository.Create(&entity); e != nil {
		return e
	}
	return nil
}

func (province *Province) Patch(id string, r *req.Province) error {
	data := map[string]interface{}{
		"name": r.Name,
		"code": r.Name,
	}
	return province.repository.Patch(id, data)
}

func (province *Province) Delete(id string) error {
	return province.repository.Delete(id)
}

func (province *Province) FindAll(q string) (d []model.Province, err error) {
	return province.repository.All(q)
}

func (province *Province) FindByID(id string) (d *model.Province, err error) {
	data, e := province.repository.FindByID(id)
	if e != nil {
		return nil, e
	}
	return data, nil
}
