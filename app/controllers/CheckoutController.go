package controller

import "net/http"
import "strconv"
import "github.com/goweb3/app/shared/view"
import "github.com/gorilla/csrf"
import "github.com/goweb3/app/models"
import "github.com/jianfengye/web-golang/web/session"
import "github.com/goweb3/app/shared/database"
import "github.com/goweb3/app/shared/cookie"
import "strings"
import service "github.com/goweb3/app/services"

func Checkout(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.SessionStart(r, w)

	userId, _ := strconv.ParseUint(sess.Get("id"), 10, 32)

	cart := models.Cart{}
	cart.FindByUserID(uint(userId))
	cart.LoadCartProducts()
	for i, _ := range cart.CartProducts {
		cart.CartProducts[i].LoadProducts()
		cart.CartProducts[i].Product.LoadProductImage()
	}
	v := view.New(r)
	v.Vars[csrf.TemplateTag] = csrf.TemplateField(r)
	if content := cookie.GetMessage(w, r, "ErrorCheckout"); content != "" {
		v.Vars["titleMessage"] = "Error"
		v.Vars["contentMessage"] = content
	} else if content := cookie.GetMessage(w, r, "SuccessCheckout"); content != "" {
		v.Vars["titleMessage"] = "Success"
		v.Vars["contentMessage"] = content
	}
	v.Vars["cart"] = cart
	v.Vars["totalPrice"] = cart.TotalPrice()
	v.Name = "checkout/index"
	v.Render(w)
}

func CheckoutPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sess, _ := session.SessionStart(r, w)
	userId, _ := strconv.ParseUint(sess.Get("id"), 10, 32)
	order := models.Order{
		UserID: uint(userId),
		NameReceiver : strings.Trim(r.FormValue("name-receiver"), " "),
		Address: strings.Trim(r.FormValue("address"), " "),
		Status: 1,
	}
	tx, _ := database.SQL.Begin()
	message := "Order failed!"	
	/* Create order */
	if service.TransactionCreateOrder(&order, tx) != nil {
		tx.Rollback()
		cookie.SetMessage(w, message, "ErrorCheckout")
		http.Redirect(w, r, "/checkout", http.StatusFound)
	}
	cart := models.Cart{}
	cart.FindByUserID(uint(userId))
	cart.LoadCartProducts()
	for i := 0; i < len(cart.CartProducts); i++ {
		orderProduct := models.OrderProduct{
			OrderID:   order.ID,
			ProductID: cart.CartProducts[i].ProductID,
			Quantity:  cart.CartProducts[i].Quantity,
			Price:     cart.CartProducts[i].PriceFollowQuantity(),
		}
		/* Create orderProduct */
		if  service.TransactionCreateOrderProduct(&orderProduct, tx) != nil {
			cookie.SetMessage(w, message, "ErrorCheckout")
			tx.Rollback()
			http.Redirect(w, r, "/checkout", http.StatusFound)
		}
		/* Delete cartProduct */
		
		if service.TransactionDeleteCartProduct(&cart.CartProducts[i], tx) != nil {
			cookie.SetMessage(w, message, "ErrorCheckout")
			tx.Rollback()
			http.Redirect(w, r, "/checkout", http.StatusFound)
		}
	}
	/* Delete cart */
	
	if service.TransactionDeleteCart(&cart, tx) != nil {
		cookie.SetMessage(w, message, "ErrorCheckout")
		tx.Rollback()
		http.Redirect(w, r, "/checkout", http.StatusFound)
	}
	/* Create payment */
	payment := models.Payment{
		OrderID : order.ID,
		AccountNumber : strings.Trim(r.FormValue("card_number"), " "),
		Bank : strings.Trim(r.FormValue("bank"), " "),
	}
	if service.TransactionCreatePayment(&payment, tx) != nil {
		cookie.SetMessage(w, message, "ErrorCheckout")
		tx.Rollback()
		http.Redirect(w, r, "/checkout", http.StatusFound)
	}
	message = "Order successful! Thank you!"
	cookie.SetMessage(w, message, "SuccessCheckout")
	tx.Commit()
	http.Redirect(w, r, "/checkout", http.StatusSeeOther)
}
