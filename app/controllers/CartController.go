package controller

import "net/http"
import "github.com/goweb3/app/shared/view"

func Cart(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "cart/index"
	v.Render(w)
}