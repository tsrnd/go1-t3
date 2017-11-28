package controller

import "net/http"
import "github.com/goweb3/app/shared/view"

func News(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "news/index"
	v.Render(w)
}