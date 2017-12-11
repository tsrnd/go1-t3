package controller

import "net/http"
import "github.com/goweb3/app/models"
import "strings"
import "github.com/goweb3/app/shared/flash"

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := models.User{
		Name: strings.Trim(r.FormValue("name"), " "),
		Email: strings.Trim(r.FormValue("email"), " "),
		Password: strings.Trim(r.FormValue("password"), " "),
	}
	err := user.HashPassword()
	if err == nil {
		flash.SetFlash(w, flash.Flash{"Password cannot hash!", flash.FlashError})
	}
	err = user.Create()
	if err != nil {

		flash.SetFlash(w, flash.Flash{"User cannot create. Please try again!", flash.FlashError})		
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

var GetUserController = &UserController{Render: renderView}
