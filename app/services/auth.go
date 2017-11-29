package service

import (
	model "github.com/goweb3/app/models"
	"github.com/goweb3/app/shared/session"
	"net/http"
	"github.com/goweb3/app/shared/passhash"
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
	sess := session.Instance(r)
	if (err == nil && passhash.MatchString(user.Password, password)) {
		// Login successfully
		session.Empty(sess)
		sess.Values["id"] = user.Id
		sess.Values["email"] = user.Email
		sess.Values["name"] = user.Name
		sess.Save(r, w)
		return nil		
	}
	return err
}