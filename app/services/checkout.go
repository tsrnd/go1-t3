package service

import "github.com/goweb3/app/models"
import "database/sql"
import "time"

/**
* Process create order with transaction
*
* return err
**/
func TransactionCreateOrder(order *models.Order, tx *sql.Tx) (err error) {
	statement := "insert into orders (user_id, name_receiver, address, status) values ($1, $2, $3, $4) returning id"
	stmt,_ := tx.Prepare(statement)

	err = stmt.QueryRow(order.UserID, order.NameReceiver, order.Address, order.Status).Scan(&order.ID)
	return
}

/**
* Process create payment with transaction
*
* return err
**/
func TransactionCreatePayment(payment *models.Payment, tx *sql.Tx) (err error) {
	statement := "insert into payments (order_id, account_number, bank) values ($1, $2, $3) returning id"		
	stmt,_ := tx.Prepare(statement)

	err = stmt.QueryRow(payment.OrderID, payment.AccountNumber, payment.Bank).Scan(&payment.ID)
	
	return	
}

/**
* Process create orderProduct with transaction
*
* return err
**/
func TransactionCreateOrderProduct(orderProduct *models.OrderProduct, tx *sql.Tx) (err error) {
	statement := "insert into order_products (order_id, product_id, quantity, price) values ($1, $2, $3, $4) returning id"		
	stmt,_ := tx.Prepare(statement)

	err = stmt.QueryRow(orderProduct.OrderID, orderProduct.ProductID, orderProduct.Quantity, orderProduct.Price).Scan(&orderProduct.ID)
	return	
}

/**
* Process delete cart with transaction
*
* return err
**/
func TransactionDeleteCart(cart *models.Cart, tx *sql.Tx) (err error) {
	stmt,_ := tx.Prepare("update carts set deleted_at = $1 where id = $2")
	_, err = stmt.Exec(time.Now(), cart.ID)	
	return
}

/**
* Process delete cartProduct with transaction
*
* return err
**/
func TransactionDeleteCartProduct(CartProduct *models.CartProduct, tx *sql.Tx) (err error) {
	stmt,_ := tx.Prepare("update cart_products set deleted_at = $1 where id = $2")
	_, err = stmt.Exec(time.Now(), CartProduct.ID)	
	return
}