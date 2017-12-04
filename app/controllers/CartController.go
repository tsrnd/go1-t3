package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	productID, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := service.ProcessAddToCard(w, r, uint(productID))
	if err == nil {
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
