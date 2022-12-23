package admin

import (
	"errors"
	"go-survia/src/lib"
	"go-survia/src/model"
	"go-survia/src/repositories"
	request "go-survia/src/request/admin"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Campaign struct{}

var campaignRepository repositories.Campaign

func (campaign Campaign) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		campaign.post(c)
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
		// var r campaignRequest
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

		// var startAt *datatypes.Date
		// var finishAt *datatypes.Date
		// var image *string

		// if r.StartAt != "" {
		// 	tmp, e := time.Parse("2006-01-02", r.StartAt)
		// 	if e != nil {
		// 		c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
		// 			Code:    http.StatusBadRequest,
		// 			Message: "start_at value must be date format",
		// 			Data:    nil,
		// 		})
		// 		return
		// 	}
		// 	startAt = (*datatypes.Date)(&tmp)
		// }

		// if r.FinishAt != "" {
		// 	tmp, e := time.Parse("2006-01-02", r.FinishAt)
		// 	if e != nil {
		// 		c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
		// 			Code:    http.StatusBadRequest,
		// 			Message: "finish_at value must be date format",
		// 			Data:    nil,
		// 		})
		// 		return
		// 	}
		// 	finishAt = (*datatypes.Date)(&tmp)
		// }

		// file, _ := c.FormFile("image")

		// if file != nil {
		// 	//remove old image
		// 	var oldImage *string = data.Image
		// 	if oldImage != nil {
		// 		_, err := os.Stat(*oldImage)
		// 		if err == nil {
		// 			os.Remove(*oldImage)
		// 			if err != nil {
		// 				c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
		// 					Code:    http.StatusInternalServerError,
		// 					Data:    nil,
		// 					Message: err.Error(),
		// 				})
		// 				return
		// 			}
		// 		}
		// 	}
		// 	//upload new image
		// 	ext := filepath.Ext(file.Filename)
		// 	fileName := "assets/campaigns/" + uuid.New().String() + ext
		// 	image = &fileName
		// 	if errorUpload := c.SaveUploadedFile(file, fileName); errorUpload != nil {
		// 		c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
		// 			Code:    http.StatusInternalServerError,
		// 			Data:    nil,
		// 			Message: errorUpload.Error(),
		// 		})
		// 		return
		// 	}
		// }

		// patchData := map[string]interface{}{
		// 	"title":             r.Title,
		// 	"description":       r.Description,
		// 	"short_description": r.ShortDescription,
		// 	"image":             image,
		// 	"start_at":          startAt,
		// 	"finish_at":         finishAt,
		// 	"background":        r.Background,
		// }
		// patchResult, err := campaignRepository.Patch(data, patchData)
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
		return
	}

	//delete method
	if c.Request.Method == "DELETE" {
		//remove image
		// var oldImage *string = data.Image
		// if oldImage != nil {
		// 	_, err := os.Stat(*oldImage)
		// 	if err == nil {
		// 		os.Remove(*oldImage)
		// 		if err != nil {
		// 			c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
		// 				Code:    http.StatusInternalServerError,
		// 				Data:    nil,
		// 				Message: err.Error(),
		// 			})
		// 			return
		// 		}
		// 	}
		// }

		// err := campaignRepository.Delete(data)
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
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func (Campaign) post(c *gin.Context)  {
	var r request.AdminCampaignRequest
	c.Bind(&r)
	m, e := lib.ValidateRequest(&r)
	if e != nil {
		lib.JSONBadRequestResponse(c, "invalid data request", m)
		return
	}


}

func (Campaign) postData(r *request.AdminCampaignRequest) (m *model.Campaign, err error)  {
	// var startAt *datatypes.Date
	// var finishAt *datatypes.Date
	// var image *string

	// if r.StartAt != "" {
		
	// }


	return nil, nil
}
func postNewCampaign(c *gin.Context) {
	// var r campaignRequest
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

	// var startAt *datatypes.Date
	// var finishAt *datatypes.Date
	// var image *string

	// if r.StartAt != "" {
	// 	tmp, e := time.Parse("2006-01-02", r.StartAt)
	// 	if e != nil {
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
	// 			Code:    http.StatusBadRequest,
	// 			Message: "start_at value must be date format",
	// 			Data:    nil,
	// 		})
	// 		return
	// 	}
	// 	startAt = (*datatypes.Date)(&tmp)
	// }

	// if r.FinishAt != "" {
	// 	tmp, e := time.Parse("2006-01-02", r.FinishAt)
	// 	if e != nil {
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
	// 			Code:    http.StatusBadRequest,
	// 			Message: "finish_at value must be date format",
	// 			Data:    nil,
	// 		})
	// 		return
	// 	}
	// 	finishAt = (*datatypes.Date)(&tmp)
	// }

	// file, _ := c.FormFile("image")

	// if file != nil {
	// 	ext := filepath.Ext(file.Filename)
	// 	fileName := "assets/campaigns/" + uuid.New().String() + ext
	// 	image = &fileName
	// 	if errorUpload := c.SaveUploadedFile(file, fileName); errorUpload != nil {
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
	// 			Code:    http.StatusInternalServerError,
	// 			Data:    nil,
	// 			Message: errorUpload.Error(),
	// 		})
	// 		return
	// 	}
	// }

	// model := model.Campaign{
	// 	Title:            r.Title,
	// 	Description:      r.Description,
	// 	ShortDescription: r.ShortDescription,
	// 	StartAt:          startAt,
	// 	FinishAt:         finishAt,
	// 	Background:       r.Background,
	// 	Image:            image,
	// }
	// data, err := campaignRepository.Create(&model)
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
