package controller

import (
	"net/http"

	"github.com/goweb3/app/shared/view"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	v := view.New(w, r)
	v.Name = "home/index"
	v.Render(w)
}
