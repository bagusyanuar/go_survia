package lib

import (
	"gorm.io/gorm"
)

type CustomModel interface {
	FindAll(tx *gorm.DB) (*[]interface{}, error)
}
