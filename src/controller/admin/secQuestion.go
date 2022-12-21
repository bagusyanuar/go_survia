package admin

import (
	"errors"
	"go-survia/database"
	"go-survia/src/lib"
	"go-survia/src/model"
	"go-survia/src/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SecQuestion struct{}
type secQuestionRequest struct {
	Question string                     `json:"question" validate:"required"`
	Answers  []secQuestionAnswerRequest `json:"answers" validate:"required"`
}

type secQuestionAnswerRequest struct {
	ID     *uuid.UUID `json:"id"`
	Answer string     `json:"answer" validate:"required"`
	Score  uint       `json:"score" validate:"required"`
}

var secQuestionRepository repositories.SecQuestion

func (SecQuestion) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		postNewSecQuestion(c)
		return
	}
	q := c.Query("q")
	data, err := secQuestionRepository.All(q)
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

func (SecQuestion) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := secQuestionRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, lib.Response{
				Code:    http.StatusNotFound,
				Data:    nil,
				Message: err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: err.Error(),
		})
		return
	}

	// if c.Request.Method == "PATCH" {
	// 	patchSecQuestion(c, data)
	// 	return
	// }

	// if c.Request.Method == "DELETE" {
	// 	deleteQuestion(c, data)
	// 	return
	// }
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func postNewSecQuestion(c *gin.Context) {
	var r secQuestionRequest
	err := c.BindJSON(&r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
			Code:    http.StatusBadRequest,
			Message: "invalid data request",
			Data:    err.Error(),
		})
		return
	}
	v := validator.New()
	if e := v.Struct(&r); e != nil {
		messages := lib.ErrorMessageValidation(e)
		c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
			Code:    http.StatusBadRequest,
			Message: "invalid data request",
			Data:    messages,
		})
		return
	}
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	lastSecQuestionIndex := 0

	var lastSecQuestion *model.SecQuestion
	if err := tx.Model(&model.SecQuestion{}).Order("index_of DESC").First(&lastSecQuestion).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: err.Error(),
			})
			return
		}
	}
	lastSecQuestionIndex = int(lastSecQuestion.IndexOf) + 1

	type SecQuestion struct {
		model.SecQuestion
		Answers  []model.SecAnswer
	}
	sqModel := SecQuestion{
		SecQuestion: model.SecQuestion{
			Question: r.Question,
			IndexOf: uint(lastSecQuestionIndex),
		},
	}
	var arrSaModel []model.SecAnswer
	for k, v := range r.Answers {
		saModel := model.SecAnswer{
			Answer:  v.Answer,
			Score:   int(v.Score),
			IndexOf: uint((k + 1)),
		}
		arrSaModel = append(arrSaModel, saModel)
		// if err := tx.Create(&saModel).Error; err != nil {
		// 	tx.Rollback()
		// }
	}
	sqModel.Answers = arrSaModel
	if err := tx.Debug().Create(&sqModel).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: err.Error(),
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Message: "success",
	})
}

func patchSecQuestion(c *gin.Context, id string) {
	var r secQuestionRequest
	c.Bind(&r)
	v := validator.New()
	if e := v.Struct(&r); e != nil {
		messages := lib.ErrorMessageValidation(e)
		c.AbortWithStatusJSON(http.StatusBadRequest, lib.Response{
			Code:    http.StatusBadRequest,
			Message: "invalid data request",
			Data:    messages,
		})
		return
	}

	// err := secQuestionRepository.Patch(id, data)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, lib.Response{
	// 		Code:    http.StatusInternalServerError,
	// 		Data:    nil,
	// 		Message: err.Error(),
	// 	})
	// 	return
	// }
	c.JSON(http.StatusOK, lib.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    nil,
	})
}

func deleteQuestion(c *gin.Context, d *model.SecQuestion) {
	err := secQuestionRepository.Delete(d)
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
	})
}
