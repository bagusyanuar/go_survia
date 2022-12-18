package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
)

type Campaign struct{}

func (Campaign) All(q string) (b []model.Campaign, err error) {
	var campaigns []model.Campaign
	if err = database.DB.Unscoped().Model(&model.Campaign{}).Where("title LIKE ?", "%"+q+"%").Or("description LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&campaigns).Error; err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (Campaign) FindByID(id string) (r *model.Campaign, err error) {
	var campaign *model.Campaign
	if err = database.DB.Model(&model.Campaign{}).First(&campaign, "id = ?", id).Error; err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (Campaign) Create(m *model.Campaign) (r *model.Campaign, err error) {
	if err := database.DB.Create(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (Campaign) Patch(m *model.Campaign, d interface{}) (r *model.Campaign, err error) {
	if err = database.DB.Model(&m).Updates(d).Error; err != nil {
		return m, err
	}
	return m, nil
}

func (Campaign) Delete(m *model.Campaign) (err error) {
	if err = database.DB.Delete(&m).Error; err != nil {
		return err
	}
	return nil
}
