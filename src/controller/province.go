package controller

import (
	"errors"
	"go-survia/src/lib"
	req "go-survia/src/request"
	"go-survia/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Province struct {
	service service.Province
	request req.Province
}

func (province *Province) Index(c *gin.Context) {
	q := c.Query("q")
	data, err := province.service.FindAll(q)
	if err != nil {
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, data)
}

func (province *Province) Store(c *gin.Context) {
	c.Bind(&province.request)
	if m, e := lib.ValidateRequest(&province.request); e != nil {
		lib.JSONBadRequestResponse(c, e.Error(), m)
		return
	}

	if e := province.service.Create(&province.request); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}

func (province *Province) Show(c *gin.Context) {
	id := c.Param("id")
	data, e := province.service.FindByID(id)
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

func (province *Province) Update(c *gin.Context) {
	id := c.Param("id")
	c.Bind(&province.request)
	if m, e := lib.ValidateRequest(&province.request); e != nil {
		lib.JSONBadRequestResponse(c, e.Error(), m)
		return
	}
	if e := province.service.Patch(id, &province.request); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}

func (province *Province) Destroy(c *gin.Context) {
	id := c.Param("id")
	if e := province.service.Delete(id); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}
