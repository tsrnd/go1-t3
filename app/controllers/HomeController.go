package controller

import "net/http"
import "github.com/goweb3/app/shared/view"
import service "github.com/goweb3/app/services"
import "fmt"

func Index(w http.ResponseWriter, r *http.Request) {
	products, err := service.ProcessHompage(r)
	if err != nil {
		fmt.Println(err)
	}
	v := view.New(r)
	v.Vars["products"] = products
	v.Name = "home/index"
	v.Render(w)
}
