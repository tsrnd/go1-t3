package middleware

import (
	"strings"
	"net/http"
	"github.com/goweb3/app/validations"
	"github.com/goweb3/app/shared/cookie"
)
// ValidateLoginFormMiddleware check validate login form
func ValidateLoginFormMiddleware() Middleware {
	
		// Create a new Middleware
		return func(f http.HandlerFunc) http.HandlerFunc {
	
			// Define the http.HandlerFunc
			return func(w http.ResponseWriter, r *http.Request) {
				// Do middleware things
				r.ParseForm()
				email := strings.Trim(r.FormValue("email"), " ")
				password := strings.Trim(r.FormValue("password"), " ")
				loginRequest := validations.LoginRequest {
					Email : email,
					Password : password,
				}
				err := loginRequest.ValidateStruct()
				message := make(map[string] string)
				if (err != nil) {	
					message = validations.CustomErrorMessage(err)
					for key, val := range message {
						cookie.SetMessage(w, val, "Login" + key)
					}
					http.Redirect(w, r, "/login", http.StatusFound)
					return					
				}
				f(w, r)
			}
		}
	}
	