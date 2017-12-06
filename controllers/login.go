package controllers

import (
	"github.com/astaxie/beego"
	"github.com/goweb3/services"
	"github.com/goweb3/utils"
)
type LoginController struct {
	beego.Controller
}
var (
	auth services.AuthService
	session utils.Session
)

func (c *LoginController) Get() {
	c.Data["url"] = ""
	c.TplName = "auth/login.html"
}

func (this *LoginController) Post() {
	session.GetSession(this)
	email := this.GetString("email")
	password := this.GetString("password")
	auth.Login(email, password)
}