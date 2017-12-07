package controllers

import (
	"github.com/astaxie/beego"
	"github.com/goweb3/services"
	"github.com/goweb3/models"
)

type UserController struct {
	beego.Controller
}
var (
	userService services.UserService
)

func (this *UserController) Index() {
	this.Data["url"] = ""
	this.TplName = "auth/login.html"
}

func (this *UserController) Register() {
	user := models.User{}
	this.ParseForm(&user)
	err := userService.Register(user)
	flash := beego.NewFlash()
	if err != nil {
		flash.Error("Register fail!")
		flash.Store(&this.Controller)
		this.Redirect("/login", 302)
	}
	flash.Notice("Register success!")
	flash.Store(&this.Controller)
	this.Redirect("/login", 302)
}
