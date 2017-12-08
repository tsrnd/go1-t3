package routers

import (
	"github.com/astaxie/beego"
	"github.com/goweb3/controllers"
	"github.com/goweb3/filters"
)

func init() {
	loadFilters()
	beego.Router("/", &controllers.HomeController{}, "get:Index")
	beego.Router("/login", &controllers.LoginController{}, "get:Index;post:Login")
	beego.Router("/guest/register", &controllers.UserController{}, "post:Register")
	beego.Router("/cart", &controllers.CartController{}, "get:Index")
	beego.Router("/cart/:id:int", &controllers.CartController{}, "post:AddToCart")
}
func loadFilters() {
	beego.InsertFilter("/cart", beego.BeforeRouter, filters.ProtectUserPages)
}
