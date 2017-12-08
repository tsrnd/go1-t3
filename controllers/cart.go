package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	service "github.com/goweb3/services"
)

type CartController struct {
	beego.Controller
}

func (this *CartController) Index() {
	this.Data["url"] = "/"
	this.TplName = "cart/index.html"
}

func (this *CartController) AddToCart() {
	productID, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	// sess := this.StartSession()
	// user := reflect.ValueOf(sess.Get("auth"))
	// userID := user.FieldByName("Id")
	userID := uint(1)
	service.ProcessAddToCart(uint(productID), uint(userID))
	this.Ctx.Redirect(302, "/cart")
}
