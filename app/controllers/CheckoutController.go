package controller

import "net/http"
import "github.com/goweb3/app/shared/view"

func Checkout(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "checkout/index"
	v.Render(w)
}
