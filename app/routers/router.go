package router

import "github.com/gorilla/mux"

import controler "github.com/goweb3/app/controllers"

func Init() *mux.Router {
	r := mux.NewRouter()
	// Serve static files, no directory browsing
	r.PathPrefix("/assets/").HandlerFunc(controler.Static)
	// r.HandleFunc("/assets/**", controler.Static)
	r.HandleFunc("/", controler.HelloWorld).Methods("GET")
	r.HandleFunc("/login", controler.Login).Methods("GET")
	r.HandleFunc("/news", controler.News).Methods("GET")
	r.HandleFunc("/contact", controler.Contact).Methods("GET")
	r.HandleFunc("/shoe", controler.Login).Methods("GET")
	r.HandleFunc("/cart", controler.Cart).Methods("GET")
	r.HandleFunc("/checkout", controler.Checkout).Methods("GET")
	return r
}
