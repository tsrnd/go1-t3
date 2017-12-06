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
	session utils.SessionUtil
)

func (c *LoginController) Index() {
	c.Data["url"] = ""
	c.TplName = "auth/login.html"
}

func (this *LoginController) Login() {
	session.GetSession(this.Controller)
	email := this.GetString("email")
	password := this.GetString("password")
	err := auth.Login(email, password)
	if err != nil {
		this.Ctx.Redirect(302, "/login")
	}
	this.Ctx.Redirect(302, "/")
}
