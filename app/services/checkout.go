package service

import "github.com/goweb3/app/models"
import "database/sql"
import "time"
import "net/http"
import "github.com/jianfengye/web-golang/web/session"
import "strconv"
import "strings"
import "github.com/goweb3/app/shared/database"

/**
* Process checkout order with transaction
*
* return err
**/
func CheckoutOrder(w http.ResponseWriter, r *http.Request) (err error) {
	sess, _ := session.SessionStart(r, w)
	userId, _ := strconv.ParseUint(sess.Get("id"), 10, 32)
	cart := models.Cart{}
	cart.FindByUserID(uint(userId))
	cart.LoadCartProducts()
	order := models.Order{
		UserID:       uint(userId),
		NameReceiver: strings.Trim(r.FormValue("name-receiver"), " "),
		Address:      strings.Trim(r.FormValue("address"), " "),
		Status:       1,
	}
	/* begin transaction */
	tx, _ := database.SQL.Begin()
	/* Create order */
	if err := TransactionCreateOrder(&order, tx); err != nil {
		tx.Rollback()
	}
	for i := 0; i < len(cart.CartProducts); i++ {
		orderProduct := models.OrderProduct{
			OrderID:   order.ID,
			ProductID: cart.CartProducts[i].ProductID,
			Quantity:  cart.CartProducts[i].Quantity,
			Price:     cart.CartProducts[i].PriceFollowQuantity(),
		}
		/* Create orderProduct */
		if err := TransactionCreateOrderProduct(&orderProduct, tx); err != nil {
			tx.Rollback()
		}
		/* Delete cartProduct */
		if err := TransactionDeleteCartProduct(&cart.CartProducts[i], tx); err != nil {
			tx.Rollback()
		}
	}
	/* Delete cart */
	if err := TransactionDeleteCart(&cart, tx); err != nil {
		tx.Rollback()
	}
	/* Create payment */
	payment := models.Payment{
		OrderID:       order.ID,
		AccountNumber: strings.Trim(r.FormValue("card_number"), " "),
		Bank:          strings.Trim(r.FormValue("bank"), " "),
	}
	if err := TransactionCreatePayment(&payment, tx); err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return
}

/**
* Process create order with transaction
*
* return err
**/
func TransactionCreateOrder(order *models.Order, tx *sql.Tx) (err error) {
	statement := "insert into orders (user_id, name_receiver, address, status) values ($1, $2, $3, $4) returning id"
	stmt, _ := tx.Prepare(statement)

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
	stmt, _ := tx.Prepare(statement)

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
	stmt, _ := tx.Prepare(statement)

	err = stmt.QueryRow(orderProduct.OrderID, orderProduct.ProductID, orderProduct.Quantity, orderProduct.Price).Scan(&orderProduct.ID)
	return
}

/**
* Process delete cart with transaction
*
* return err
**/
func TransactionDeleteCart(cart *models.Cart, tx *sql.Tx) (err error) {
	stmt, _ := tx.Prepare("update carts set deleted_at = $1 where id = $2")
	_, err = stmt.Exec(time.Now(), cart.ID)
	return
}

/**
* Process delete cartProduct with transaction
*
* return err
**/
func TransactionDeleteCartProduct(CartProduct *models.CartProduct, tx *sql.Tx) (err error) {
	stmt, _ := tx.Prepare("update cart_products set deleted_at = $1 where id = $2")
	_, err = stmt.Exec(time.Now(), CartProduct.ID)
	return
}
