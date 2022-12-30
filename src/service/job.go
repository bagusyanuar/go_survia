package service

import (
	"go-survia/src/model"
	"go-survia/src/repositories"
	req "go-survia/src/request"
)

type Job struct {
	repository repositories.Job
}

func (job *Job) Create(r *req.Job) error {
	entity := model.Job{
		Name: r.Name,
	}
	if e := job.repository.Create(&entity); e != nil {
		return e
	}
	return nil
}

func (job *Job) Patch(id string, r *req.Job) error {
	data := map[string]interface{}{
		"name": r.Name,
	}
	return job.repository.Patch(id, data)
}

func (job *Job) Delete(id string) error {
	return job.repository.Delete(id)
}

func (job *Job) FindAll(q string) (d []model.Job, err error) {
	return job.repository.All(q)
}

func (job *Job) FindByID(id string) (d *model.Job, err error) {
	data, e := job.repository.FindByID(id)
	if e != nil {
		return nil, e
	}
	return data, nil
}
