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

type Job struct{}
type jobRequest struct {
	Name string `form:"name" validate:"required"`
}

var jobRepository repositories.Job

func (Job) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		postNewJob(c)
		return
	}
	q := c.Query("q")
	data, err := jobRepository.All(q)
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

func (Job) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := jobRepository.FindByID(id)
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
		patchJob(c, data)
		return
	}

	if c.Request.Method == "DELETE" {
		deleteJob(c, data)
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func postNewJob(c *gin.Context) {
	// var r jobRequest
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

	// model := model.Job{
	// 	Name: r.Name,
	// }
	// data, err := jobRepository.Create(&model)
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

func patchJob(c *gin.Context, d *model.Job) {
	// var r jobRequest
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
	// 	"name": r.Name,
	// }
	// result, err := jobRepository.Patch(d, data)
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

func deleteJob(c *gin.Context, d *model.Job) {
	// err := jobRepository.Delete(d)
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
