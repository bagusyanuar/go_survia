package admin

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
