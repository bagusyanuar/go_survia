package auth

import (
	"go-survia/src/model"
	"go-survia/src/repositories"

	"github.com/gin-gonic/gin"
)


type AuthAdmin struct{}

func (AuthAdmin) SignIn(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	
	repository := repositories.AuthAdmin{
		User:  model.User{
			Username: username,
			Password: &password,
		},
	}

	user, err := repository.SignIn()
	
}