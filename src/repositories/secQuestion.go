package repositories

import (
	"go-survia/database"
	"go-survia/src/model"

	"gorm.io/gorm"
)

type SecQuestion struct{}

type apiSecQuestionResponse struct {
	model.SecQuestion
	Answers []model.SecAnswer `gorm:"foreignKey:SecQuestionID" json:"answers"`
}

func (SecQuestion) All(q string) (b []apiSecQuestionResponse, err error) {
	var secQuestions []apiSecQuestionResponse
	if err = database.DB.Unscoped().
		Model(&model.SecQuestion{}).
		Preload("Answers", func(db *gorm.DB) *gorm.DB {
			return db.Order("index_of DESC")
		}).
		Where("question LIKE ?", "%"+q+"%").
		Order("created_at ASC").
		Find(&secQuestions).Error; err != nil {
		return secQuestions, err
	}
	return secQuestions, nil
}

func (SecQuestion) FindByID(id string) (r *apiSecQuestionResponse, err error) {
	var secQuestion *apiSecQuestionResponse
	if err = database.DB.Model(&model.SecQuestion{}).Preload("Answers").First(&secQuestion, "id = ?", id).Error; err != nil {
		return secQuestion, err
	}
	return secQuestion, nil
}

func (SecQuestion) Create(m *model.SecQuestion) (r *model.SecQuestion, err error) {
	if err := database.DB.Create(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (SecQuestion) Patch(id string, d map[string]interface{}) (err error) {

	var sq *model.SecQuestion
	q := d["question"]
	qUpdate := map[string]interface{}{
		"question": q,
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

	// for _, v := range d["anwers"] {

	// }
	tx.Commit()
	return nil
}

func (SecQuestion) Delete(m *model.SecQuestion) (err error) {
	if err = database.DB.Delete(&m).Error; err != nil {
		return err
	}
	return nil
}
