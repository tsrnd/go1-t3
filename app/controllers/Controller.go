package controller

import (
	"net/http"
)

type Controller interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}
