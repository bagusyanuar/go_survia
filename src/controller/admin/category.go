package admin

import (
	"errors"
	"go-survia/src/lib"
	request "go-survia/src/request/admin"
	"go-survia/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Category struct {
	service service.Category
	request request.AdminCategory
}

func (category *Category) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		c.Bind(&category.request)
		messages, err := category.service.Create(&category.request)
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
	data, err := category.service.FindAll(q)
	if err != nil {
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, data)
}

func (category *Category) FindByID(c *gin.Context) {
	id := c.Param("id")

	//update method
	if c.Request.Method == "PATCH" {
		c.Bind(&category.request)
		messages, err := category.service.Patch(id, &category.request)
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
		err := category.service.Delete(id)
		if err != nil {
			lib.JSONErrorResponse(c, err.Error(), nil)
			return
		}
		lib.JSONSuccessResponse(c, nil)
		return
	}

	data, err := category.service.FindByID(id)
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
