package services

import (
	model "github.com/goweb3/models"
	"github.com/astaxie/beego/orm"
	"github.com/goweb3/utils"
	"errors"
	"github.com/astaxie/beego"
)
 
type AuthService struct {
}

func (auth *AuthService) Login(email string, password string) (err error) {
	user := model.User{}
	o := orm.NewOrm()
	qs := o.QueryTable(user)
	qs.Filter("email", email).One(&user)
	flash := beego.NewFlash()
	if (user != model.User{}) && utils.MatchString(user.Password, password) {
		sess := utils.Controller.StartSession()
		sess.Set("user", user)
		flash.Notice("Login success!")
		flash.Store(utils.Controller)
		return nil
	}
	flash.Error("Login fail!")
	flash.Store(utils.Controller)
	return errors.New("")
}