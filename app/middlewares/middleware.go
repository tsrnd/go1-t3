package middleware

import (
	"fmt"
	"strings"
	"net/http"
	"github.com/goweb3/app/models"
	
)



type Middleware func(http.HandlerFunc) http.HandlerFunc

// ValidateRegisterFormMiddleware check user loginned or not
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
			message := make(map[string] string)
			
			if len(email)==0 {
				message["email-required"] = "Email field is required"
			} else {
				user := models.User{}	
				err := user.CheckExistEmail(email)
				if  err == nil {
					message["email-exist"] = "Email already exists"
				}
			}
			if len(name) == 0 {
				message["name-required"] = "Name field is required"
			}
			if len(password) == 0 {
				message["password-required"] = "Password field is required"
			} else if len(password) <= 4 {
				message["password-min"] = "Password must be greater than 4 characters"				
			} else if len(password) > 20 {
				message["password-max"] = "Password must be less than 20 characters"				
			}
			if (len(message) != 0) {		
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