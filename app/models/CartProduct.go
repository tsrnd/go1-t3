package models

import (
	"github.com/goweb3/app/shared/database"
)

type CartProduct struct {
	BaseModel
	CartID    uint    `db:"cart_id"`
	ProductID uint    `db:"product_id"`
	Quantity  uint    `db:"quantity"`
	Product   Product
}

/**
*
* Price follow quantity
**/
func (cartProduct *CartProduct) PriceFollowQuantity() uint {
	return uint(cartProduct.Quantity) * uint(cartProduct.Product.Price)
}

/**
*
* Load Product
**/
func (cartProduct *CartProduct) LoadProducts() (err error) {
	err = database.SQL.QueryRow("SELECT id, name, description, quantity, price FROM products WHERE deleted_at is null AND id = $1", cartProduct.ProductID).Scan(&cartProduct.Product.ID, &cartProduct.Product.Name, &cartProduct.Product.Description, &cartProduct.Product.Quantity, &cartProduct.Product.Price)
	return
}

/**
*
* Delete
**/
func (cartProduct *CartProduct) Delete() (err error) {
	// err := database.SQL.Delete(&cartProduct).Error
	return err
}

/**
*
* Create cart product
**/
func (cartProduct *CartProduct) Create() (err error) {
	// err = database.SQL.Create(&cartProduct).Error
	return
}

/**
*
* Update cart product
**/
func (cartProduct *CartProduct) Update() (err error) {
	// database.SQL.Save(cartProduct)
	return
}

/**
*
* Find cart product by cart id and product id
**/
func (cartProduct *CartProduct) FindByCartIDAndProductID(cartID uint, productID uint) (err error) {
	// err = database.SQL.Where("cart_id = ? AND product_id =?", cartID, productID).First(&cartProduct).Error
	return err
}

/**
*
* Find cart products by cart id
**/
func (cartProduct *CartProduct) FindByCartID(cartID uint) (cartProducts []CartProduct) {
	// database.SQL.Where("cart_id = ?", cartID).Find(&cartProducts)
	return
}

/**
*
* Get all cart products by cart id
**/
func (cartProduct *CartProduct) GetByCartID(cartID uint) (cartProducts []CartProduct) {
	// database.SQL.Preload("Product").Where("cart_id = ?", cartID).Find(&cartProducts)
	return
}
