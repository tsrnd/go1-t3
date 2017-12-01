package service

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/goweb3/app/models"
	"github.com/goweb3/app/shared/flash"
	"github.com/jianfengye/web-golang/web/session"
)

func ProcessAddToCard(w http.ResponseWriter, r *http.Request, productID int) error {
	var err error
	sess, _ := session.SessionStart(r, w)
	userID, _ := strconv.Atoi(sess.Get("id"))
	product := models.Product{}
	err = product.FindById(productID)
	if err != nil {
		flash.SetFlash(w, flash.Flash{"Product does not exist", flash.FlashError})
		return err
	}
	cart := models.Cart{}
	err = cart.CheckExistCart(userID)
	if err != nil { // cart not exist
		err = ProcessCreateCard(userID, product.ID, &cart)
		if err != nil {
			return err
		}
	} else { // cart exist
		err = ProcessCheckExistCartProduct(cart.ID, product.ID)
		if err != nil {
			return err
		}
	}
	fmt.Println("session", userID, "productid", productID)
	return nil
}

/**
* Process create card
*
* return error
**/
func ProcessCreateCard(userID int, productID int, cart *models.Cart) (err error) {
	quantity := 1
	*cart = models.Cart{
		ID:        0,
		UserID:    userID,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: time.Time{},
	}
	err = cart.Create()
	cartProduct := models.CartProduct{}
	ProcessCartProductData(cart.ID, productID, quantity, &cartProduct)
	err = cartProduct.Create()
	return
}

/**
* Process check exist cart product
*
* return error
**/
func ProcessCheckExistCartProduct(cardID int, productID int) error {
	var err error
	cartProduct := models.CartProduct{}
	quantity := 1
	err = cartProduct.CheckExistCartProduct(cardID, productID)
	if err != nil { // cart product not exist
		ProcessCartProductData(cardID, productID, quantity, &cartProduct)
		_ = cartProduct.Create()
	} else { // cart product exist
		quantity = cartProduct.Quantity
		quantity++
		ProcessCartProductData(cardID, productID, quantity, &cartProduct)
		err = cartProduct.Update()
	}
	return err
}

/**
* Process cart product data
*
* return void
**/
func ProcessCartProductData(cardID int, productID int, quantity int, cartProduct *models.CartProduct) {
	*cartProduct = models.CartProduct{
		ID:        0,
		CartID:    cardID,
		ProductID: productID,
		Quantity:  quantity,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: time.Time{},
	}
}
