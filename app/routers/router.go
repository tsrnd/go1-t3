package router

import "github.com/gorilla/mux"
import controler "github.com/goweb3/app/controllers"

func Init() *mux.Router {
	r := mux.NewRouter()
	// Serve static files, no directory browsing
	r.PathPrefix("/assets/").HandlerFunc(controler.Static)
	// r.HandleFunc("/assets/**", controler.Static)
	r.HandleFunc("/", controler.HelloWorld).Methods("GET")
	return r
}
