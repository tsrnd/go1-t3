package controller

import (
	"net/http"
	"strings"

	"github.com/goweb3/app/models"
)

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := models.User{
		Name:     strings.Trim(r.Form["name"][0], " "),
		Email:    strings.Trim(r.Form["email"][0], " "),
		Password: strings.Trim(r.Form["password"][0], " "),
	}
	message, statusCode := make([]string, 0), http.StatusOK
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

var GetUserController = &UserController{Render: renderView}
