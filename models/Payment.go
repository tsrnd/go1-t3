package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/astaxie/beego/orm"
)
type Payment struct {
	Id uint
	Order *Order `orm:"column(order_id);rel(fk)"`
	AccountNumber string
	Bank string
}

func init() {
    orm.RegisterModel(new(Payment))
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
