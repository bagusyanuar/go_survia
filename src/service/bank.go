package service

import (
	"go-survia/src/lib"
	"go-survia/src/model"
	"go-survia/src/repositories"
	adminRequest "go-survia/src/request/admin"
	adminResponse "go-survia/src/response/admin"
)

type Bank struct {
	repository repositories.Bank
}

func (bank *Bank) Create(request *adminRequest.AdminBank) (d interface{}, err error) {
	messages, e := lib.ValidateRequest(request)
	if e != nil {
		return messages, lib.ErrBadRequest
	}
	entity := model.Bank{
		Name: request.Name,
		Code: request.Code,
	}
	e = bank.repository.Create(&entity)
	if e != nil {
		return nil, e
	}
	return nil, nil
}

func (bank *Bank) Patch(id string, request *adminRequest.AdminBank) (d interface{}, err error) {
	messages, e := lib.ValidateRequest(request)
	if e != nil {
		return messages, lib.ErrBadRequest
	}
	data := map[string]interface{}{
		"name": request.Name,
		"code": request.Code,
	}
	return nil, bank.repository.Patch(id, data)
}

func (bank *Bank) Delete(id string) error {
	return bank.repository.Delete(id)
}

func (bank *Bank) FindAll(q string) (b []adminResponse.APIBank, err error) {
	return bank.repository.All(q)
}

func (bank *Bank) FindByID(id string) (r *adminResponse.APIBank, err error) {
	entity, e := bank.repository.FindByID(id)
	if e != nil {
		return nil, e
	}
	return entity, nil
}
