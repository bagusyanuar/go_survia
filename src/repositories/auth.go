package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
)

type AuthAdmin struct {
	model.User
	Admin *model.Admin `gorm:"foreignKey:UserID" json:"admin"`
}

func (auth *AuthAdmin) SignIn() (admin *AuthAdmin, err error) {
	username := auth.User.Username
	if err = database.DB.Debug().
		Preload("Admin").
		Joins("JOIN admins ON users.id = admins.user_id").
		Where("username = ? ", username).First(&auth).Error; err != nil {
		return nil, err
	}
	return auth, nil
}
