package tests

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	controller "github.com/goweb3/app/controllers"
	"github.com/goweb3/app/shared/view"
)

var render = func(w http.ResponseWriter, v *view.View) {
	return
}

func TestLoginSuccess(t *testing.T) {
	form := url.Values{}
	form.Add("email", "duy@gmail.com")
	form.Add("password", "123456")
	req, err := http.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	loginController := &controller.LoginController{Render: render}
	handler := http.HandlerFunc(loginController.Login)
	handler.ServeHTTP(rr, req)
}
