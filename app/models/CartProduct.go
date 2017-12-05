package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/jinzhu/gorm"
)

type CartProduct struct {
	gorm.Model
	CartID    uint    `schema:"cart_id"`
	ProductID uint    `schema:"product_id"`
	Quantity  uint    `schema:"quantity"`
	Product   Product `gorm:"ForeignKey:ProductID"`
}

/**
*
* Price follow quantity
**/
func (cartProduct *CartProduct) PriceFollowQuantity() uint {
	return uint(cartProduct.Quantity) * uint(cartProduct.Product.Price)
}

/**
*
* Delete
**/
func (cartProduct *CartProduct) Delete() error {
	err := database.SQL.Delete(&cartProduct).Error
	return err
}

/**
*
* Create cart product
**/
func (cartProduct *CartProduct) Create() (err error) {
	err = database.SQL.Create(&cartProduct).Error
	return
}

/**
*
* Update cart product
**/
func (cartProduct *CartProduct) Update() (err error) {
	database.SQL.Save(cartProduct)
	return
}

/**
*
* Find cart product by cart id and product id
**/
func (cartProduct *CartProduct) FindByCartIDAndProductID(cartID uint, productID uint) error {
	var err error
	err = database.SQL.Where("cart_id = ? AND product_id =?", cartID, productID).First(&cartProduct).Error
	return err
}

/**
*
* Find cart products by cart id
**/
func (cartProduct *CartProduct) FindByCartID(cartID uint) (cartProducts []CartProduct) {
	database.SQL.Where("cart_id = ?", cartID).Find(&cartProducts)
	return
}

/**
*
* Get all cart products by cart id
**/
func (cartProduct *CartProduct) GetByCartID(cartID uint) (cartProducts []CartProduct) {
	database.SQL.Preload("Product").Where("cart_id = ?", cartID).Find(&cartProducts)
	return
}
