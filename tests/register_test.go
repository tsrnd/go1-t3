package tests

import (
	"github.com/goweb3/app/models"
	"os"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"github.com/manveru/faker"
	controller "github.com/goweb3/app/controllers"
	middlewares "github.com/goweb3/app/middlewares"	
)
var fake, err = faker.New("en")

func TestMain(m *testing.M) {
	user := models.User{
		Name: "duy",
		Email: "duy@gmail.com",
		Password: "123456",
	}
	if user.FindByEmail(user.Email) != nil {
		user.Create()
	} 
	code := m.Run()
	os.Exit(code)
}

func makeRequestRegister(username string, email string, password string) (req *http.Request, err error) {
	form := url.Values{}
	form.Add("name", username)
	form.Add("email", email)
	form.Add("password", password)
	req, err = http.NewRequest("POST", "/regitser", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return
}

func TestRegisterSuccess(t *testing.T) {
	rr := httptest.NewRecorder()
	email := fake.Email()
	name := fake.Characters(5)
	req, err := makeRequestRegister(name, email, "123456")
	if err != nil {
		t.Fatal(err)
	}

	userController := &controller.UserController{Render: render}
	handler := http.HandlerFunc(middlewares.Chain(userController.Create, middlewares.ValidateRegisterFormMiddleware()))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != 303 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, 303)
	}
	if rr.Header()["Location"][0] != "/login" {
			t.Errorf("Redirect failed for successful register!")
	}
	user := models.User{}
	if user.FindByEmail(email) != nil {
		t.Errorf("handler cannot register account")
	} 
}

func TestRegisterFailValidation(t *testing.T) {
	ValidationRegister(t, "", fake.Email(), "123456", "RegisterUsernameREQUIRED")
	ValidationRegister(t, "ht", fake.Email(), "123456", "RegisterUsernameMIN")
	ValidationRegister(t, fake.Characters(41), fake.Email(), "123456", "RegisterUsernameMAX")
	ValidationRegister(t, "***", fake.Email(), "123456", "RegisterUsernameALPHANUM")
	ValidationRegister(t, fake.Characters(5), "", "123456", "RegisterEmailREQUIRED")
	ValidationRegister(t, fake.Characters(5), "duy@gmail.com", "123456", "RegisterEmailEXIST")
	ValidationRegister(t, fake.Characters(5), fake.Characters(5), "123456", "RegisterEmailEMAIL")
	ValidationRegister(t, fake.Characters(5), fake.Email(), "", "RegisterPasswordREQUIRED")
	ValidationRegister(t, fake.Characters(5), fake.Email(), "123", "RegisterPasswordMIN")
}

func ValidationRegister(t *testing.T, name string, email string, password string, expectMessage string) {
	rr := httptest.NewRecorder()
	req, err := makeRequestRegister(name, email, password)
	if err != nil {
		t.Fatal(err)
	}
	userController := &controller.UserController{Render: render}
	handler := http.HandlerFunc(middlewares.Chain(userController.Create, middlewares.ValidateRegisterFormMiddleware()))
	handler.ServeHTTP(rr, req)
	if rr.Header()["Location"][0] != "/login" {
		t.Errorf("Redirect failed for register!")
	}
	pass := false
	for _, val := range rr.Header()["Set-Cookie"] {
		if (strings.Contains(val, expectMessage)) {
			pass = true
			break
		}
	}
	if !pass {
		t.Errorf("Validate failed for %v", expectMessage)		
	}
}
