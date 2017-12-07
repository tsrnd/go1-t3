package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/astaxie/beego/orm"
)

type OrderProduct struct {
	Id uint
	Order *Order `orm:"column(order_id);rel(fk)"`
	Product *Product `orm:"column(product_id);rel(fk)"`
	Quantity uint
	Price int
}
func init() {
    orm.RegisterModel(new(OrderProduct))
}
/**
*
* Create orderProduct
**/
func(orderProduct *OrderProduct) Create() (err error) {
	err = database.SQL.Create(&orderProduct).Error
	return
}