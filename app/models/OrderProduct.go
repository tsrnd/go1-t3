package models

import (
	// "github.com/goweb3/app/shared/database"
)

type OrderProduct struct {
	BaseModel
	OrderID   uint `db:"order_id"`
	ProductID uint `db:"product_id"`
	Quantity  uint `db:"quantity"`
	Price     uint `db:"price"`
}

/**
*
* Create orderProduct
**/
func (orderProduct *OrderProduct) Create() (err error) {
	// err = database.SQL.Create(&orderProduct).Error
	return
}
