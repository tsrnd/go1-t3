package controller

import "net/http"
import "github.com/goweb3/app/models"
import "time"
import "strings"

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := models.User{
		0,
		strings.Trim(r.Form["name"][0], " "),
		strings.Trim(r.Form["email"][0], " "),
		strings.Trim(r.Form["password"][0], " "),
		time.Time{},
		time.Time{},
		time.Time{},
	}
	message, statusCode := make([] string, 0), http.StatusOK
	err := user.HashPassword()
	if err == nil {
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
