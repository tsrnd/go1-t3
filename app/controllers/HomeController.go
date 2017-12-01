package controller

import "net/http"
import "github.com/goweb3/app/shared/view"
import service "github.com/goweb3/app/services"
import "fmt"

func Index(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	err := service.ProcessHompage(r, v.Vars)
	if err != nil {
		fmt.Println(err)
	}
	v.Name = "home/index"
	v.Render(w)
}
