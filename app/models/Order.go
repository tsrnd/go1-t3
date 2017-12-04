package models

import (
	"github.com/jinzhu/gorm"
	"github.com/goweb3/app/shared/database"
)
type Order struct {
	gorm.Model
	UserID uint
	NameReceiver string
	Address string
	Status uint
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
