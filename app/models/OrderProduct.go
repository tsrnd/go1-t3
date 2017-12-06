package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/jinzhu/gorm"
)

type OrderProduct struct {
	gorm.Model
	OrderID   uint `schema:"order_id"`
	ProductID uint `schema:"product_id"`
	Quantity  uint `schema:"quantity"`
	Price     uint `schema:"price"`
}

/**
*
* Create orderProduct
**/
func (orderProduct *OrderProduct) Create() (err error) {
	err = database.SQL.Create(&orderProduct).Error
	return
}
