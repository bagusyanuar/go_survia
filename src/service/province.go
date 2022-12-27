package service

import (
	"go-survia/src/lib"
	"go-survia/src/model"
	"go-survia/src/repositories"
	adminRequest "go-survia/src/request/admin"
)

type Province struct {
	repository repositories.Province
}

func (province *Province) FindAll(response interface{}, query string, table string)  {
	
}

func (province *Province) Create(request *adminRequest.AdminProvince) error {
	entity := model.Province{
		Code: request.Code,
		Name: request.Name,
	}
	if e := province.repository.Create(&entity); e != nil {
		return e
	}
	return nil
}

func (province *Province) ValidateRequest(request *adminRequest.AdminProvince) (m interface{}, err error) {
	messages, e := lib.ValidateRequest(request)
	if e != nil {
		return messages, lib.ErrBadRequest
	}
	return nil, nil
}
