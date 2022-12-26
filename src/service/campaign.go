package service

import (
	"go-survia/src/lib"
	"go-survia/src/model"
	"go-survia/src/repositories"
	adminRequest "go-survia/src/request/admin"
	adminResponse "go-survia/src/response/admin"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Campaign struct {
	repository repositories.Campaign
}

func (campaign *Campaign) Create(c *gin.Context, request *adminRequest.AdminCampaign) (d interface{}, f *multipart.FileHeader, fname *string, err error) {
	var startAt *datatypes.Date
	var finishAt *datatypes.Date
	var image *string

	messages, e := lib.ValidateRequest(request)
	if e != nil {
		return messages, nil, image, lib.ErrBadRequest
	}

	if request.StartAt != "" {
		tmp, e := time.Parse("2006-01-02", request.StartAt)
		if e != nil {
			msg := map[string]interface{}{
				"key":     "start_at",
				"message": "invalid format",
			}
			messages = append(messages, msg)
			return messages, nil, image, lib.ErrBadRequest
		}
		startAt = (*datatypes.Date)(&tmp)
	}

	if request.FinishAt != "" {
		tmp, e := time.Parse("2006-01-02", request.FinishAt)
		if e != nil {
			msg := map[string]interface{}{
				"key":     "finish_at",
				"message": "invalid format",
			}
			messages = append(messages, msg)
			return messages, nil, image, lib.ErrBadRequest
		}
		finishAt = (*datatypes.Date)(&tmp)
	}

	file, _ := c.FormFile("image")

	if file != nil {
		ext := filepath.Ext(file.Filename)
		fileName := "assets/campaigns/" + uuid.New().String() + ext
		image = &fileName
	}

	entity := model.Campaign{
		Title:            request.Title,
		Description:      request.Description,
		ShortDescription: request.ShortDescription,
		StartAt:          startAt,
		FinishAt:         finishAt,
		Background:       request.Background,
		Image:            image,
	}
	e = campaign.repository.Create(&entity)
	if e != nil {
		return nil, file, image, e
	}
	return nil, file, image, nil
}

func (campaign *Campaign) Patch(c *gin.Context, id string, request *adminRequest.AdminCampaign) (d interface{}, f *multipart.FileHeader, fname *string, err error) {
	var startAt *datatypes.Date
	var finishAt *datatypes.Date
	var image *string

	messages, e := lib.ValidateRequest(request)
	if e != nil {
		return messages, nil, image, lib.ErrBadRequest
	}

	if request.StartAt != "" {
		tmp, e := time.Parse("2006-01-02", request.StartAt)
		if e != nil {
			msg := map[string]interface{}{
				"key":     "start_at",
				"message": "invalid format",
			}
			messages = append(messages, msg)
			return messages, nil, image, lib.ErrBadRequest
		}
		startAt = (*datatypes.Date)(&tmp)
	}

	if request.FinishAt != "" {
		tmp, e := time.Parse("2006-01-02", request.FinishAt)
		if e != nil {
			msg := map[string]interface{}{
				"key":     "finish_at",
				"message": "invalid format",
			}
			messages = append(messages, msg)
			return messages, nil, image, lib.ErrBadRequest
		}
		finishAt = (*datatypes.Date)(&tmp)
	}
	data := map[string]interface{}{
		"title":             request.Title,
		"description":       request.Description,
		"short_description": request.ShortDescription,
		"start_at":          startAt,
		"finish_at":         finishAt,
		"background":        request.Background,
	}
	file, _ := c.FormFile("image")

	if file != nil {
		ext := filepath.Ext(file.Filename)
		fileName := "assets/campaigns/" + uuid.New().String() + ext
		image = &fileName
		data["image"] = fileName
	}

	return nil, file, image, campaign.repository.Patch(id, data)
}

func (campaign *Campaign) Delete(id string) error {
	return campaign.repository.Delete(id)
}

func (campaign *Campaign) FindAll(q string) (b []adminResponse.APICampaign, err error) {
	return campaign.repository.All(q)
}

func (campaign *Campaign) FindByID(id string) (r *adminResponse.APICampaign, err error) {
	entity, e := campaign.repository.FindByID(id)
	if e != nil {
		return nil, e
	}
	return entity, nil
}

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
