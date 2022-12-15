package admin

import (
	"errors"
	"go-survia/database"
	"go-survia/src/lib"
	"go-survia/src/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Category struct{}

type apiResponse struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

var listCategory []apiResponse
var category *apiResponse

func (Category) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		name := c.PostForm("name")
		request := model.Category{
			Name: name,
		}
		err := create(request)
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
	results, err := findAll(q)
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

func (Category) FindByID(c *gin.Context) {
	id := c.Param("id")
	if c.Request.Method == "POST" {
		name := c.PostForm("name")
		data := map[string]interface{}{
			"name": name,
		}
		err := patch(id, data)
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

	if c.Request.Method == "DELETE" {
		err := delete(id)
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
	result, err := findById(id)
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
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    result,
	})
}

//logical
func findAll(q string) (r []apiResponse, err error) {
	//unscoped for show deleted item
	if err = database.DB.Unscoped().Model(&model.Category{}).Where("name LIKE ?", "%"+q+"%").Find(&listCategory).Error; err != nil {
		return listCategory, err
	}
	return listCategory, nil
}

func create(d model.Category) (err error) {
	if err = database.DB.Create(&d).Error; err != nil {
		return err
	}
	return nil
}

func findById(id string) (r *apiResponse, err error) {
	if err = database.DB.Model(&model.Category{}).First(&category, "id = ?", id).Error; err != nil {
		return category, err
	}
	return category, nil
}

func patch(id string, data interface{}) (err error) {
	if err = database.DB.Model(&model.Category{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func delete(id string) (err error) {
	if err = database.DB.Where("id = ?", id).Delete(&model.Category{}).Error; err != nil {
		return err
	}
	return nil
}
