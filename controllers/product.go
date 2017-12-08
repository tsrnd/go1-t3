package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/goweb3/models"
)

type ProductController struct {
	beego.Controller
}

func (this *ProductController) Index() {

}

func (this *ProductController) Show() {
	o := orm.NewOrm()
	var id, _ = this.GetUint32(":id")
	var uint_id = uint(id)
	product := &models.Product{Id: uint_id}
	o.Read(product)
	o.LoadRelated(product, "Images")
	var data = &productData{
		Name:         product.Name,
		Id:           product.Id,
		Price:        product.Price,
		Images:       getMapImage(product.Images...),
		PrimaryImage: "/static/images/product-details/" + product.Images[0].Image,
	}

	fmt.Println(data)

	this.Data["product"] = data

	this.TplName = "product/index.html"
}

func getMapImage(images ...*models.ProductImage) map[int][]string {
	var a = make(map[int][]string)
	var mapKey = 0
	for key, image := range images {
		if key%3 == 0 {
			mapKey = key / 3
		}
		a[mapKey] = append(a[mapKey], "/static/images/product-details/"+image.Image)
	}
	fmt.Println(a)
	return a
}

type productData struct {
	Name         string
	Id           uint
	Price        int
	Images       map[int][]string
	PrimaryImage string
}
