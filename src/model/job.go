package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (job *Job) BeforeCreate(tx *gorm.DB) (err error) {
	job.ID = uuid.New()
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()
	return
}

func (Job) TableName() string {
	return "jobs"
}
