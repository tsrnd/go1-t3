package service

import (
	model "github.com/goweb3/app/models"
	"net/http"

)

/**
*
*
**/
func ProcessHompage(r *http.Request,data map[string] interface{}) (err error) {

	var product = model.Product{}
	products, err := product.GetTopProducts()
	data["products"] = products
	return
}