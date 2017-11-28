package controller

import "net/http"
import "github.com/goweb3/app/shared/view"
import "fmt"

func LoginGet(w http.ResponseWriter, r *http.Request){
    v := view.New(r)
    v.Name = "login/login"
    v.Render(w)  
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
    
    // Form values
	email := r.FormValue("email")
    password := r.FormValue("password")
    fmt.Println(email)
    fmt.Println(password)
    v := view.New(r)
    v.Name = "login/login"
    v.Render(w)
}