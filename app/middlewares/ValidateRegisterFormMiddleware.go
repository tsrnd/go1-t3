package middleware

import (
	"github.com/goweb3/app/models"
	"strings"
	"net/http"
	"github.com/goweb3/app/validations"
	"github.com/goweb3/app/shared/cookie"
)
// ValidateRegisterFormMiddleware check validate register form
func ValidateRegisterFormMiddleware() Middleware {
	
		// Create a new Middleware
		return func(f http.HandlerFunc) http.HandlerFunc {
	
			// Define the http.HandlerFunc
			return func(w http.ResponseWriter, r *http.Request) {
				// Do middleware things
				r.ParseForm()
				email := strings.Trim(r.FormValue("email"), " ")
				name := strings.Trim(r.FormValue("name"), " ")
				password := strings.Trim(r.FormValue("password"), " ")
				registerRequest := validations.RegisterRequest {
					Username : name,
					Email : email,
					Password : password,
				}
				err := registerRequest.ValidateStruct()
				message := make(map[string] string)
				if (err != nil) {	
					message = validations.CustomErrorMessage(err)
				}
				user := models.User{}
				err = user.FindByEmail(email)
				if  err == nil {
					message["EmailEXIST"] = "Email already exists"
				}
				for key, val := range message {
					cookie.SetMessage(w, val, "Register"+key)
				}
				if len(message) != 0 {
					http.Redirect(w, r, "/login", http.StatusFound)
					return					
				}
				f(w, r)
			}
		}
	}
	