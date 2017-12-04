package controller

import "net/http"
import "github.com/gorilla/csrf"
import "github.com/goweb3/app/shared/view"
import service "github.com/goweb3/app/services"

/**
*
* Get view Login
**/
func Login(w http.ResponseWriter, r *http.Request) {
	v := view.New(w, r)
	v.Vars[csrf.TemplateTag] = csrf.TemplateField(r)
	v.Name = "auth/login"
	v.Render(w)
}

/**
*
* Post Login
**/
func LoginPost(w http.ResponseWriter, r *http.Request) {
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
