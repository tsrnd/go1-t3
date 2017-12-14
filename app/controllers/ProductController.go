package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/goweb3/app/dto"
	"github.com/goweb3/app/models"
	"github.com/goweb3/app/shared/view"
)

type ProductController struct {
	Render render
}

func (this *ProductController) Index(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	products, _ := (&models.Product{}).GetAll()
	for _, product := range products {
		product.LoadProductImage()
	}
	v.Vars["products"] = products
	v.Name = "home/index"
	this.Render(w, v)
}

func (this *ProductController) Show(w http.ResponseWriter, r *http.Request) {
	product_id, _ := strconv.Atoi(mux.Vars(r)["id"])
	v := view.New(r)
	product := &models.Product{}
	err := product.FindByID(uint(product_id))
	if err != nil {
		Error404(w, r)
		return
	}
	product.LoadProductImage()
	data := &dto.DataTransferProduct{
		Id:         product.ID,
		Name:       product.Name,
		Descrition: product.Description,
		Quantity:   product.Quantity,
		Price:      product.Price,
		PrimaryImage: func() string {
			var s string = ""
			if product.ProductImages != nil {
				s = "/asset/images/product-details/" + product.ProductImages[0].Image
			}
			return s
		}(),
		Images: func() (images map[int][]string) {
			images = make(map[int][]string)
			for key, image := range product.ProductImages {
				var keymap = 0
				if key%3 == 0 && key > 0 {
					keymap++
				}
				images[keymap] = append(images[keymap], "/asset/images/product-details/"+image.Image)
			}
			return
		}(),
	}

	v.Vars["product"] = data
	v.Name = "product/index"
	this.Render(w, v)
}

var GetProductController = &ProductController{Render: renderView}
