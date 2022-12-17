package admin

import (
	"errors"
	"go-survia/database"
	"go-survia/src/lib"
	"go-survia/src/model"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Campaign struct{}

var listCampaign []model.Campaign
var campaign *model.Campaign

func (Campaign) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		title := c.PostForm("title")
		description := c.PostForm("description")
		shortDescription := c.PostForm("short_description")
		startAt, _ := time.Parse("2006-01-02", c.PostForm("start_at"))
		finishAt, _ := time.Parse("2006-01-02", c.PostForm("finish_at"))
		background := c.PostForm("background")

		var startValue *datatypes.Date
		var finishValue *datatypes.Date
		var imageValue *string
		// finishValue := datatypes.Date(finishAt)

		if c.PostForm("start_at") != "" {
			startValue = (*datatypes.Date)(&startAt)
		}
		if c.PostForm("finish_at") != "" {
			finishValue = (*datatypes.Date)(&finishAt)
		}

		file, _ := c.FormFile("image")
		if file != nil {
			ext := filepath.Ext(file.Filename)
			fileName := "assets/campaigns/" + uuid.New().String() + ext
			imageValue = &fileName
			if errorUpload := c.SaveUploadedFile(file, fileName); errorUpload != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
					Code:    http.StatusInternalServerError,
					Data:    nil,
					Message: errorUpload.Error(),
				})
				return
			}
		}
		request := model.Campaign{
			Title:            title,
			Description:      description,
			ShortDescription: shortDescription,
			StartAt:          startValue,
			FinishAt:         finishValue,
			Background:       background,
			Image:            imageValue,
		}
		err := createCampaign(request)
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
	q := c.Query("q")
	results, err := findAllCampaign(q)
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
		Data:    results,
	})
}

func (campaign *Campaign) FindByID(c *gin.Context) {
	id := c.Param("id")
	result, err := campaign.findById(id)
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
	if c.Request.Method == "POST" {
		title := c.PostForm("title")
		description := c.PostForm("description")
		shortDescription := c.PostForm("short_description")
		startAt, _ := time.Parse("2006-01-02", c.PostForm("start_at"))
		finishAt, _ := time.Parse("2006-01-02", c.PostForm("finish_at"))
		background := c.PostForm("background")

		var startValue *datatypes.Date
		var finishValue *datatypes.Date
		var imageValue *string = result.Image
		// finishValue := datatypes.Date(finishAt)

		if c.PostForm("start_at") != "" {
			startValue = (*datatypes.Date)(&startAt)
		}
		if c.PostForm("finish_at") != "" {
			finishValue = (*datatypes.Date)(&finishAt)
		}

		file, _ := c.FormFile("image")
		if file != nil {
			ext := filepath.Ext(file.Filename)
			fileName := "assets/campaigns/" + uuid.New().String() + ext
			imageValue = &fileName
			if errorUpload := c.SaveUploadedFile(file, fileName); errorUpload != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
					Code:    http.StatusInternalServerError,
					Data:    nil,
					Message: errorUpload.Error(),
				})
				return
			}
		}

		data := map[string]interface{}{
			"title":             title,
			"description":       description,
			"short_description": shortDescription,
			"image":             imageValue,
			"start_at":          startValue,
			"finish_at":         finishValue,
			"background":        background,
		}
		r, e := campaign.patch(result, data)
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
			Data:    r,
		})
		return
	}

	//delete method
	if c.Request.Method == "DELETE" {
		var imageValue *string = result.Image
		if imageValue != nil {
			_, err := os.Stat(*imageValue)
			if err == nil {
				os.Remove(*imageValue)
				// if err != nil {
				// 	c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
				// 		Code:    http.StatusInternalServerError,
				// 		Data:    nil,
				// 		Message: err.Error(),
				// 	})
				// 	return
				// }
			}
		}

		err := campaign.delete(result)
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
		Data:    result,
	})
}

//logical
func findAllCampaign(q string) (r []model.Campaign, err error) {
	//unscoped for show deleted item
	if err = database.DB.Unscoped().Model(&model.Campaign{}).Where("title LIKE ?", "%"+q+"%").Or("description LIKE ?", "%"+q+"%").Find(&listCampaign).Error; err != nil {
		return listCampaign, err
	}
	return listCampaign, nil
}

func createCampaign(d model.Campaign) (err error) {
	if err = database.DB.Create(&d).Error; err != nil {
		return err
	}
	return nil
}

func (Campaign) findById(id string) (r *model.Campaign, err error) {
	if err = database.DB.Model(&model.Campaign{}).First(&campaign, "id = ?", id).Error; err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (Campaign) patch(m *model.Campaign, data interface{}) (r *model.Campaign, err error) {
	if err = database.DB.Model(&m).Updates(data).Error; err != nil {
		return m, err
	}
	return m, nil
}
func (Campaign) delete(m *model.Campaign) (err error) {
	if err = database.DB.Delete(&m).Error; err != nil {
		return err
	}
	return nil
}
