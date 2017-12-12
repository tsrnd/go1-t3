package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/goweb3/app/shared/view"
)

func (this *Product) Index(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Vars["products"] = (&model.Product{}.GetAll(10))
}

func (this *Product) Show(w http.ResponseWriter, r *http.Request) {
	product_id, _ = strconv.Atoi(mux.Vars["id"])
	v := view.New(r)
	product = &model.Product{}
	v.Vars["product"] = product.FindById(product_id)
	v.Name = "product/index"
	this.Render(w, v)
}

var GetProductController = &ProductConroller{Render: renderView}
