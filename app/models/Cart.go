package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/jinzhu/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint `gorm:"index"`
}

/**
*
* Create cart
**/
func (cart *Cart) Create() (err error) {
	err = database.SQL.Create(&cart).Error
	return
}

/**
*
* Find cart by user id
**/
func (cart *Cart) FindByUserID(userID uint) error {
	var err error
	err = database.SQL.Where("user_id = ?", userID).First(&cart).Error
	return err
}
