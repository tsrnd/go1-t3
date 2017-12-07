package routers

import (
	"github.com/goweb3/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "get:Index")
	beego.Router("/login", &controllers.LoginController{}, "get:Index;post:Login")
	beego.Router("/guest/register", &controllers.UserController{}, "post:Register")
}