package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

/**
*
* Hash password of user
**/
func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err == nil {
		user.Password = string(bytes)
	}
	return err
}

/**
*
* Find user by name
**/
func (user *User) FindByName(name string) error {
	var err error
	err = database.SQL.Where("name = ?", name).First(&user).Error
	return err
}

/**
*
* Create user
**/
func (user *User) Create() (err error) {
	err = database.SQL.Create(&user).Error
	return
}

/**
*
* Find user by Email
**/
func (user *User) FindByEmail(email string) (err error) {
	err = database.SQL.Where("email = ?", email).First(&user).Error
	return err
}
