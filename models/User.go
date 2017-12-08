package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/goweb3/app/shared/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint
	Name     string
	Email    string
	Password string
	Cart     []*Cart `orm:"reverse(many)"`
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

/**
*
* Find user by ID
**/
func (user *User) FindByID() (err error) {
	o := orm.NewOrm()
	err = o.Read(user)
	return
}
