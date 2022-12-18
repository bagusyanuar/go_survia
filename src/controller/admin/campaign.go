package admin

import (
	"errors"
	"go-survia/src/lib"
	"go-survia/src/model"
	"go-survia/src/repositories"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Campaign struct{}

var campaignRepository repositories.Campaign

type campaignRequest struct {
	Title            string `form:"title" validate:"required"`
	Description      string `form:"description" validate:"required"`
	ShortDescription string `form:"short_description" validate:"required"`
	StartAt          string `form:"start_at"`
	FinishAt         string `form:"finish_at"`
	Background       string `form:"background" validate:"required"`
}

// var listCampaign []model.Campaign
// var campaign *model.Campaign

func (Campaign) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		postNewCampaign(c)
		return
	}
	q := c.Query("q")
	data, err := campaignRepository.All(q)
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

func (campaign *Campaign) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := campaignRepository.FindByID(id)
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

	//patch method
	if c.Request.Method == "PATCH" {
		var r campaignRequest
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

		var startAt *datatypes.Date
		var finishAt *datatypes.Date
		var image *string

		if r.StartAt != "" {
			tmp, e := time.Parse("2006-01-02", r.StartAt)
			if e != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
					Code:    http.StatusBadRequest,
					Message: "start_at value must be date format",
					Data:    nil,
				})
				return
			}
			startAt = (*datatypes.Date)(&tmp)
		}

		if r.FinishAt != "" {
			tmp, e := time.Parse("2006-01-02", r.FinishAt)
			if e != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
					Code:    http.StatusBadRequest,
					Message: "finish_at value must be date format",
					Data:    nil,
				})
				return
			}
			finishAt = (*datatypes.Date)(&tmp)
		}

		file, _ := c.FormFile("image")

		if file != nil {
			//remove old image
			var oldImage *string = data.Image
			if oldImage != nil {
				_, err := os.Stat(*oldImage)
				if err == nil {
					os.Remove(*oldImage)
					if err != nil {
						c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
							Code:    http.StatusInternalServerError,
							Data:    nil,
							Message: err.Error(),
						})
						return
					}
				}
			}
			//upload new image
			ext := filepath.Ext(file.Filename)
			fileName := "assets/campaigns/" + uuid.New().String() + ext
			image = &fileName
			if errorUpload := c.SaveUploadedFile(file, fileName); errorUpload != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
					Code:    http.StatusInternalServerError,
					Data:    nil,
					Message: errorUpload.Error(),
				})
				return
			}
		}

		patchData := map[string]interface{}{
			"title":             r.Title,
			"description":       r.Description,
			"short_description": r.ShortDescription,
			"image":             image,
			"start_at":          startAt,
			"finish_at":         finishAt,
			"background":        r.Background,
		}
		patchResult, err := campaignRepository.Patch(data, patchData)
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
			Data:    patchResult,
		})
		return
	}

	//delete method
	if c.Request.Method == "DELETE" {
		//remove image
		var oldImage *string = data.Image
		if oldImage != nil {
			_, err := os.Stat(*oldImage)
			if err == nil {
				os.Remove(*oldImage)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
						Code:    http.StatusInternalServerError,
						Data:    nil,
						Message: err.Error(),
					})
					return
				}
			}
		}

		err := campaignRepository.Delete(data)
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

func postNewCampaign(c *gin.Context) {
	var r campaignRequest
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

	var startAt *datatypes.Date
	var finishAt *datatypes.Date
	var image *string

	if r.StartAt != "" {
		tmp, e := time.Parse("2006-01-02", r.StartAt)
		if e != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
				Code:    http.StatusBadRequest,
				Message: "start_at value must be date format",
				Data:    nil,
			})
			return
		}
		startAt = (*datatypes.Date)(&tmp)
	}

	if r.FinishAt != "" {
		tmp, e := time.Parse("2006-01-02", r.FinishAt)
		if e != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
				Code:    http.StatusBadRequest,
				Message: "finish_at value must be date format",
				Data:    nil,
			})
			return
		}
		finishAt = (*datatypes.Date)(&tmp)
	}

	file, _ := c.FormFile("image")

	if file != nil {
		ext := filepath.Ext(file.Filename)
		fileName := "assets/campaigns/" + uuid.New().String() + ext
		image = &fileName
		if errorUpload := c.SaveUploadedFile(file, fileName); errorUpload != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: errorUpload.Error(),
			})
			return
		}
	}

	model := model.Campaign{
		Title:            r.Title,
		Description:      r.Description,
		ShortDescription: r.ShortDescription,
		StartAt:          startAt,
		FinishAt:         finishAt,
		Background:       r.Background,
		Image:            image,
	}
	data, err := campaignRepository.Create(&model)
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
