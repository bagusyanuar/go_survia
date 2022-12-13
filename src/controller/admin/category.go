package admin

import (
	"go-survia/database"
	"go-survia/src/lib"
	"go-survia/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type Category struct{}

type CategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (Category) Index(c *gin.Context) {
	var response []CategoryResponse
	if err := database.DB.Debug().Model(&model.Category{}).Find(&response).Error; err != nil {
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
		Data:    response,
	})
}
