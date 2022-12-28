package service

import (
	"go-survia/src/lib"
	"go-survia/src/model"
	"go-survia/src/repositories"
	adminRequest "go-survia/src/request/admin"
	adminResponse "go-survia/src/response/admin"
)

type Category struct {
	repository repositories.Category
}

func (category *Category) Create(request *adminRequest.AdminCategory) (d interface{}, err error) {
	messages, e := lib.ValidateRequest(request)
	if e != nil {
		return messages, lib.ErrBadRequest
	}
	entity := model.Category{
		Name: request.Name,
	}
	e = category.repository.Create(&entity)
	if e != nil {
		return nil, e
	}
	return nil, nil
}

func (category *Category) Patch(id string, request *adminRequest.AdminCategory) (d interface{}, err error) {
	messages, e := lib.ValidateRequest(request)
	if e != nil {
		return messages, lib.ErrBadRequest
	}
	data := map[string]interface{}{
		"name": request.Name,
	}
	return nil, category.repository.Patch(id, data)
}

func (category *Category) Delete(id string) error {
	return category.repository.Delete(id)
}

func (category *Category) FindAll(q string) (res []model.Category, err error) {
	return category.repository.All(q)
}

func (category *Category) FindByID(id string) (r *adminResponse.APICategory, err error) {
	entity, e := category.repository.FindByID(id)
	if e != nil {
		return nil, e
	}
	return entity, nil
}
