package auth

import (
	"go-survia/src/lib"
	"go-survia/src/model"
	"go-survia/src/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type AuthAdmin struct{}

func (AuthAdmin) SignIn(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	repository := repositories.AuthAdmin{
		User: model.User{
			Username: username,
			Password: &password,
		},
	}

	user, err := repository.SignIn()
	if err != nil {
		errorResponse := lib.ErrorSignIn(err)
		c.AbortWithStatusJSON(errorResponse.Code, errorResponse)
		return
	}

	jwt := lib.JWT{}
	claim := lib.JWTClaims{
		Unique: uuid.UUID(user.ID),
		Email:  *user.Email,
		Role:   "admin",
	}

	accessToken, errorTokenize := jwt.GenerateToked(claim)
	if errorTokenize != nil {
		c.AbortWithStatusJSON(500, lib.Response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "error tokenize " + errorTokenize.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lib.Response{
		Code: http.StatusOK,
		Data: map[string]interface{}{
			"accessToken": accessToken,
			"user":        user,
		},
		Message: "success sign in",
	})
}
