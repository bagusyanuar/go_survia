package repositories

import (
	"errors"
	"go-survia/database"
	"go-survia/src/model"

	"gorm.io/gorm"
)

type SecQuestion struct{}

func (SecQuestion) All(q string) (d []model.SecQuestionWithAnswers, err error) {
	var data []model.SecQuestionWithAnswers
	if err = database.DB.Unscoped().
		Model(&model.SecQuestionWithAnswers{}).
		Preload("Answers", func(db *gorm.DB) *gorm.DB {
			return db.Order("index_of ASC")
		}).
		Where("question LIKE ?", "%"+q+"%").
		Order("index_of ASC").
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (SecQuestion) FindByID(id string) (d *model.SecQuestionWithAnswers, err error) {
	var data *model.SecQuestionWithAnswers
	if err = database.DB.Model(&model.SecQuestionWithAnswers{}).Preload("Answers").First(&data, "id = ?", id).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (SecQuestion) Create(entity *model.SecQuestionWithAnswers) error {
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := database.DB.Create(&entity).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (SecQuestion) Patch(id string, d interface{}) error {

	if err := database.DB.Debug().Model(&model.SecQuestionWithAnswers{}).Where("id = ?", id).Updates(d).Error; err != nil {
		return err
	}
	return nil
}

func (SecQuestion) Delete(id string) error {
	if err := database.DB.Where("id = ?", id).Delete(&model.Category{}).Error; err != nil {
		return err
	}
	return nil
}

func (SecQuestion) FindLastIndex() (i int, err error) {
	value := 0
	var lastSecQuestion *model.SecQuestion
	if err := database.DB.Model(&model.SecQuestion{}).Order("index_of DESC").First(&lastSecQuestion).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, err
		}
	}
	value = int(lastSecQuestion.IndexOf) + 1
	return value, nil
}
