package admin

import (
	"go-survia/database"
	"go-survia/src/lib"
	"go-survia/src/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
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
		finishAt,_ := time.Parse("2006-01-02", c.PostForm("finish_at"))
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
			c.JSON(http.StatusOK, lib.Response{
				Code:    http.StatusOK,
				Message: file.Header.Get("Content-Type"),
			})
			return
		}
		request := model.Campaign{
			Title:            title,
			Description:      description,
			ShortDescription: shortDescription,
			StartAt:          startValue,
			FinishAt:         finishValue,
			Background:       background,
			Image: imageValue,
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
