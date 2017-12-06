package routers

import (
	"github.com/astaxie/beego"
	"github.com/goweb3/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/go", &controllers.UserController{})
}
