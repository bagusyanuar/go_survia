package lib

import (
	"gorm.io/gorm"
)

type CustomModel interface {
	FindAll(tx *gorm.DB) (*[]interface{}, error)
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
