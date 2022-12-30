package controller

import (
	"go-survia/src/lib"
	req "go-survia/src/request"
	"go-survia/src/service"

	"github.com/gin-gonic/gin"
)

type SecQuestion struct {
	service service.SecQuestion
	request req.SecQuestion
}

func (secQuestion *SecQuestion) Index(c *gin.Context) {
	q := c.Query("q")
	data, err := secQuestion.service.FindAll(q)
	if err != nil {
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, data)
}

func (secQuestion *SecQuestion) Store(c *gin.Context) {
	c.BindJSON(&secQuestion.request)
	if m, e := lib.ValidateRequest(&secQuestion.request); e != nil {
		lib.JSONBadRequestResponse(c, e.Error(), m)
		return
	}

	if e := secQuestion.service.Create(&secQuestion.request); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}