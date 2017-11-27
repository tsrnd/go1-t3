package controller

import "net/http"
import "github.com/tsrnd/goweb3/app/shared/view"
import "github.com/tsrnd/goweb3/app/models"
import "log"

func HelloWorld(w http.ResponseWriter, r *http.Request){
    user := models.User{}
    err := user.FindByName("duy")
    if err != nil {
        log.Fatalln(err)
    }
    v := view.New(r)
    v.Name = "home/index"
    v.Render(w)  
}