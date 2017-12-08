package routers

import (
	"github.com/astaxie/beego"
	"github.com/goweb3/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "get:Index")
	beego.Router("/login", &controllers.LoginController{}, "get:Index;post:Login")
	beego.Router("/guest/register", &controllers.UserController{}, "get:Index;post:Register")
	beego.Router("/products/:id", &controllers.ProductController{}, "get:Show")
}
