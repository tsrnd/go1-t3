package routers

import (
	"github.com/astaxie/beego"
	"github.com/goweb3/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/cart", &controllers.CartController{}, "get:Index")
	beego.Router("/cart/:id:int", &controllers.CartController{}, "post:AddToCart")
}
