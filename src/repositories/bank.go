package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
	request "go-survia/src/request/admin"
	adminResponse "go-survia/src/response/admin"
)

type Bank struct{}

func (Bank) All(q string) (b []adminResponse.APIBankResponse, err error) {
	var banks []adminResponse.APIBankResponse
	if err = database.DB.Unscoped().Model(&model.Bank{}).Where("name LIKE ?", "%"+q+"%").Find(&banks).Error; err != nil {
		return banks, err
	}
	return banks, nil
}

func (Bank) FindByID(id string) (r *adminResponse.APIBankResponse, err error) {
	var bank *adminResponse.APIBankResponse
	if err = database.DB.Model(&model.Bank{}).First(&bank, "id = ?", id).Error; err != nil {
		return bank, err
	}
	return bank, nil
}
func (Bank) Create(request *request.AdminBankRequest) (r *model.Bank, err error) {
	m := model.Bank{
		Code: request.Code,
		Name: request.Name,
	}
	if err := database.DB.Create(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (Bank) Patch(id string, d interface{}) (err error) {
	var bank *model.Bank
	if err = database.DB.Model(&bank).Where("id = ?", id).Updates(d).Error; err != nil {
		return err
	}
	return nil
}

func (Bank) Delete(id string) (err error) {
	var bank *model.Bank
	if err = database.DB.Where("id = ?", id).Delete(&bank).Error; err != nil {
		return err
	}
	return nil
}
