package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/goweb3/app/shared/view"

	service "github.com/goweb3/app/services"
)

/**
*
* Get view Login
**/
func (l *LoginController) Index(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Vars[csrf.TemplateTag] = csrf.TemplateField(r)
	v.Name = "auth/login"
	l.Render(w, v)
}

/**
*
* Post Login
**/
func (l *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	fmt.Println(email)
	err := service.Auth(w, r)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

/**
*
*
**/
func Logout(w http.ResponseWriter, r *http.Request) {
	service.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}

var GetLoginController = &LoginController{Render: renderView}
