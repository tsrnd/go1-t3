package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/goweb3/app/shared/database"
)

type CartProduct struct {
	Id       uint
	Cart     *Cart    `orm:"column(cart_id);rel(fk)"`
	Product  *Product `orm:"column(product_id);rel(fk)"`
	Quantity uint
}

func (p *CartProduct) TableName() string {
	return "cart_products"
}

func init() {
	orm.RegisterModel(new(CartProduct))
}

/**
*
* Price follow quantity
**/
func (cartProduct *CartProduct) PriceFollowQuantity() int {
	return int(cartProduct.Quantity) * cartProduct.Product.Price
}

/**
*
* Delete
**/
func (cartProduct *CartProduct) Delete() error {
	err := database.SQL.Delete(&cartProduct).Error
	return err
}

/**
*
* Create cart product
**/
func (cartProduct *CartProduct) Create() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(cartProduct)
	return
}

/**
*
* Update cart product
**/
func (cartProduct *CartProduct) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(cartProduct)
	return
}

/**
*
* Find cart product by cart id and product id
**/
func (c *CartProduct) FindByCartIDAndProductID(cartID uint, productID uint) (err error) {
	o := orm.NewOrm()
	err = o.QueryTable("cart_products").Filter("cart_id", cartID).Filter("product_id", productID).One(c)
	return
}

/**
*
* Find cart products by cart id
**/
func (cartProduct *CartProduct) FindByCartID(cartID uint) (cartProducts []CartProduct) {
	database.SQL.Where("cart_id = ?", cartID).Find(&cartProducts)
	return
}
