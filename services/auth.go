package services

import (
	"fmt"
	model "github.com/goweb3/models"
	"github.com/astaxie/beego/orm"
	"github.com/goweb3/utils"
	"errors"
)

type AuthService struct {
}

func (auth *AuthService) Login(email string, password string) (err error) {
	user := model.User{}
	o := orm.NewOrm()
	qs := o.QueryTable(user)
	qs.Filter("email", email).One(&user)
	if (user != model.User{}) && utils.MatchString(user.Password, password) {
		utils.Session.Set("user", user)
		fmt.Println(utils.Session.Get("user"))
		return nil
	}
	return errors.New("")
}