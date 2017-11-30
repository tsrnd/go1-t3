package service

import (
	model "github.com/goweb3/app/models"
	"net/http"
	"github.com/goweb3/app/shared/passhash"
	"github.com/jianfengye/web-golang/web/session"
	"strconv"
	"github.com/goweb3/app/shared/flash"
)

/**
* Authentication function
*
* return error
**/
func Auth(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")
	var err error
	user := model.User{Email: email, Password: password}
	err = user.FindByEmail(email)
	sess,_ := session.SessionStart(r, w)
	if (err == nil && passhash.MatchString(user.Password, password)) {
		// Login successfully
		flash.SetFlash(w, "success", []byte("Login success!"))
		sess.Set("id", strconv.Itoa(user.Id))
		sess.Set("email", user.Email)
		sess.Set("name", user.Name)
		return nil	
	}
	flash.SetFlash(w, "error", []byte("Login fail!"))
	return err
}