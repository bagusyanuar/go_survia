package user

import (
	"go-survia/database"
	"go-survia/src/model"
)

var users *[]model.User

func Find() (data *[]model.User, err error) {
	if err = database.DB.Debug().Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
