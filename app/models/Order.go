package models

import (
	// "github.com/goweb3/app/shared/database"
)
type Order struct {
	BaseModel
	UserID uint			`db:"user_id"`
	NameReceiver string	`db:"name_receiver"`
	Address string		`db:"address"`
	Status uint			`db:"status"`
}

/**
*
* Create order
**/
func(order *Order) Create() (err error) {
	// err = database.SQL.Create(&order).Error
	return
}

/**
*
* Find order by order_id
**/
func (order *Order) FindById(id uint) (err error) {
	// err = database.SQL.Where("id = ?", id).First(&order).Error
	return err
}
