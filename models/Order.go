package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/astaxie/beego/orm"
)
type Order struct {
	Id uint
	User *User `orm:"column(user_id);rel(fk)"`
	NameReceiver string
	Address string
	Status uint
}
func init() {
    orm.RegisterModel(new(Order))
}
/**
*
* Create order
**/
func(order *Order) Create() (err error) {
	err = database.SQL.Create(&order).Error
	return
}

/**
*
* Find order by order_id
**/
func (order *Order) FindById(id uint) (err error) {
	err = database.SQL.Where("id = ?", id).First(&order).Error
	return err
}
