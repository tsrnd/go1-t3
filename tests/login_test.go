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

func makeRequestLogin(username string, password string) (req *http.Request, err error) {
	form := url.Values{}
	form.Add("email", username)
	form.Add("password", password)
	req, err = http.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return
}

func TestLoginSuccess(t *testing.T) {
	rr := httptest.NewRecorder()

	req, err := makeRequestLogin("duy@gmail.com", "123456")
	if err != nil {
		t.Fatal(err)
	}

	loginController := &controller.LoginController{Render: render}
	handler := http.HandlerFunc(loginController.Login)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != 302 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, 302)
	}
}

func TestLoginFailUserName(t *testing.T) {
	rr := httptest.NewRecorder()

	req, err := makeRequestLogin("abc123123@gmail.com", "123456")
	if err != nil {
		t.Fatal(err)
	}

	loginController := &controller.LoginController{Render: render}
	handler := http.HandlerFunc(loginController.Login)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != 303 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, 303)
	}
}

func TestLoginFailPassword(t *testing.T) {
	rr := httptest.NewRecorder()

	req, err := makeRequestLogin("duy@gmail.com", "1234567")
	if err != nil {
		t.Fatal(err)
	}

	loginController := &controller.LoginController{Render: render}
	handler := http.HandlerFunc(loginController.Login)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != 303 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, 303)
	}
}

func TestLoginFailUserNameAndPassword(t *testing.T) {
	rr := httptest.NewRecorder()

	req, err := makeRequestLogin("abc123123@gmail.com", "1234567")
	if err != nil {
		t.Fatal(err)
	}

	loginController := &controller.LoginController{Render: render}
	handler := http.HandlerFunc(loginController.Login)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != 303 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, 303)
	}
}
