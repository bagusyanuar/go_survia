package database

import (
	"encoding/json"
	"fmt"
	"go-survia/src/model"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type adminUser struct {
	model.Admin
	User model.User
}

func Seed() {
	email := "administrator@gmail.com"
	username := "administrator"
	password := "administrator"
	name := "Administrator"

	roles, _ := json.Marshal([]string{"admin"})
	hash, errHashing := bcrypt.GenerateFromPassword([]byte(password), 13)
	if errHashing != nil {
		log.Fatal("error hashing")
	}
	password = string(hash)
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user := model.User{
		Username: username,
		Email:    &email,
		Password: &password,
		Roles:    roles,
	}

	admin := adminUser{
		Admin: model.Admin{
			Name:     name,
			IsActive: true,
		},
		User: user,
	}

	if err := tx.Debug().Create(&admin).Error; err != nil {
		log.Fatal("error insert data")
	}
	tx.Commit()
	fmt.Println("success seed admin")
}
