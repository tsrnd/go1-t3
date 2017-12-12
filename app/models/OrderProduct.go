package models

import (
	"github.com/goweb3/app/shared/database"
)

type OrderProduct struct {
	BaseModel
	OrderID   uint `db:"order_id"`
	ProductID uint `db:"product_id"`
	Quantity  uint `db:"quantity"`
	Price     uint `db:"price"`
	Product *Product
}

/**
*
* Create orderProduct
**/
func (orderProduct *OrderProduct) Create() (err error) {
	statement := "insert into order_products (order_id, product_id, quantity, price) values ($1, $2, $3, $4) returning id"
	stmt, err := database.SQL.Prepare(statement)
	if err != nil {
		return
	}
	err = stmt.QueryRow(orderProduct.OrderID, orderProduct.ProductID, orderProduct.Quantity, orderProduct.Price).Scan(&orderProduct.ID)
	return
}

/**
*
* Load Product
**/
func (orderProduct *OrderProduct) LoadProducts() (err error) {
	err = database.SQL.QueryRow("SELECT id, name, description, quantity, price FROM products WHERE deleted_at is null AND id = $1", orderProduct.ProductID).Scan(&orderProduct.Product.ID, &orderProduct.Product.Name, &orderProduct.Product.Description, &orderProduct.Product.Quantity, &orderProduct.Product.Price)
	return
}
