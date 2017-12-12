package models

import (
	"github.com/goweb3/app/shared/database"
)

type Order struct {
	BaseModel
	UserID        uint   `db:"user_id"`
	NameReceiver  string `db:"name_receiver"`
	Address       string `db:"address"`
	Status        uint   `db:"status"`
	OrderProducts []OrderProduct
}

/**
*
* Create order
**/
func (order *Order) Create() (err error) {
	statement := "insert into orders (user_id, name_receiver, address, status) values ($1, $2, $3, $4) returning id"
	stmt, err := database.SQL.Prepare(statement)
	if err != nil {
		return
	}
	err = stmt.QueryRow(order.UserID, order.NameReceiver, order.Address, order.Status).Scan(&order.ID)
	return
}

/**
*
* Find order by order_id
**/
func (order *Order) FindById(id uint) (err error) {
	err = database.SQL.QueryRow("SELECT id, user_id, name_receiver, address, status FROM users WHERE deleted_at is null AND id = $1", id).Scan(&order.ID, &order.UserID, &order.NameReceiver, &order.Address, &order.Status)
	return err
}

/**
*
* Load OrderProducts
**/
func (order *Order) LoadOrderProducts() (err error) {
	rows, err := database.SQL.Query("select id, order_id, product_id, quantity from order_products where deleted_at is null AND order_id = $1", order.ID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		orderProduct := OrderProduct{}
		err := rows.Scan(&orderProduct.ID, &orderProduct.OrderID, &orderProduct.Quantity)
		if err != nil {
			return err
		}
		order.OrderProducts = append(order.OrderProducts, orderProduct)
	}
	return
}
