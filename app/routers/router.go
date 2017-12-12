package router

import "github.com/gorilla/mux"
import "github.com/gorilla/csrf"
import "net/http"
import controller "github.com/goweb3/app/controllers"
import middlewares "github.com/goweb3/app/middlewares"

// public Routes for test
func Routes() *mux.Router {
	r := mux.NewRouter()
	// Serve static files, no directory browsing
	r.PathPrefix("/assets/").HandlerFunc(controller.Static)
	r.HandleFunc("/", controller.Home).Methods("GET")
	r.HandleFunc("/login", controller.GetLoginController.Index).Methods("GET")
	r.HandleFunc("/news", controller.News).Methods("GET")
	r.HandleFunc("/contact", controller.Contact).Methods("GET")
	r.HandleFunc("/shoe", controller.Shoe).Methods("GET")
	r.HandleFunc("/register", middlewares.Chain(controller.GetUserController.Create, middlewares.ValidateRegisterFormMiddleware())).Methods("POST")
	r.HandleFunc("/login", controller.GetLoginController.Login).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("GET")
	r.HandleFunc("/checkout", middlewares.Chain(controller.Checkout, middlewares.LoginMiddleware())).Methods("GET")
	r.HandleFunc("/checkout", middlewares.Chain(controller.CheckoutPost, middlewares.LoginMiddleware())).Methods("POST")

	// Product
	r.HandleFunc("/products", controller.GetProductController.Index).Methods("GET")
	r.HandleFunc("/products/{id:[0-9]+}", controller.GetProductController.Show).Methods("GET")

	// Cart
	r.HandleFunc("/cart", controller.GetCartController.Index).Methods("GET")
	r.HandleFunc("/cart/{id:[0-9]+}", controller.GetCartController.Store).Methods("GET")
	r.HandleFunc("/cart/del/{id:[0-9]+}", controller.GetCartController.Destroy).Methods("GET")
	return r
}

func HTTP() http.Handler {
	return middleware(Routes())
}

func middleware(h http.Handler) http.Handler {
	CSRF := csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false), csrf.CookieName("_csrf"))
	return CSRF(h)
}
