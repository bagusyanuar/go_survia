package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Province struct {
	ID        uuid.UUID      `json:"id"`
	Code      int            `json:"code"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (province *Province) BeforeCreate(tx *gorm.DB) (err error) {
	province.ID = uuid.New()
	province.CreatedAt = time.Now()
	province.UpdatedAt = time.Now()
	return
}

func (Province) TableName() string {
	return "provinces"
}
