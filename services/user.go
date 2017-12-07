package services

import (
	"github.com/astaxie/beego/orm"
	model "github.com/goweb3/models"

)
type UserService struct{

}

func (service *UserService) Register(data model.User) (err error) {
	o := orm.NewOrm()
	user := new(model.User)
	user.SetUser(data)
	user.HashPassword()
	_, err = o.Insert(user)
	return
}