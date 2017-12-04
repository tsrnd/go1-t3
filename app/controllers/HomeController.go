package controller

import (
	"log"
	"net/http"

	service "github.com/goweb3/app/services"
	"github.com/goweb3/app/shared/view"
)

func Home(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	err := service.ProcessHompage(r, v.Vars)
	if err != nil {
		log.Fatal(err.Error())
	}
	v.Name = "home/index"
	v.Render(w)
}
