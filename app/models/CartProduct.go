package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/jinzhu/gorm"
)

type CartProduct struct {
	gorm.Model
	CartID    uint `gorm:"index"`
	ProductID uint
	Quantity  uint
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
