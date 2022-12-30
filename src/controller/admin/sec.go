package admin

import (
	"errors"
	"go-survia/src/lib"
	"go-survia/src/model"
	"go-survia/src/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Sec struct{}
type secRequest struct {
	Name   string `form:"name" validate:"required"`
	Bottom int    `form:"bottom" validate:"required"`
	Top    int    `form:"top" validate:"required"`
}

var secRepository repositories.Sec

func (Sec) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		postNewSec(c)
		return
	}
	q := c.Query("q")
	data, err := secRepository.All(q)
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

func (Sec) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := secRepository.FindByID(id)
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
		patchSec(c, data)
		return
	}

	if c.Request.Method == "DELETE" {
		deleteSec(c, data)
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func postNewSec(c *gin.Context) {
	// var r secRequest
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

	// model := model.Sec{
	// 	Name: r.Name,
	// 	Bottom: r.Bottom,
	// 	Top: r.Top,
	// }
	// data, err := secRepository.Create(&model)
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

func patchSec(c *gin.Context, d *model.Sec) {
	// var r secRequest
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
	// data := map[string]interface{}{
	// 	"name":   r.Name,
	// 	"bottom": r.Bottom,
	// 	"top":    r.Top,
	// }
	// result, err := secRepository.Patch(d, data)
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
	// 	Data:    result,
	// })
}

func deleteSec(c *gin.Context, d *model.Sec) {
	// err := secRepository.Delete(d)
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
	// })
}
