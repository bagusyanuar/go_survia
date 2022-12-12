package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID    uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"user_id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	IsActive  bool           `gorm:"type:tinyint(1);not null" json:"is_active"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

func (admin *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	admin.ID = uuid.New()
	admin.CreatedAt = time.Now()
	admin.UpdatedAt = time.Now()
	return
}