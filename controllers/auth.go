package controllers

import (
	"github.com/astaxie/beego"
)

// AuthController operations for Auth
type AuthController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Register", c.Register)
}

// Login ...
// @Title Create
// @Description create Auth
// @Param	body		body 	models.Auth	true		"body for Auth content"
// @Success 201 {object} models.Auth
// @Failure 403 body is empty
// @router /login [post]
func (c *AuthController) Login() {

}

// Register ...
// @Title GetOne
// @Description get Auth by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Auth
// @Failure 403 :id is empty
// @router /register [post]
func (c *AuthController) Register() {

}
