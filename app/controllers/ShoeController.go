package controller

import "net/http"
import "github.com/goweb3/app/shared/view"

func Shoe(w http.ResponseWriter, r *http.Request) {
	v := view.New(w, r)
	v.Name = "category/shoe/index"
	v.Render(w)
}
