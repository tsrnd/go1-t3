package controller

import "net/http"
import "github.com/goweb3/app/shared/view"

func Contact(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "contact/index"
	v.Render(w)
}
