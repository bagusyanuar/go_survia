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

func (category *Category) Create(request *adminRequest.AdminCategoryRequest) (data interface{}, err error) {

	d, e := lib.ValidateRequest(request)
	if e != nil {
		return d, e
	}
	entity := model.Category{
		Name: request.Name,
	}
	m, e := category.repository.Create(&entity)
	if e != nil {
		return nil, e
	}
	return m, nil
}
func (c *Category) FindAll(q string) (b []adminResponse.APICategory, err error) {
	return c.repository.All(q)
}
