package admin

import (
	"errors"
	"go-survia/src/lib"
	"go-survia/src/model"
	"go-survia/src/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type City struct{}
type cityRequest struct {
	ProvinceID string `form:"province_id" validate:"required"`
	Code       int    `form:"code" validate:"required"`
	Name       string `form:"name" validate:"required"`
}

var cityRepository repositories.City

func (City) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		postNewCity(c)
		return
	}
	q := c.Query("q")
	data, err := cityRepository.All(q)
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

func (City) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := cityRepository.FindByID(id)
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
	if c.Request.Method == "PATCH" {
		patchCity(c, data)
		return
	}

	if c.Request.Method == "DELETE" {
		deleteCity(c, data)
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}
func postNewCity(c *gin.Context) {
	// var r cityRequest
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

	// provinceId, e := uuid.Parse(r.ProvinceID)
	// if e != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
	// 		Code:    http.StatusBadRequest,
	// 		Message: "invalid province id format",
	// 	})
	// 	return
	// }
	// model := model.City{
	// 	ProvinceID: provinceId,
	// 	Code:       r.Code,
	// 	Name:       r.Name,
	// }
	// data, err := cityRepository.Create(&model)
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
	// 	Data:    data,
	// })
}

func patchCity(c *gin.Context, d *model.City) {
	var r cityRequest
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
	provinceId, e := uuid.Parse(r.ProvinceID)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
			Code:    http.StatusBadRequest,
			Message: "invalid province id format",
		})
		return
	}
	data := map[string]interface{}{
		"province_id": provinceId,
		"code":        r.Code,
		"name":        r.Name,
	}
	result, err := cityRepository.Patch(d, data)
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
}

func deleteCity(c *gin.Context, d *model.City) {
	err := cityRepository.Delete(d)
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
