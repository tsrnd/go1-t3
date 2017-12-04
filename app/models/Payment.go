package models

import (
	"github.com/jinzhu/gorm"
	"github.com/goweb3/app/shared/database"
)
type Payment struct {
	gorm.Model
	OrderID uint
	PaymentMetnod string
	AccountNumber string
	Bank string
}

/**
*
* Create payment
**/
func(payment *Payment) Create() (err error) {
	err = database.SQL.Create(&payment).Error
	return
}

/**
*
* Find payment by order_id
**/
func (payment *Payment) FindByOrderId(orderId uint) (err error) {
	err = database.SQL.Where("order_id = ?", orderId).First(&payment).Error
	return err
}