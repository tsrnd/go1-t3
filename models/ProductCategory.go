package models

import (
	"github.com/astaxie/beego/orm"
)

type ProductCategory struct {
	Id uint
	Product *Product 	`orm:"column(product_id);rel(fk)"`
	Category *Category	`orm:"column(category_id);rel(fk)"`
}
func init() {
    orm.RegisterModel(new(ProductCategory))
}