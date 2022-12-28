package controller

import (
	"go-survia/src/lib"
	"go-survia/src/service"

	"github.com/gin-gonic/gin"
)

type Category struct {
	service service.Category
}

func (category *Category) GetData(c *gin.Context) {
	q := c.Query("q")
	data, err := category.service.FindAll(q)
	if err != nil {
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, data)
}

func (category *Category) StoreData(c *gin.Context) {

}
