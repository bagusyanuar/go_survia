package controller

import (
	"errors"
	"go-survia/src/lib"
	req "go-survia/src/request"
	"go-survia/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Category struct {
	service service.Category
	request req.Category
}

func (category *Category) Index(c *gin.Context) {
	q := c.Query("q")
	data, err := category.service.FindAll(q)
	if err != nil {
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, data)
}

func (category *Category) Store(c *gin.Context) {
	c.Bind(&category.request)
	if m, e := lib.ValidateRequest(&category.request); e != nil {
		lib.JSONBadRequestResponse(c, e.Error(), m)
		return
	}

	if e := category.service.Create(&category.request); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}

func (category *Category) Show(c *gin.Context) {
	id := c.Param("id")
	data, e := category.service.FindByID(id)
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

func (category *Category) Update(c *gin.Context) {
	id := c.Param("id")
	c.Bind(&category.request)
	if m, e := lib.ValidateRequest(&category.request); e != nil {
		lib.JSONBadRequestResponse(c, e.Error(), m)
		return
	}
	if e := category.service.Patch(id, &category.request); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}

func (category *Category) Destroy(c *gin.Context) {
	id := c.Param("id")
	if e := category.service.Delete(id); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}
