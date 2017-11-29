package models

import (
	"fmt"
	"time"

	"github.com/goweb3/app/shared/database"
)

type CartProduct struct {
	ID        int       `db:"id" bson:"id"`
	CartID    int       `db:"cart_id" bson:"cart_id"`
	ProductID int       `db:"product_id" bson:"product_id"`
	Quantity  int       `db:"quantity" bson:"quantity"`
	CreatedAt time.Time `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time `db:"updated_at" bson:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" bson:"deleted_at"`
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
