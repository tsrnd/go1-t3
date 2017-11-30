package controller

import "net/http"
import "github.com/gorilla/csrf"
import "github.com/goweb3/app/shared/view"
import service "github.com/goweb3/app/services"
import "github.com/goweb3/app/shared/flash"	

/**
*
* Get view Login
**/
func Login(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Vars[csrf.TemplateTag] = csrf.TemplateField(r)
	v.Name = "auth/login"
	emailRequired,_ := flash.GetFlash(w, r, "email-required")
	emailExist,_ := flash.GetFlash(w, r, "email-exist")
	nameRequired,_ := flash.GetFlash(w, r, "name-required")
	passwordRequired,_ := flash.GetFlash(w, r, "password-required")
	passwordMin,_ := flash.GetFlash(w, r, "password-min")
	passwordMax,_ := flash.GetFlash(w, r, "password-max")
	v.Vars["emailRequired"] = string(emailRequired)
	v.Vars["emailExist"] = string(emailExist)
	v.Vars["nameRequired"] = string(nameRequired)
	v.Vars["passwordRequired"] = string(passwordRequired)
	v.Vars["passwordMin"] = string(passwordMin)
	v.Vars["passwordMax"] = string(passwordMax)
	v.Render(w)
}

/**
*
* Post Login
**/
func LoginPost(w http.ResponseWriter, r *http.Request) {
	err := service.Auth(w, r)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		Login(w, r)
	}
}
