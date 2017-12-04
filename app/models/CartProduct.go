package models

import (
	"github.com/jinzhu/gorm"
	"github.com/goweb3/app/shared/database"		
)
type CartProduct struct {
	gorm.Model
	CartID uint 		`schema:"cart_id"`
	ProductID uint		`schema:"product_id"`
	Quantity uint 		`schema:"quantity"`
	Product Product
}

/**
*
* Price follow quantity
**/
func (cartProduct *CartProduct) PriceFollowQuantity() (int) {
	return int(cartProduct.Quantity) * cartProduct.Product.Price
}


/**
*
* Delete
**/
func (cartProduct *CartProduct) Delete() (error) {
	err := database.SQL.Delete(&cartProduct).Error
	return err
}
