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
			return db.Order("index_of DESC")
		}).
		Where("question LIKE ?", "%"+q+"%").
		Order("created_at ASC").
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
		return err
	}
	return nil
}

func (SecQuestion) Patch(iid string, d interface{}) (err error) {

	// var sq *model.SecQuestion
	// qUpdate := map[string]interface{}{
	// 	"question": d.Question,
	// }

	// tx := database.DB.Begin()
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		tx.Rollback()
	// 	}
	// }()
	// if err = database.DB.Model(&sq).Where("id = ?", id).Updates(qUpdate).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	// tx.Commit()
	return nil
}

func (SecQuestion) Delete(m *model.SecQuestion) (err error) {
	if err = database.DB.Delete(&m).Error; err != nil {
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
