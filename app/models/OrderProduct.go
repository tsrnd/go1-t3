package models

import (
	"github.com/jinzhu/gorm"
	"github.com/goweb3/app/shared/database"
)

type OrderProduct struct {
	gorm.Model
	OrderID uint		`schema:"order_id"`
	ProductID uint		`schema:"product_id"`
	Quantity uint		`schema:"quantity"`
	Price int			`schema:"price"`
}

/**
*
* Create orderProduct
**/
func(orderProduct *OrderProduct) Create() (err error) {
	err = database.SQL.Create(&orderProduct).Error
	return
}