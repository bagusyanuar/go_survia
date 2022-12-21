package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SecAnswer struct {
	ID            uuid.UUID      `json:"id"`
	SecQuestionID uuid.UUID      `json:"sec_question_id"`
	Answer        string         `json:"answer"`
	Score         int            `json:"score"`
	IndexOf       uint           `json:"index_of"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	Question      *SecQuestion   `gorm:"foreignKey:SecQuestionID" json:"question"`
}

func (secAnswer *SecAnswer) BeforeCreate(tx *gorm.DB) (err error) {
	secAnswer.ID = uuid.New()
	secAnswer.CreatedAt = time.Now()
	secAnswer.UpdatedAt = time.Now()
	return
}

func (SecAnswer) TableName() string {
	return "sec_answers"
}
