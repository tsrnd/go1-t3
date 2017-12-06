package service

import (
	"net/http"
	"strconv"

	"github.com/goweb3/app/models"
	"github.com/goweb3/app/shared/flash"
	"github.com/jianfengye/web-golang/web/session"
)

/**
*
* Process add product to cart
**/
func ProcessAddToCard(w http.ResponseWriter, r *http.Request, productID uint) error {
	var err error
	sess, _ := session.SessionStart(r, w)
	userID, _ := strconv.Atoi(sess.Get("id"))
	product := models.Product{}
	err = product.FindByID(productID)
	if err != nil {
		flash.SetFlash(w, flash.Flash{"Product does not exist", flash.FlashError})
		return err
	}
	cart := models.Cart{}
	err = cart.FindByUserID(uint(userID))
	if err != nil { // cart not exist
		err = ProcessCreateCard(uint(userID), product.ID, &cart)
		if err != nil {
			return err
		}
	} else { // cart exist
		err = ProcessCheckExistCartProduct(cart.ID, product.ID)
		if err != nil {
			return err
		}
	}
	flash.SetFlash(w, flash.Flash{"Add success", flash.FlashSuccess})
	return nil
}

/**
* Process create card
*
* return error
**/
func ProcessCreateCard(userID uint, productID uint, cart *models.Cart) (err error) {
	var quantity uint = 1
	*cart = models.Cart{
		UserID: userID,
	}
	err = cart.Create()
	cartProduct := models.CartProduct{}
	ProcessCartProductData(uint(cart.ID), productID, quantity, &cartProduct)
	err = cartProduct.Create()
	return
}

/**
* Process check exist cart product
*
* return error
**/
func ProcessCheckExistCartProduct(cardID uint, productID uint) error {
	var err error
	cartProduct := models.CartProduct{}
	var quantity uint = 1
	err = cartProduct.FindByCartIDAndProductID(cardID, productID)
	if err != nil { // cart product not exist
		ProcessCartProductData(cardID, productID, quantity, &cartProduct)
		err = cartProduct.Create()
	} else { // cart product exist
		quantity = cartProduct.Quantity
		quantity++
		cartProduct.Quantity = quantity
		err = cartProduct.Update()
	}
	return err
}

/**
* Process cart product data
*
* return void
**/
func ProcessCartProductData(cardID uint, productID uint, quantity uint, cartProduct *models.CartProduct) {
	*cartProduct = models.CartProduct{
		CartID:    cardID,
		ProductID: productID,
		Quantity:  quantity,
	}
}

/**
* Process get count cart products data by cart id
*
* return uint
**/
func ProcessGetCountCartProduct(userID uint) uint {
	cart := models.Cart{}
	cart.FindByUserID(userID)
	cartProduct := models.CartProduct{}
	cartProducts := cartProduct.FindByCartID(cart.ID)
	sum := 0
	for _, v := range cartProducts {
		sum += int(v.Quantity)
	}
	return uint(sum)
}

/**
* Process cart page
*
* return products
**/
func ProcessCartPage(w http.ResponseWriter, r *http.Request, data map[string]interface{}) (err error) {
	sess, _ := session.SessionStart(r, w)
	userID, _ := strconv.Atoi(sess.Get("id"))
	cart := models.Cart{}
	cart.FindByUserID(uint(userID))
	cartProduct := models.CartProduct{}
	cartProducts := cartProduct.GetByCartID(cart.ID)
	sum := 0
	for _, v := range cartProducts {
		sum += int(v.Quantity) * v.Product.Price
	}
	data["cartProducts"] = cartProducts
	data["priceTotal"] = sum
	return
}

/**
* Process del cart
*
* return err
**/
func ProcessDelCartProduct(w http.ResponseWriter, r *http.Request, productID uint) (err error) {
	sess, _ := session.SessionStart(r, w)
	userID, _ := strconv.Atoi(sess.Get("id"))
	cart := models.Cart{}
	cart.FindByUserID(uint(userID))
	cartProduct := models.CartProduct{}
	err = cartProduct.FindByCartIDAndProductID(cart.ID, productID)
	if err != nil {
		flash.SetFlash(w, flash.Flash{"Cart Product does not exist", flash.FlashError})
		return err
	}
	cartProduct.Delete()
	flash.SetFlash(w, flash.Flash{"Delete success", flash.FlashSuccess})
	return
}
