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

	var data = &productData{
		Name:  product.Name,
		Id:    product.Id,
		Price: product.Price,
	}

	fmt.Println(data)
	this.Data["product"] = data

	this.TplName = "product/index.html"
}

type productData struct {
	Name  string
	Id    uint
	Price int
}
