package main

import (
	"log"
	"net/http"
	"github.com/goweb3/config"
)

func main() {
	db := config.DB()
	cache := config.Cache()
	router := config.Router(db, cache)
	port := config.Port()
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
