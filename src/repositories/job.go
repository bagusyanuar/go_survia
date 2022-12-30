package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
)

type Job struct{}

func (Job) All(q string) (d []model.Job, err error) {
	var jobs []model.Job
	if err = database.DB.Unscoped().Model(&model.Job{}).Where("name LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&jobs).Error; err != nil {
		return jobs, err
	}
	return jobs, nil
}

func (Job) FindByID(id string) (d *model.Job, err error) {
	var job *model.Job
	if err = database.DB.Model(&model.Job{}).First(&job, "id = ?", id).Error; err != nil {
		return job, err
	}
	return job, nil
}

func (Job) Create(entity *model.Job) error {
	if err := database.DB.Create(&entity).Error; err != nil {
		return err
	}
	return nil
}

func (Job) Patch(id string, d interface{}) error {
	if err := database.DB.Debug().Model(&model.Job{}).Where("id = ?", id).Updates(d).Error; err != nil {
		return err
	}
	return nil
}

func (Job) Delete(id string) error {
	if err := database.DB.Where("id = ?", id).Delete(&model.Job{}).Error; err != nil {
		return err
	}
	return nil
}
