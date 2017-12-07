package routers

import (
	"github.com/astaxie/beego"
	"github.com/goweb3/controllers"
)

type Router struct {
	Url           string
	Controller    beego.ControllerInterface
	MethodMapping []string
}

type Routers []Router

var routers = Routers{
	Router{
		Url:           "/users",
		Controller:    &controllers.UserController{},
		MethodMapping: "get:GetAll;post:Post"
	},
	Router{
		Url:           "/products",
		Controller:    &controllers.ProductController{},
		MethodMapping: []string{"get:GetAll"},
	},
	Router{
		Url:           "/categories",
		Controller:    &controllers.CategoryController{},
		MethodMapping: []string{"get:GetAll"},
	},
}

func init() {
<<<<<<< HEAD
	for _, router := range routers {
		beego.Router(router.Url, router.Controller, router.MethodMapping...)
	}
}
=======
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/guest/register", &controllers.UserController{}, "post:Register")
}
>>>>>>> 222b04d51827c0f13d0f60cb79795a8f0110f1f5
