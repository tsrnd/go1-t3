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
	err := user.HashPassword()
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return		
	}
	err = user.Create()	
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return		
	}
	http.Redirect(w, r, "/login", http.StatusOK)
}
