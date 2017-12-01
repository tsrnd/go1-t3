package service

import (
	model "github.com/goweb3/app/models"
	"net/http"
	"github.com/goweb3/app/shared/passhash"
	"github.com/jianfengye/web-golang/web/session"
	"fmt"
	"github.com/goweb3/app/shared/flash"
	"time"
	"errors"
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
		flash.SetFlash(w, flash.Flash{"Login success!", flash.FlashSuccess})
		sess.Set("id", fmt.Sprint(user.ID))
		sess.Set("email", user.Email)
		sess.Set("name", user.Name)
		return nil	
	}
	flash.SetFlash(w, flash.Flash{"Login fail!", flash.FlashError})
	return errors.New("")
}

/**
*
*
**/
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie(session.CookieName)
	if cookie != nil {
		sessionid := cookie.Value
		session.Sessions[sessionid] = nil
		dc := &http.Cookie{Name: session.CookieName, MaxAge: -1, Expires: time.Unix(1, 0)}
		http.SetCookie(w, dc)
	}
}