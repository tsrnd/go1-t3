package controller

import "net/http"
import "strconv"
import "strings"
import "github.com/goweb3/app/shared/view"
import "github.com/gorilla/csrf"
import "github.com/goweb3/app/models"
import "github.com/jianfengye/web-golang/web/session"	
import "github.com/goweb3/app/shared/database"

func Checkout(w http.ResponseWriter, r *http.Request) {
	sess,_ := session.SessionStart(r, w)
	
	userId,_ := strconv.ParseInt(sess.Get("id"), 10, 32)

	cart := models.Cart{}
	cart.FindByUserId(int(userId))
	database.SQL.Model(&cart).Related(&cart.CartProducts)
	for i,_ := range cart.CartProducts {
		database.SQL.Model(&cart.CartProducts[i]).Related(&cart.CartProducts[i].Product)
		database.SQL.Model(&cart.CartProducts[i].Product).Related(&cart.CartProducts[i].Product.ProductImages)
		
	}
	v := view.New(r)
	v.Vars[csrf.TemplateTag] = csrf.TemplateField(r)
	v.Vars["cart"] = cart
	v.Vars["totalPrice"] = cart.TotalPrice()
	v.Name = "checkout/index"
	v.Render(w)
}


func CheckoutPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sess,_ := session.SessionStart(r, w)
	userId,_ := strconv.ParseInt(sess.Get("id"), 10, 32)
	order := models.Order{
		UserID: uint(userId),
		NameReceiver : strings.Trim(r.FormValue("name-receiver"), " "),
		Address: strings.Trim(r.FormValue("address"), " "),
		Status: 1,
	}
	message := make([] string, 0)
	/* Begin transaction */
	db := database.SQL.Begin()
	/* Create order */
	if err := db.Create(&order).Error; err != nil {
		db.Rollback()
		message = append(message, "Can not create order!")
		http.Redirect(w, r, "/checkout", http.StatusFound)
	}
	cart := models.Cart{}
	cart.FindByUserId(int(userId))
	db.Model(&cart).Related(&cart.CartProducts)
	for i := 0; i< len(cart.CartProducts); i++ {
		orderProduct := models.OrderProduct{
			OrderID : order.ID,
			ProductID : cart.CartProducts[i].ProductID,
			Quantity : cart.CartProducts[i].Quantity,
			Price : cart.CartProducts[i].PriceFollowQuantity(),
		}
		/* Create orderProduct */
		if err := db.Create(&orderProduct).Error; err != nil {
			db.Rollback()
			message = append(message, "Can not create order product!")
			http.Redirect(w, r, "/checkout", http.StatusFound)
		}
		/* Delete cartProduct */
		if err := db.Delete(&cart.CartProducts[i]).Error; err != nil {
			db.Rollback()
			message = append(message, "Can not delete cart product!")
			http.Redirect(w, r, "/checkout", http.StatusFound)
		}
	}
	/* Create cart */
	if err := db.Delete(&cart).Error; err != nil {
		db.Rollback()
		message = append(message, "Can not delete cart!")
		http.Redirect(w, r, "/checkout", http.StatusFound)
	}
	/* Create payment */
	payment := models.Payment{
		OrderID : order.ID,
		AccountNumber : strings.Trim(r.FormValue("car_number"), " "),
		Bank : strings.Trim(r.FormValue("bank"), " "),
	}
	if err := db.Create(&payment).Error; err != nil {
		db.Rollback()
		message = append(message, "Can not create payment!")
		http.Redirect(w, r, "/checkout", http.StatusFound)
	}
	db.Commit()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
