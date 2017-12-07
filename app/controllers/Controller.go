package controller

import (
	"net/http"

	"github.com/goweb3/app/shared/view"
)

type Controller interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}

type render func(w http.ResponseWriter, v *view.View)

var renderView = func(w http.ResponseWriter, v *view.View) {
	v.Render(w)
}

type LoginController struct {
	Render render
}

type UserController struct {
	Render render
}
