package controllers

import (
	"github.com/astaxie/beego"
	service "github.com/goweb3/services"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	service.ProcessHompage(c.Data)
	c.Data["url"] = ""
	c.TplName = "home/index.html"
}
