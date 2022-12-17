package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
)

type Bank struct{}

var listBanks []model.Bank
var bank *model.Bank

func (Bank) All(q string) (b []model.Bank, err error) {
	if err = database.DB.Unscoped().Model(&model.Bank{}).Where("name LIKE ?", "%"+q+"%").Find(&listBanks).Error; err != nil {
		return listBanks, err
	}
	return listBanks, nil
}

func (Bank) FindByID(id string) (r *model.Bank, err error) {
	if err = database.DB.Model(&model.Bank{}).First(&bank, "id = ?", id).Error; err != nil {
		return bank, err
	}
	return bank, nil
}
func (Bank) Create(m *model.Bank) (r *model.Bank, err error) {
	if err := database.DB.Create(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (Bank) Patch(m *model.Bank, d interface{}) (r *model.Bank, err error) {
	if err = database.DB.Model(&m).Updates(d).Error; err != nil {
		return m, err
	}
	return m, nil
}

func (Bank) Delete(m *model.Bank) (err error) {
	if err = database.DB.Delete(&m).Error; err != nil {
		return err
	}
	return nil
}
