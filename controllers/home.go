package controllers

import (
	"github.com/astaxie/beego"
	service "github.com/goweb3/services"
	
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["notice"]; ok {
        c.Data["notice"] = n
    }
	service.ProcessHompage(c.Data)
	c.Data["url"] = ""
	c.TplName = "home/index.html"
}
