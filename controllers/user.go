package controllers

import (
	"github.com/astaxie/beego"
	"github.com/goweb3/models"
	"github.com/goweb3/validations"
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
	registerForm := validations.RegisterForm{}
	this.ParseForm(&registerForm)
	flash := beego.NewFlash()
	if !registerForm.Validate() {
		flash.Error("Register fail!")
		flash.Store(&this.Controller)
		this.Redirect("/login", 302)
	} else {
		user := models.User{}
		this.ParseForm(&user)
		err := userService.Register(user)
		if err != nil {
			flash.Error("Register fail!")
			flash.Store(&this.Controller)
			this.Redirect("/login", 302)
		}
		flash.Notice("Register success!")
		flash.Store(&this.Controller)
		this.Redirect("/login", 302)
	}
}
