package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bank struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Code      int            `json:"code"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (bank *Bank) BeforeCreate(tx *gorm.DB) (err error) {
	bank.ID = uuid.New()
	bank.CreatedAt = time.Now()
	bank.UpdatedAt = time.Now()
	return
}

func (Bank) TableName() string {
	return "banks"
}