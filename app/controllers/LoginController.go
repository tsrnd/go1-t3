package controller

import "net/http"
import "github.com/goweb3/app/shared/view"

func Login(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "auth/login"
	v.Render(w)
}
