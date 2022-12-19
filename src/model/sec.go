package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sec struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Bottom    int            `json:"bottom"`
	Top       int            `json:"top"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (sec *Sec) BeforeCreate(tx *gorm.DB) (err error) {
	sec.ID = uuid.New()
	sec.CreatedAt = time.Now()
	sec.UpdatedAt = time.Now()
	return
}

func (Sec) TableName() string {
	return "secs"
}