package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	controller "github.com/goweb3/app/controllers"
)

func init() {
	Init()
}

func makeRequestProductDetail() (req *http.Request, err error) {
	req, err = http.NewRequest("GET", "/products/1", nil)
	return
}
func TestProductController_Show(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := makeRequestProductDetail()
	if err != nil {
		t.Fatal(err)
	}
	productController := &controller.ProductController{Render: render}
	handler := http.HandlerFunc(productController.Show)
	handler.ServeHTTP(rr, req)
	fmt.Println(productController)
	if status := rr.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, 200)
	}
}
