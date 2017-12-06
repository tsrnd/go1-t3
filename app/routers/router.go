package router

import "github.com/gorilla/mux"
import "github.com/gorilla/csrf"
import "net/http"
import controler "github.com/goweb3/app/controllers"
import middlewares "github.com/goweb3/app/middlewares"

// public Routes for test
func Routes() *mux.Router {
	r := mux.NewRouter()
	// Serve static files, no directory browsing
	r.PathPrefix("/assets/").HandlerFunc(controler.Static)
	r.HandleFunc("/", controler.Home).Methods("GET")
	r.HandleFunc("/login", controler.GetLoginController.Index).Methods("GET")
	r.HandleFunc("/news", controler.News).Methods("GET")
	r.HandleFunc("/contact", controler.Contact).Methods("GET")
	r.HandleFunc("/shoe", controler.Shoe).Methods("GET")
	r.HandleFunc("/register", middlewares.Chain(controler.GetUserController.Create, middlewares.ValidateRegisterFormMiddleware())).Methods("POST")
	r.HandleFunc("/login", controler.GetLoginController.Login).Methods("POST")
	r.HandleFunc("/logout", controler.Logout).Methods("GET")
	r.HandleFunc("/checkout", middlewares.Chain(controler.Checkout, middlewares.LoginMiddleware())).Methods("GET")
	r.HandleFunc("/checkout", middlewares.Chain(controler.CheckoutPost, middlewares.LoginMiddleware())).Methods("POST")

	// Cart
	r.HandleFunc("/cart", controler.Cart).Methods("GET")
	r.HandleFunc("/cart/{id:[0-9]+}", controler.AddToCart).Methods("GET")
	r.HandleFunc("/cart/del/{id:[0-9]+}", controler.DelCartProduct).Methods("GET")
	return r
}

func HTTP() http.Handler {
	return middleware(Routes())
}

func middleware(h http.Handler) http.Handler {
	CSRF := csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false), csrf.CookieName("_csrf"))
	return CSRF(h)
}
