package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/astaxie/beego/orm"
	"github.com/goweb3/utils"
)

type User struct {
	Id uint `form:"-"`
	Name string `form:"name"`
	Email string `form:"email"`
	Password string `form:"password"`
}
func init() {
    orm.RegisterModel(new(User))
}
func (p *User) TableName() string {
    return "users"
}
/**
*
* Hash password of user
**/
func (user *User) HashPassword() error {
	pass, err := utils.HashString(user.Password)
	user.Password = pass
	return err
}

func (user *User) SetUser(userTmp User) {
	user.Id = userTmp.Id
	user.Email = userTmp.Email
	user.Name = userTmp.Name
	user.Password = userTmp.Password
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
