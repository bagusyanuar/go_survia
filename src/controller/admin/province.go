package admin

import (
	"errors"
	"go-survia/src/lib"
	"go-survia/src/model"
	"go-survia/src/repositories"
	request "go-survia/src/request/admin"
	"go-survia/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Province struct {
	request request.AdminProvince
	service service.Province
}
type provinceRequest struct {
	Code int    `form:"code" validate:"required"`
	Name string `form:"name" validate:"required"`
}

var provinceRepository repositories.Province

func (province *Province) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		// postNewProvince(c)
		c.Bind(&province.request)

		if m, e := province.service.ValidateRequest(&province.request); e != nil {
			lib.JSONBadRequestResponse(c, e.Error(), m)
			return
		}

		if e := province.service.Create(&province.request); e != nil {
			lib.JSONErrorResponse(c, e.Error(), nil)
		}
		lib.JSONSuccessResponse(c, nil)
		return
	}
	q := c.Query("q")
	data, err := provinceRepository.All(q)
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

func (Province) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := provinceRepository.FindByID(id)
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
		patchProvince(c, data)
		return
	}

	if c.Request.Method == "DELETE" {
		deleteProvince(c, data)
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func postNewProvince(c *gin.Context) {
	// var r provinceRequest
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

	// model := model.Province{
	// 	Code: r.Code,
	// 	Name: r.Name,
	// }
	// data, err := provinceRepository.Create(&model)
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

func patchProvince(c *gin.Context, d *model.Province) {
	var r provinceRequest
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
	data := map[string]interface{}{
		"code": r.Code,
		"name": r.Name,
	}
	result, err := provinceRepository.Patch(d, data)
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

func deleteProvince(c *gin.Context, d *model.Province) {
	err := provinceRepository.Delete(d)
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
