package admin

import (
	"go-survia/src/lib"
	repository "go-survia/src/repositories/admin"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Bank struct{}

var bankRepository repository.Bank

func (Bank) Index(c *gin.Context)  {
	q := c.Query("q")
	data, err := bankRepository.All(q)
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