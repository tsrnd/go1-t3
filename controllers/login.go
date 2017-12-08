package controllers

import (
	"html/template"

	"github.com/astaxie/beego"
	"github.com/goweb3/services"
	"github.com/goweb3/utils"
)

type LoginController struct {
	beego.Controller
}

var (
	auth    services.AuthService
	session utils.SessionUtil
)

func (c *LoginController) Index() {
	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["error"]; ok {
		c.Data["error"] = n
	}
	if n, ok := flash.Data["notice"]; ok {
		c.Data["notice"] = n
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["url"] = ""
	c.TplName = "auth/login.html"
}

func (this *LoginController) Login() {
	session.SetControler(this.Controller)
	email := this.GetString("email")
	password := this.GetString("password")
	err := auth.Login(email, password)
	if err != nil {
		this.Ctx.Redirect(302, "/login")
	}
	this.Ctx.Redirect(302, "/")
}
