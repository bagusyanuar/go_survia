package admin

import (
	"errors"
	"go-survia/src/lib"
	"go-survia/src/repositories"
	request "go-survia/src/request/admin"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func (Category) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := categoryRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, lib.Response{
				Code:    http.StatusNotFound,
				Data:    nil,
				Message: err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: err.Error(),
		})
		return
	}

	//update method
	if c.Request.Method == "PATCH" {
		// var r categoryRequest
		// c.Bind(&r)
		// v := validator.New()
		// if e := v.Struct(&r); e != nil {
		// 	messages := lib.ErrorMessageValidation(e)
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
		// 		Code:    http.StatusBadRequest,
		// 		Message: "invalid data request",
		// 		Data:    messages,
		// 	})
		// 	return
		// }
		// patchData := map[string]interface{}{
		// 	"name": r.Name,
		// }

		// patchResult, err := categoryRepository.Patch(data, patchData)
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
		// 		Code:    http.StatusInternalServerError,
		// 		Data:    nil,
		// 		Message: err.Error(),
		// 	})
		// 	return
		// }
		// c.JSON(http.StatusOK, lib.Response{
		// 	Code:    http.StatusOK,
		// 	Message: "success",
		// 	Data:    patchResult,
		// })
		// return
	}

	if c.Request.Method == "DELETE" {
		err := categoryRepository.Delete(data)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, lib.Response{
			Code:    http.StatusOK,
			Message: "success",
		})
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
	v := validator.New()
	if e := v.Struct(&r); e != nil {
		messages := lib.ErrorMessageValidation(e)
		c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
			Code:    http.StatusBadRequest,
			Message: "invalid data request",
			Data:    messages,
		})
		return
	}
	_, err := categoryRepository.Create(&r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Message: "success",
	})
}
