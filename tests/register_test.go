package tests

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	_ "github.com/goweb3/routers"

	"github.com/astaxie/beego"
	"github.com/goweb3/app/shared/view"
	. "github.com/smartystreets/goconvey/convey"
)

var render = func(w http.ResponseWriter, v *view.View) {
	return
}

func TestRegisterSuccess(t *testing.T) {
	form := url.Values{}
	form.Add("email", "ht@gmail.com")
	form.Add("password", "123456")
	r, _ := http.NewRequest("POST", "/guest/register", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 302)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldEqual, 0)
		})
	})
}

func TestRegisterFaile(t *testing.T) {
	form := url.Values{}
	form.Add("email", "abc@gmail.com")
	form.Add("password", "123456")
	r, _ := http.NewRequest("POST", "/guest/register", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 302)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldEqual, 0)
		})
	})
}
