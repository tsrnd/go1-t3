package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/astaxie/beego/orm"
)

type CartProduct struct {
	Id uint
	Cart *Cart `orm:"column(cart_id);rel(fk)"`		
	Product *Product `orm:"column(product_id);rel(fk)"`	
	Quantity uint
}
func init() {
    orm.RegisterModel(new(CartProduct))
}
/**
*
* Price follow quantity
**/
func (cartProduct *CartProduct) PriceFollowQuantity() (int) {
	return int(cartProduct.Quantity) * cartProduct.Product.Price
}


/**
*
* Delete
**/
func (cartProduct *CartProduct) Delete() (error) {
	err := database.SQL.Delete(&cartProduct).Error
	return err
}

/**
*
* Create cart product
**/
func (cartProduct *CartProduct) Create() (err error) {
	err = database.SQL.Create(&cartProduct).Error
	return
}

/**
*
* Update cart product
**/
func (cartProduct *CartProduct) Update() (err error) {
	database.SQL.Save(cartProduct)
	return
}

/**
*
* Find cart product by cart id and product id
**/
func (cartProduct *CartProduct) FindByCartIDAndProductID(cartID uint, productID uint) error {
	var err error
	err = database.SQL.Where("cart_id = ? AND product_id =?", cartID, productID).First(&cartProduct).Error
	return err
}

/**
*
* Find cart products by cart id
**/
func (cartProduct *CartProduct) FindByCartID(cartID uint) (cartProducts []CartProduct) {
	database.SQL.Where("cart_id = ?", cartID).Find(&cartProducts)
	return
}
