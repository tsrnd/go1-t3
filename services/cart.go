package services

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	model "github.com/goweb3/models"
)

/**
* Get data home page
*
* return err
**/
func ProcessCartpage(data map[interface{}]interface{}) (err error) {
	o := orm.NewOrm()
	product := new(model.Product)
	var products []*model.Product
	qs := o.QueryTable(product)
	qs.Limit(10).RelatedSel().All(&products)
	data["products"] = products
	return
}

/**
* Process Add product to card
*
* return error
**/
func ProcessAddToCart(productID uint, userID uint) (err error) {
	product := model.Product{}
	product.Id = productID
	err = product.FindByID()
	if err != nil {
		fmt.Println("product does not exist")
		return err
	}
	cart := model.Cart{}
	err = cart.FindByUserID(userID)
	if err != nil {
		fmt.Println("cart does not exist")
		ProcessCreateCard(userID, &cart, &product)
	} else {
		fmt.Println("cart already exist")
		ProcessCheckExistCartProduct(productID, &cart, &product)
	}
	return
}

/**
* Process create card
*
* return error
**/
func ProcessCreateCard(userID uint, cart *model.Cart, product *model.Product) (err error) {
	user := &model.User{}
	user.Id = userID
	cart.User = user
	cart.Create()
	cartProduct := model.CartProduct{
		Cart:     cart,
		Product:  product,
		Quantity: uint(1),
	}
	ProcessCreateCartProduct(&cartProduct, cart, product)
	cartProduct.Create()
	return
}

/**
* Process check exist cart product
*
* return error
**/
func ProcessCheckExistCartProduct(productID uint, cart *model.Cart, product *model.Product) (err error) {
	cartProduct := model.CartProduct{}
	err = cartProduct.FindByCartIDAndProductID(cart.Id, productID)
	if err != nil {
		fmt.Println("cart product does not exist")
		ProcessCreateCartProduct(&cartProduct, cart, product)
	} else {
		fmt.Println("cart product already exist")
		cartProduct.Quantity++
		cartProduct.Update()
	}
	return
}

/**
* Process create card product
*
* return error
**/
func ProcessCreateCartProduct(cartProduct *model.CartProduct, cart *model.Cart, product *model.Product) (err error) {
	cartProduct.Cart = cart
	cartProduct.Product = product
	cartProduct.Quantity = uint(1)
	cartProduct.Create()
	return
}
