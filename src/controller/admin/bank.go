package admin

import (
	"errors"
	"go-survia/src/lib"
	"go-survia/src/model"
	"go-survia/src/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Bank struct{}

var bankRepository repositories.Bank

type bankRequest struct {
	Name string `form:"name" validate:"required" json:"name"`
	Code int    `form:"code" validate:"required,numeric" json:"code"`
}

func (Bank) Index(c *gin.Context) {

	if c.Request.Method == "POST" {
		var r bankRequest
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

		m := model.Bank{
			Code: r.Code,
			Name: r.Name,
		}

		d, e := bankRepository.Create(&m)
		if e != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: e.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, lib.Response{
			Code:    http.StatusOK,
			Message: "success",
			Data:    d,
		})
		return
	}
	q := c.Query("q")
	data, err := bankRepository.All(q)
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

func (Bank) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := bankRepository.FindByID(id)
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
		var r bankRequest
		c.Bind(&r)
		v := validator.New()
		if err := v.Struct(&r); err != nil {
			messages := lib.ErrorMessageValidation(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
				Code:    http.StatusBadRequest,
				Message: "invalid data request",
				Data:    messages,
			})
			return
		}

		patchData := map[string]interface{}{
			"name": r.Name,
			"code": r.Code,
		}
		result, err := bankRepository.Patch(data, patchData)
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
			Data:    result,
		})
		return
	}

	//delete method
	if c.Request.Method == "DELETE" {
		err := bankRepository.Delete(data)
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
