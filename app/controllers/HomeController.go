package controller

import "net/http"
import "github.com/goweb3/app/shared/view"


func HelloWorld(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "home/index"
	v.Render(w)
}
