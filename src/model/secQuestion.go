package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SecQuestion struct {
	ID        uuid.UUID      `json:"id"`
	Question  string         `json:"question"`
	IndexOf   uint           `json:"index_of"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (secQuestion *SecQuestion) BeforeCreate(tx *gorm.DB) (err error) {
	secQuestion.ID = uuid.New()
	secQuestion.CreatedAt = time.Now()
	secQuestion.UpdatedAt = time.Now()
	return
}

func (SecQuestion) TableName() string {
	return "sec_questions"
}

type SecQuestionWithAnswers struct {
	ID        uuid.UUID      `json:"id"`
	Question  string         `json:"question"`
	IndexOf   uint           `json:"index_of"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Answers   []SecAnswer    `gorm:"foreignKey:SecQuestionID" json:"answers"`
}

func (secQuestionWithAnswers *SecQuestionWithAnswers) BeforeCreate(tx *gorm.DB) (err error) {
	secQuestionWithAnswers.ID = uuid.New()
	secQuestionWithAnswers.CreatedAt = time.Now()
	secQuestionWithAnswers.UpdatedAt = time.Now()
	return
}

func (SecQuestionWithAnswers) TableName() string {
	return "sec_questions"
}
