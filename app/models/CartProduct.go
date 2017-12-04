package models

import (
	"fmt"

	"github.com/goweb3/app/shared/database"
	"github.com/jinzhu/gorm"
)

type CartProduct struct {
	gorm.Model
	CartID    uint `gorm:"index"`
	ProductID uint
	Quantity  uint
}

func (cartProduct *CartProduct) Create() (err error) {
	statement := "insert into cart_products (cart_id, product_id, quantity) values ($1,$2,$3) returning id"
	stmt, err := database.SQL.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(cartProduct.CartID, cartProduct.ProductID, cartProduct.Quantity).Scan(&cartProduct.ID)
	return
}

func (cartProduct *CartProduct) Update() (err error) {
	statement := "UPDATE cart_products SET quantity = $3 WHERE cart_id = $1 AND product_id =$2  returning id"
	stmt, err := database.SQL.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(cartProduct.CartID, cartProduct.ProductID, cartProduct.Quantity).Scan(&cartProduct.ID)
	return
}

func (cartProduct *CartProduct) CheckExistCartProduct(cartID int, productID int) error {
	var err error
	fmt.Println("cccc", cartID, productID)
	err = database.SQL.QueryRow("SELECT id, cart_id, product_id, quantity FROM cart_products WHERE cart_id = $1 AND product_id = $2", cartID, productID).Scan(&cartProduct.ID, &cartProduct.CartID, &cartProduct.ProductID, &cartProduct.Quantity)
	return err
}
