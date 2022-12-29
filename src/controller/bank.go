package controller

import (
	"errors"
	"go-survia/src/lib"
	req "go-survia/src/request"
	"go-survia/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Bank struct {
	service service.Bank
	request req.Bank
}

func (bank *Bank) Index(c *gin.Context) {
	q := c.Query("q")
	data, err := bank.service.FindAll(q)
	if err != nil {
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, data)
}

func (bank *Bank) Store(c *gin.Context) {
	c.Bind(&bank.request)
	if m, e := lib.ValidateRequest(&bank.request); e != nil {
		lib.JSONBadRequestResponse(c, e.Error(), m)
		return
	}

	if e := bank.service.Create(&bank.request); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}

func (bank *Bank) Show(c *gin.Context) {
	id := c.Param("id")
	data, e := bank.service.FindByID(id)
	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			lib.JSONNotFoundResponse(c, e.Error(), nil)
			return
		}
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, data)
}

func (bank *Bank) Update(c *gin.Context) {
	id := c.Param("id")
	c.Bind(&bank.request)
	if m, e := lib.ValidateRequest(&bank.request); e != nil {
		lib.JSONBadRequestResponse(c, e.Error(), m)
		return
	}
	if e := bank.service.Patch(id, &bank.request); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}

func (bank *Bank) Destroy(c *gin.Context) {
	id := c.Param("id")
	if e := bank.service.Delete(id); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}
