package admin

import (
	"go-survia/src/model"
)

type APISecQuestionResponse struct {
	model.SecQuestion
	Answers   []model.SecAnswer `gorm:"foreignKey:SecQuestionID" json:"answers"`
}