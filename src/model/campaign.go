package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Campaign struct {
	ID               uuid.UUID       `json:"id"`
	Title            string          `json:"title"`
	Description      string          `json:"description"`
	ShortDescription string          `json:"short_description"`
	Image            *string         `json:"image"`
	StartAt          *datatypes.Date `json:"start_at"`
	FinishAt         *datatypes.Date `json:"finish_at"`
	Status           uint            `json:"status"`
	Background       string          `json:"background"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	DeletedAt        gorm.DeletedAt  `json:"deleted_at"`
}

func (campaign *Campaign) BeforeCreate(tx *gorm.DB) (err error) {
	campaign.ID = uuid.New()
	campaign.CreatedAt = time.Now()
	campaign.UpdatedAt = time.Now()
	return
}

func (Campaign) TableName() string {
	return "campaigns"
}
