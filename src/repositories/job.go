package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
)

type Job struct{}

func (Job) All(q string) (b []model.Job, err error) {
	var jobs []model.Job
	if err = database.DB.Unscoped().Model(&model.Job{}).Where("name LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&jobs).Error; err != nil {
		return jobs, err
	}
	return jobs, nil
}

func (Job) FindByID(id string) (r *model.Job, err error) {
	var job *model.Job
	if err = database.DB.Model(&model.Job{}).First(&job, "id = ?", id).Error; err != nil {
		return job, err
	}
	return job, nil
}

func (Job) Create(m *model.Job) (r *model.Job, err error) {
	if err := database.DB.Create(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (Job) Patch(m *model.Job, d interface{}) (r *model.Job, err error) {
	if err = database.DB.Model(&m).Updates(d).Error; err != nil {
		return m, err
	}
	return m, nil
}

func (Job) Delete(m *model.Job) (err error) {
	if err = database.DB.Delete(&m).Error; err != nil {
		return err
	}
	return nil
}
