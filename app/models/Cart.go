package models

import (
	"github.com/jinzhu/gorm"
	"github.com/goweb3/app/shared/database"	
)

type Cart struct {
	gorm.Model	
	CartProducts []CartProduct
	UserID uint `schema:"user_id"`
}

/**
*
* Find cart by user_id
**/
func (cart *Cart) FindByUserId(userId int) (err error) {
	err = database.SQL.Where("user_id = ?", userId).First(&cart).Error
	return err
}

/**
*
* Total price cart
**/
func (cart *Cart) TotalPrice() (int) {
	sum := 0
	for _, v := range cart.CartProducts {
		sum += int(v.Quantity) * v.Product.Price
	}
	return sum + 20000
}

/**
*
* Delete
**/
func (cart *Cart) Delete() (error) {
	err := database.SQL.Delete(&cart).Error
	return err
}