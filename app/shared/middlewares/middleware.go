package middleware

import (
	"net/http"
	"github.com/goweb3/app/controllers"
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
			if len(r.Form["email"][0])==0 {
				controller.Login(w, r)
				return
			}
			user := models.User{}	
			exist, err := user.CheckExistEmail(r.Form["email"][0])
			if  exist || err != nil {
				controller.Login(w, r)
				return
			}
			if len(r.Form["password"][0])==0{
				controller.Login(w, r)
				return
			}

			// for k, v := range r.Form {
			// 	fmt.Println("key:", k)
			// 	fmt.Println("Typeval:", reflect.TypeOf(v))
			// 	fmt.Println(v[0])
			// }
			// Call the next middleware/handler in chain
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