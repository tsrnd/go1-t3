package services

import (
	"github.com/astaxie/beego/orm"
	model "github.com/goweb3/models"
)

/**
*
*
**/
func ProcessHompage(data map[interface{}]interface{}) (err error) {
	o := orm.NewOrm()
	product := new(model.Product)
	var products []*model.Product
	qs := o.QueryTable(product)
	qs.Limit(10).RelatedSel().All(&products)
	data["products"] = products
	for _, product = range products {
		_, err = o.LoadRelated(product, "ProductImages")
	}
	return
}
