package service

import (
	"go-survia/src/model"
	"go-survia/src/repositories"
	req "go-survia/src/request"
)

type Sec struct {
	repository repositories.Sec
}

func (sec *Sec) Create(r *req.Sec) error {
	entity := model.Sec{
		Name:   r.Name,
		Bottom: r.Bottom,
		Top:    r.Top,
	}
	if e := sec.repository.Create(&entity); e != nil {
		return e
	}
	return nil
}

func (sec *Sec) Patch(id string, r *req.Sec) error {
	data := map[string]interface{}{
		"name":   r.Name,
		"bottom": r.Bottom,
		"top":    r.Top,
	}
	return sec.repository.Patch(id, data)
}

func (sec *Sec) Delete(id string) error {
	return sec.repository.Delete(id)
}

func (sec *Sec) FindAll(q string) (d []model.Sec, err error) {
	return sec.repository.All(q)
}

func (sec *Sec) FindByID(id string) (d *model.Sec, err error) {
	data, e := sec.repository.FindByID(id)
	if e != nil {
		return nil, e
	}
	return data, nil
}
