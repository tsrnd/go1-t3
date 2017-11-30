package service

import (
	model "github.com/goweb3/app/models"
	"net/http"
)

/**
*
*
**/
func ProcessHompage(r *http.Request) (products []model.Product, err error){
	var product = model.Product{}
	products, err = product.Limit(10)
	return 
}