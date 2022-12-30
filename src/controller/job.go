package controller

import (
	"errors"
	"go-survia/src/lib"
	req "go-survia/src/request"
	"go-survia/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Job struct {
	service service.Job
	request req.Job
}

func (job *Job) Index(c *gin.Context) {
	q := c.Query("q")
	data, err := job.service.FindAll(q)
	if err != nil {
		lib.JSONErrorResponse(c, err.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, data)
}

func (job *Job) Store(c *gin.Context) {
	c.Bind(&job.request)
	if m, e := lib.ValidateRequest(&job.request); e != nil {
		lib.JSONBadRequestResponse(c, e.Error(), m)
		return
	}

	if e := job.service.Create(&job.request); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}

func (job *Job) Show(c *gin.Context) {
	id := c.Param("id")
	data, e := job.service.FindByID(id)
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

func (job *Job) Update(c *gin.Context) {
	id := c.Param("id")
	c.Bind(&job.request)
	if m, e := lib.ValidateRequest(&job.request); e != nil {
		lib.JSONBadRequestResponse(c, e.Error(), m)
		return
	}
	if e := job.service.Patch(id, &job.request); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}

func (job *Job) Destroy(c *gin.Context) {
	id := c.Param("id")
	if e := job.service.Delete(id); e != nil {
		lib.JSONErrorResponse(c, e.Error(), nil)
		return
	}
	lib.JSONSuccessResponse(c, nil)
}
