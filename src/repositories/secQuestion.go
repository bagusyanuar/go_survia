package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
	adminResponse "go-survia/src/response/admin"

	"gorm.io/gorm"
)

type SecQuestion struct{}

type apiSecQuestionResponse struct {
	model.SecQuestion
	Answers []model.SecAnswer `gorm:"foreignKey:SecQuestionID" json:"answers"`
}

func (SecQuestion) All(q string) (b []adminResponse.APISecQuestionResponse, err error) {
	var response []adminResponse.APISecQuestionResponse
	if err = database.DB.Unscoped().Debug().
		Model(&model.SecQuestion{}).
		Preload("Answers", func(db *gorm.DB) *gorm.DB {
			return db.Order("index_of DESC")
		}).
		Where("question LIKE ?", "%"+q+"%").
		Order("created_at ASC").
		Find(&response).Error; err != nil {
		return response, err
	}
	return response, nil
}

func (SecQuestion) FindByID(id string) (r *adminResponse.APISecQuestionResponse, err error) {
	var response *adminResponse.APISecQuestionResponse
	if err = database.DB.Model(&model.SecQuestion{}).Preload("Answers").First(&response, "id = ?", id).Error; err != nil {
		return response, err
	}
	return response, nil
}

func (SecQuestion) Create(m *model.SecQuestion) (r *model.SecQuestion, err error) {
	if err := database.DB.Create(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (SecQuestion) Patch(id string, d *apiSecQuestionResponse) (err error) {

	var sq *model.SecQuestion
	qUpdate := map[string]interface{}{
		"question": d.Question,
	}

	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = database.DB.Model(&sq).Where("id = ?", id).Updates(qUpdate).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (SecQuestion) Delete(m *model.SecQuestion) (err error) {
	if err = database.DB.Delete(&m).Error; err != nil {
		return err
	}
	return nil
}
