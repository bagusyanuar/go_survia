package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
)

type Bank struct{}

func (Bank) All(q string) (res []model.Bank, err error) {
	var banks []model.Bank
	if err = database.DB.Unscoped().Model(&model.Bank{}).Where("name LIKE ?", "%"+q+"%").Find(&banks).Error; err != nil {
		return banks, err
	}
	return banks, nil
}

func (Bank) FindByID(id string) (res *model.Bank, err error) {
	var bank *model.Bank
	if err = database.DB.Model(&model.Bank{}).First(&bank, "id = ?", id).Error; err != nil {
		return bank, err
	}
	return bank, nil
}
func (Bank) Create(entity *model.Bank) error {
	if err := database.DB.Create(&entity).Error; err != nil {
		return err
	}
	return nil
}

func (Bank) Patch(id string, d interface{}) error {
	if err := database.DB.Debug().Model(&model.Bank{}).Where("id = ?", id).Updates(d).Error; err != nil {
		return err
	}
	return nil
}

func (Bank) Delete(id string) error {
	if err := database.DB.Where("id = ?", id).Delete(&model.Bank{}).Error; err != nil {
		return err
	}
	return nil
}
