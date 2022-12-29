package service

import (
	"go-survia/src/model"
	"go-survia/src/repositories"
	req "go-survia/src/request"
)

type Category struct {
	repository repositories.Category
}

func (category *Category) Create(r *req.Category) error {
	entity := model.Category{
		Name: r.Name,
	}
	if e := category.repository.Create(&entity); e != nil {
		return e
	}
	return nil
}

func (category *Category) Patch(id string, r *req.Category) error {
	data := map[string]interface{}{
		"name": r.Name,
	}
	return category.repository.Patch(id, data)
}

func (category *Category) Delete(id string) error {
	return category.repository.Delete(id)
}

func (category *Category) FindAll(q string) (d []model.Category, err error) {
	return category.repository.All(q)
}

func (category *Category) FindByID(id string) (d *model.Category, err error) {
	data, e := category.repository.FindByID(id)
	if e != nil {
		return nil, e
	}
	return data, nil
}
