package services

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	model "github.com/goweb3/models"
	"github.com/goweb3/utils"
)

type AuthService struct {
}

func (auth *AuthService) Login(email string, password string) (err error) {
	user := model.User{}
	o := orm.NewOrm()
	qs := o.QueryTable(user)
	err = qs.Filter("email", email).One(&user)
	flash := beego.NewFlash()
	if err == nil && utils.MatchString(user.Password, password) {
		sess := utils.Controller.StartSession()
		sess.Set("auth", user)
		flash.Notice("Login success!")
		flash.Store(utils.Controller)
		return nil
	}
	flash.Error("Login fail!")
	flash.Store(utils.Controller)
	return errors.New("")
}
