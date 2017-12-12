package controller

import (
	"log"
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
func (this *CartController) Index(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	err := service.ProcessCartPage(w, r, v.Vars)
	if err != nil {
		log.Fatal(err.Error())
	}
	v.Name = "cart/index"
	this.Render(w, v)
}

/**
* Add product to cart function
*
* return cart
**/
func (this *CartController) Store(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := service.ProcessAddToCard(w, r, uint(productID))
	if err == nil {
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (this *CartController) Destroy(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(mux.Vars(r)["id"])
	service.ProcessDelCartProduct(w, r, uint(productID))
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

var GetCartController = &CartController{Render: renderView}
