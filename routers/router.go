package routers

import (
	"github.com/goweb3/controllers"
	"github.com/astaxie/beego"
	"github.com/goweb3/filters"
)

func init() {
	loadFilters()
	beego.Router("/", &controllers.HomeController{}, "get:Index")
	beego.Router("/login", &controllers.LoginController{}, "get:Index;post:Login")
	beego.Router("/guest/register", &controllers.UserController{}, "post:Register")
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")	
}
func loadFilters() {
	beego.InsertFilter("/cart", beego.BeforeRouter, filters.ProtectUserPages)
	beego.InsertFilter("/login", beego.BeforeRouter, filters.UserIsLogged)
}