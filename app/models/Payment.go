package models

import (
	"github.com/goweb3/app/shared/database"
)
type Payment struct {
	BaseModel
	OrderID uint		 `db:"order_id"`
	AccountNumber string `db:"account_number"`
	Bank string			 `db:"bank"`
}

/**
*
* Create payment
**/
func(payment *Payment) Create() (err error) {
	statement := "insert into payments (order_id, account_number, bank) values ($1, $2, $3) returning id"
	stmt, err := database.SQL.Prepare(statement)
	if err != nil {
		return
	}
	err = stmt.QueryRow(payment.OrderID, payment.AccountNumber, payment.Bank).Scan(&payment.ID)
	return
}

/**
*
* Find payment by order_id
**/
func (payment *Payment) FindByOrderId(orderId uint) (err error) {
	err = database.SQL.QueryRow("SELECT id, order_id, account_number, bank FROM payments WHERE deleted_at is null AND order_id = $1", orderId).Scan(&payment.ID, &payment.OrderID, &payment.AccountNumber, &payment.Bank)
	return err
}
