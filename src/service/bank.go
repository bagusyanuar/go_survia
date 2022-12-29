package service

import (
	"go-survia/src/model"
	"go-survia/src/repositories"
	req "go-survia/src/request"
)

type Bank struct {
	repository repositories.Bank
}

func (bank *Bank) Create(r *req.Bank) error {
	entity := model.Bank{
		Name: r.Name,
		Code: r.Code,
	}
	if e := bank.repository.Create(&entity); e != nil {
		return e
	}
	return nil
}

func (bank *Bank) Patch(id string, r *req.Bank) error {
	data := map[string]interface{}{
		"name": r.Name,
		"code": r.Code,
	}
	return bank.repository.Patch(id, data)
}

func (bank *Bank) Delete(id string) error {
	return bank.repository.Delete(id)
}

func (bank *Bank) FindAll(q string) (d []model.Bank, err error) {
	return bank.repository.All(q)
}

func (bank *Bank) FindByID(id string) (d *model.Bank, err error) {
	data, e := bank.repository.FindByID(id)
	if e != nil {
		return nil, e
	}
	return data, nil
}
