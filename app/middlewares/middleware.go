package middleware

import (
	"fmt"
	"strings"
	"net/http"
	"github.com/goweb3/app/models"
	"github.com/jianfengye/web-golang/web/session"	
	"github.com/goweb3/app/shared/flash"	
)



type Middleware func(http.HandlerFunc) http.HandlerFunc

// ValidateRegisterFormMiddleware check validate register form
func ValidateRegisterFormMiddleware() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			// Do middleware things
			r.ParseForm()
			email := strings.Trim(r.Form["email"][0], " ")
			name := strings.Trim(r.Form["name"][0], " ")
			password := strings.Trim(r.Form["password"][0], " ")
			isPass := true			
			if len(email)==0 {
				flash.SetFlash(w, "email-required", []byte("Email field is required!"))	
				isPass = false			
			} else {
				user := models.User{}	
				err := user.FindByEmail(email)
				if  err == nil {
					flash.SetFlash(w, "email-exist", []byte("Email already exists!"))	
					isPass = false
				}
			}
			if len(name) == 0 {
				flash.SetFlash(w, "name-required", []byte("Name field is required!"))	
				isPass = false
			}
			if len(password) == 0 {
				flash.SetFlash(w, "password-required", []byte("Password field is required!"))	
				isPass = false
			} else if len(password) <= 4 {
				flash.SetFlash(w, "password-min", []byte("Password must be greater than 4 characters!"))	
				isPass = false
			} else if len(password) > 20 {
				flash.SetFlash(w, "password-max", []byte("Password must be less than 20 characters!"))	
				isPass = false
			}
			if (!isPass) {		
				http.Redirect(w, r, "/login", http.StatusFound)	
				return
			}
			
			f(w, r)
		}
	}
}

// LoginMiddleware check user loginned or not
func LoginMiddleware() Middleware {
	
		// Create a new Middleware
		return func(f http.HandlerFunc) http.HandlerFunc {
	
			// Define the http.HandlerFunc
			return func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("eeeeee")
				sess,_ := session.SessionStart(r, w)
				if sess.Get("id") == "" ||	sess.Get("name") == "" || sess.Get("email") == "" {
					flash.SetFlash(w, "warning", []byte("You are not logged in. Please login!"))
					http.Redirect(w, r, "/login", http.StatusFound)	
					return
				}
				
				f(w, r)
			}
		}
	}


// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}