package controller

import "net/http"
import "github.com/goweb3/app/models"
import "strings"
import 	"github.com/goweb3/app/shared/flash"

	"github.com/goweb3/app/models"
)

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := models.User{
		Name:     strings.Trim(r.Form["name"][0], " "),
		Email:    strings.Trim(r.Form["email"][0], " "),
		Password: strings.Trim(r.Form["password"][0], " "),
	}

	statusCode := http.StatusOK

	err := user.HashPassword()
	if err == nil {
		flash.SetFlash(w, flash.Flash{"Password cannot hash!", flash.FlashError})
		statusCode = http.StatusFound
	}
	err = user.Create()
	if err != nil {

		flash.SetFlash(w, flash.Flash{"User cannot create. Please try again!", flash.FlashError})		
		statusCode = http.StatusFound		

	}
	http.Redirect(w, r, "/login", statusCode)
}

var GetUserController = &UserController{Render: renderView}
