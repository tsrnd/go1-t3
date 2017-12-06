package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/astaxie/beego/orm"
)

type Cart struct {
	Id	uint
	CartProducts []*CartProduct `orm:"reverse(many)"`
	User *User `orm:"column(user_id);rel(fk)"`
}
func init() {
    orm.RegisterModel(new(Cart))
}
/**
*
* Find cart by user_id
**/
func (cart *Cart) FindByUserId(userId int) (err error) {
	err = database.SQL.Where("user_id = ?", userId).First(&cart).Error
	return err
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
* Total price cart
**/
func (cart *Cart) TotalPrice() (int) {
	sum := 0
	for _, v := range cart.CartProducts {
		sum += int(v.Quantity) * v.Product.Price
	}
	return sum + 20000
}

/**
*
* Delete
**/
func (cart *Cart) Delete() (error) {
	err := database.SQL.Delete(&cart).Error
	return err
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
