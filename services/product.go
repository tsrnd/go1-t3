package services

import (
	model "github.com/goweb3/models"
	"github.com/astaxie/beego/orm"
)

/**
*
*
**/
func ProcessHompage(data map[interface{}] interface{}) (err error) {
	o := orm.NewOrm()
	product := new(model.Product)
	var products []*model.Product
	qs := o.QueryTable(product)
	qs.Limit(10).RelatedSel().All(&products)
	data["products"] = products
	return
}