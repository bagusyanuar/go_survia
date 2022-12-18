package admin

import (
	"go-survia/src/repositories"

	"github.com/gin-gonic/gin"
)

type City struct{}

var cityRepository repositories.City

func (City) Index(c *gin.Context) {

}
