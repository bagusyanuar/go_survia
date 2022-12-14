package member

import (
	"go-survia/database"
	"go-survia/src/lib"
	"go-survia/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type Category struct{}

type response struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

var data []response

func (Category) Index(c *gin.Context) {
	results, err := findAll()
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

//repositories
func findAll() (r []response, err error) {
	if err = database.DB.Model(&model.Category{}).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
