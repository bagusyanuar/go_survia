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

type Bank struct{}

var bankRepository repositories.Bank

func (bank Bank) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		bank.post(c)
		return
	}
	q := c.Query("q")
	data, err := bankRepository.All(q)
	if err != nil {
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, data)
}

func (bank Bank) FindByID(c *gin.Context) {
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
		bank.patch(c, id)
		return
	}

	//delete method
	if c.Request.Method == "DELETE" {
		bank.delete(c, id)
		return
	}
	
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func (Bank) post(c *gin.Context) {
	var r request.AdminBankRequest
	c.Bind(&r)
	m, e := lib.ValidateRequest(&r)
	if e != nil {
		lib.JSONBadRequestResponse(c, "invalid data request", m)
		return
	}

	_, err := bankRepository.Create(&r)
	if err != nil {
		lib.JSONErrorResponse(c, "internal server error", nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}

func (Bank) patch(c *gin.Context, id string)  {
	var r request.AdminBankRequest
	c.Bind(&r)
	m, e := lib.ValidateRequest(&r)
	if e != nil {
		lib.JSONBadRequestResponse(c, "invalid data request", m)
		return
	}
	data := map[string]interface{}{
		"name": r.Name,
		"code": r.Code,
	}
	err := bankRepository.Patch(id, data)
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

func (Bank) delete(c *gin.Context, id string)  {
	err := bankRepository.Delete(id)
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
