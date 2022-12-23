package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomModel interface {
	FindAll(tx *gorm.DB) (*[]interface{}, error)
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONSuccessResponse(c *gin.Context, d interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    d,
	})
}

func JSONErrorResponse(c *gin.Context, m string, d interface{}) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
		Code:    http.StatusInternalServerError,
		Data:    d,
		Message: m,
	})
}

func JSONBadRequestResponse(c *gin.Context, m string, d interface{}) {
	c.AbortWithStatusJSON(http.StatusBadRequest, Response{
		Code:    http.StatusBadRequest,
		Data:    d,
		Message: m,
	})
}
func JSONNotFoundResponse(c *gin.Context, m string, d interface{}) {
	c.AbortWithStatusJSON(http.StatusNotFound, Response{
		Code:    http.StatusNotFound,
		Data:    d,
		Message: m,
	})
}
