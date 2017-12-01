package middleware

import (
	"net/http"
	"github.com/jianfengye/web-golang/web/session"	
	"github.com/goweb3/app/shared/flash"	
)

// LoginMiddleware check user loginned or not
func LoginMiddleware() Middleware {
	
		// Create a new Middleware
		return func(f http.HandlerFunc) http.HandlerFunc {
	
			// Define the http.HandlerFunc
			return func(w http.ResponseWriter, r *http.Request) {
				sess,_ := session.SessionStart(r, w)
				if sess.Get("id") == "" ||	sess.Get("name") == "" || sess.Get("email") == "" {
					flash.SetFlash(w, flash.Flash{"You are not logged in. Please login!", flash.FlashWarning})
					http.Redirect(w, r, "/login", http.StatusFound)	
					return
				}
				
				f(w, r)
			}
		}
	}