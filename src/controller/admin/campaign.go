package admin

import (
	"errors"
	"go-survia/src/lib"
	request "go-survia/src/request/admin"
	"go-survia/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Campaign struct {
	service service.Campaign
	request request.AdminCampaign
}

func (campaign *Campaign) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		c.Bind(&campaign.request)
		messages, file, fileName, err := campaign.service.Create(c, &campaign.request)
		if err != nil {
			if errors.Is(err, lib.ErrBadRequest) {
				lib.JSONBadRequestResponse(c, err.Error(), messages)
				return
			}
			lib.JSONErrorResponse(c, err.Error(), nil)
			return
		}

		if file != nil {
			if errorUpload := c.SaveUploadedFile(file, *fileName); errorUpload != nil {
				lib.JSONErrorResponse(c, errorUpload.Error(), nil)
				return
			}
		}
		lib.JSONSuccessResponse(c, nil)
		return
	}
	q := c.Query("q")
	data, err := campaign.service.FindAll(q)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: err.Error(),
		})
		return
	}
	lib.JSONSuccessResponse(c, data)
}

func (campaign *Campaign) FindByID(c *gin.Context) {
	id := c.Param("id")

	if c.Request.Method == "PATCH" {
		c.Bind(&campaign.request)
		messages, file, fileName, err := campaign.service.Patch(c, id, &campaign.request)
		if err != nil {
			if errors.Is(err, lib.ErrBadRequest) {
				lib.JSONBadRequestResponse(c, err.Error(), messages)
				return
			}
			lib.JSONErrorResponse(c, err.Error(), nil)
			return
		}

		if file != nil {
			if errorUpload := c.SaveUploadedFile(file, *fileName); errorUpload != nil {
				lib.JSONErrorResponse(c, errorUpload.Error(), nil)
				return
			}
		}
		lib.JSONSuccessResponse(c, nil)
		return
	}

	//delete method
	if c.Request.Method == "DELETE" {
		err := campaign.service.Delete(id)
		if err != nil {
			lib.JSONErrorResponse(c, err.Error(), nil)
			return
		}
		lib.JSONSuccessResponse(c, nil)
		return
	}

	data, err := campaign.service.FindByID(id)
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
	lib.JSONSuccessResponse(c, data)
}
