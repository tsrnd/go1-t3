package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/goweb3/app/shared/database"
)

type Cart struct {
	Id           uint
	CartProducts []*CartProduct `orm:"reverse(many)"`
	User         *User          `orm:"column(user_id);rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Cart))
}
func (p *Cart) TableName() string {
	return "carts"
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
	o := orm.NewOrm()
	_, err = o.Insert(cart)
	return
}

/**
*
* Total price cart
**/
func (cart *Cart) TotalPrice() int {
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
func (cart *Cart) Delete() error {
	err := database.SQL.Delete(&cart).Error
	return err
}

/**
*
* Find cart by user id
**/
func (c *Cart) FindByUserID(userID uint) (err error) {
	o := orm.NewOrm()
	err = o.QueryTable("carts").Filter("user_id", userID).One(c)
	return
}
