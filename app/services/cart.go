package service

import (
	"fmt"
	"time"

	"github.com/goweb3/app/models"
)

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
		fmt.Println(err)
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
