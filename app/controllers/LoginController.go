package controller

import "net/http"
import "github.com/gorilla/csrf"
import "github.com/goweb3/app/shared/view"
import service "github.com/goweb3/app/services"
import "github.com/goweb3/app/shared/cookie"

type LoginController struct {
	Render render
}

/**
*
* Get view Login
**/
func (l *LoginController) Index(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Vars[csrf.TemplateTag] = csrf.TemplateField(r)
	message := cookie.GetMessageStartWith(w, r, "Register")
	for key, val := range cookie.GetMessageStartWith(w, r, "Login") {
		message[key] = val
	}
	v.Vars["message"] = message
	v.Name = "auth/login"
	l.Render(w, v)
}

/**
*
* Post Login
**/
func (l *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	err := service.Auth(w, r)
	if err == nil {
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 303)
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
