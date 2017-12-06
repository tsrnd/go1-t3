package controllers

import (
	"fmt"
	"net/http"
)

type UserController struct {
	Controller
}

func (this *UserController) Get() {
	this.TplName = "auth/login.html"
}

func (this *UserController) Post() {
	this.Ctx.Request.ParseForm()
	fmt.Println(this.Ctx.Request.FormValue("username"))
	this.Ctx.Redirect(http.StatusSeeOther, "/")
}
