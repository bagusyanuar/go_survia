package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
	adminResponse "go-survia/src/response/admin"
)

type Campaign struct{}

func (Campaign) All(q string) (b []adminResponse.APICampaign, err error) {
	var campaigns []adminResponse.APICampaign
	if err = database.DB.Unscoped().Model(&model.Campaign{}).Where("title LIKE ?", "%"+q+"%").Or("description LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&campaigns).Error; err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (Campaign) FindByID(id string) (r *adminResponse.APICampaign, err error) {
	var campaign *adminResponse.APICampaign
	if err = database.DB.Model(&model.Campaign{}).First(&campaign, "id = ?", id).Error; err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (Campaign) Create(entity *model.Campaign) error {
	if err := database.DB.Create(&entity).Error; err != nil {
		return err
	}
	return nil
}

func (Campaign) Patch(id string, d interface{}) error {
	if err := database.DB.Model(&model.Campaign{}).Where("id = ?", id).Updates(d).Error; err != nil {
		return err
	}
	return nil
}

func (Campaign) Delete(id string) error {
	if err := database.DB.Where("id = ?", id).Delete(&model.Campaign{}).Error; err != nil {
		return err
	}
	return nil
}
