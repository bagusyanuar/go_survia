package lib

import (
	"gorm.io/gorm"
)

type CustomModel interface {
	FindAll(tx *gorm.DB) (*[]interface{}, error)
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
}
