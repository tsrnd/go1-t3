package controller

import "net/http"
import "github.com/goweb3/app/models"
import "strings"

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := models.User{
		Name: strings.Trim(r.FormValue("name"), " "),
		Email: strings.Trim(r.FormValue("email"), " "),
		Password: strings.Trim(r.FormValue("password"), " "),
	}
	message, statusCode := make([] string, 0), http.StatusSeeOther
	err := user.HashPassword()
	if err != nil {
		message = append(message, "Password cannot hash!")
		statusCode = http.StatusFound
	}
	err = user.Create()
	if err != nil {
		message = append(message, "User cannot create. Please try again!")
		statusCode = http.StatusFound		
	}
	http.Redirect(w, r, "/login", statusCode)
}
