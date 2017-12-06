package services

import (
	model "github.com/goweb3/models"
	"github.com/astaxie/beego/orm"
	"github.com/goweb3/utils"
)
var session utils.Session

type AuthService struct {
}

func (auth *AuthService) Login(email string, password string) (err error) {
	user := model.User{}
	o := orm.NewOrm()
	qs := o.QueryTable(user)
	qs.Filter("email", email).One(&user)
	if (user != model.User{}) && utils.MatchString(user.Password, password) {
		session.SessionStore.Set("user", user)
	}
	return
}