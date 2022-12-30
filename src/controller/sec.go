package controller

import (
	"errors"
	"go-survia/src/lib"
	req "go-survia/src/request"
	"go-survia/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Sec struct {
	service service.Sec
	request req.Sec
}

func (sec *Sec) Index(c *gin.Context) {
	q := c.Query("q")
	data, err := sec.service.FindAll(q)
	if err != nil {
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, data)
}

func (sec *Sec) Store(c *gin.Context) {
	c.Bind(&sec.request)
	if m, e := lib.ValidateRequest(&sec.request); e != nil {
		lib.JSONBadRequestResponse(c, e.Error(), m)
		return
	}

	if e := sec.service.Create(&sec.request); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}

func (sec *Sec) Show(c *gin.Context) {
	id := c.Param("id")
	data, e := sec.service.FindByID(id)
	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			lib.JSONNotFoundResponse(c, e.Error(), nil)
			return
		}
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, data)
}

func (sec *Sec) Update(c *gin.Context) {
	id := c.Param("id")
	c.Bind(&sec.request)
	if m, e := lib.ValidateRequest(&sec.request); e != nil {
		lib.JSONBadRequestResponse(c, e.Error(), m)
		return
	}
	if e := sec.service.Patch(id, &sec.request); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}

func (sec *Sec) Destroy(c *gin.Context) {
	id := c.Param("id")
	if e := sec.service.Delete(id); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}
