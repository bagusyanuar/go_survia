package admin

import (
	"errors"
	"go-survia/src/lib"
	request "go-survia/src/request/admin"
	"go-survia/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Bank struct {
	service service.Bank
	request request.AdminBank
}

func (bank *Bank) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		c.Bind(&bank.request)
		messages, err := bank.service.Create(&bank.request)
		if err != nil {
			if errors.Is(err, lib.ErrBadRequest) {
				lib.JSONBadRequestResponse(c, err.Error(), messages)
				return
			}
			lib.JSONErrorResponse(c, err.Error(), nil)
			return
		}
		lib.JSONSuccessResponse(c, nil)
		return
	}
	q := c.Query("q")
	data, err := bank.service.FindAll(q)
	if err != nil {
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, data)
}

func (bank Bank) FindByID(c *gin.Context) {
	id := c.Param("id")

	//update method
	if c.Request.Method == "PATCH" {
		c.Bind(&bank.request)
		messages, err := bank.service.Patch(id, &bank.request)
		if err != nil {
			if errors.Is(err, lib.ErrBadRequest) {
				lib.JSONBadRequestResponse(c, err.Error(), messages)
				return
			}
			lib.JSONErrorResponse(c, err.Error(), nil)
			return
		}
		lib.JSONSuccessResponse(c, nil)
		return
	}

	//delete method
	if c.Request.Method == "DELETE" {
		err := bank.service.Delete(id)
		if err != nil {
			lib.JSONErrorResponse(c, err.Error(), nil)
			return
		}
		lib.JSONSuccessResponse(c, nil)
		return
	}

	data, err := bank.service.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			lib.JSONNotFoundResponse(c, err.Error(), nil)
			return
		}
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, data)
}
