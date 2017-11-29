package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/goweb3/app/models"
	service "github.com/goweb3/app/services"
	"github.com/goweb3/app/shared/view"
)

/**
* Get cart
*
* return cart view
**/
func Cart(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "cart/index"
	v.Render(w)
}

/**
* Add product to cart function
*
* return cart
**/
func AddToCart(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)
	productID, _ := strconv.Atoi(data["id"])
	product := models.Product{}
	_ = product.FindById(productID)
	cart := models.Cart{}
	err := cart.CheckExistCart(1)
	if err != nil { // cart not exist
		cart = models.Cart{
			ID:        0,
			UserID:    1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: time.Time{},
		}
		_ = cart.Create()
		err = service.ProcessCheckExistCartProduct(cart.ID, product.ID)
	} else { // cart exist
		err = service.ProcessCheckExistCartProduct(cart.ID, product.ID)
	}
	http.Redirect(w, r, "/cart", http.StatusOK)
}
