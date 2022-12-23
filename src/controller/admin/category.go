package admin

import (
	"errors"
	"go-survia/src/lib"
	"go-survia/src/repositories"
	request "go-survia/src/request/admin"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Category struct{}

var categoryRepository repositories.Category

func (category Category) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		category.post(c)
		return
	}
	q := c.Query("q")
	data, err := categoryRepository.All(q)
	if err != nil {
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, data)
}

func (category Category) FindByID(c *gin.Context) {
	id := c.Param("id")

	//update method
	if c.Request.Method == "PATCH" {
		category.patch(c, id)
		return
	}

	//delete method
	if c.Request.Method == "DELETE" {
		category.delete(c, id)
		return
	}

	data, err := categoryRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			lib.JSONNotFoundResponse(c, err.Error(), nil)
			return
		}
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func (Category) post(c *gin.Context) {
	var r request.AdminCategoryRequest
	c.Bind(&r)
	m, e := lib.ValidateRequest(&r)
	if e != nil {
		lib.JSONBadRequestResponse(c, "invalid data request", m)
		return
	}
	_, err := categoryRepository.Create(&r)
	if err != nil {
		lib.JSONErrorResponse(c, "internal server error", nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}

func (Category) patch(c *gin.Context, id string) {
	var r request.AdminCategoryRequest
	c.Bind(&r)
	m, e := lib.ValidateRequest(&r)
	if e != nil {
		lib.JSONBadRequestResponse(c, "invalid data request", m)
		return
	}
	data := map[string]interface{}{
		"name": r.Name,
	}
	err := categoryRepository.Patch(id, data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			lib.JSONNotFoundResponse(c, err.Error(), nil)
			return
		}
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}

func (Category) delete(c *gin.Context, id string) {
	err := categoryRepository.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			lib.JSONNotFoundResponse(c, err.Error(), nil)
			return
		}
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}
