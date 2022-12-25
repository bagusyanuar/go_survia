package member

import (
	"go-survia/src/repositories"

	"github.com/gin-gonic/gin"
)

type Category struct{}

var categoryRepository repositories.Category

func (Category) Index(c *gin.Context) {
	// q := c.Query("q")
	// data, err := categoryRepository.GetCategories(q)
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
