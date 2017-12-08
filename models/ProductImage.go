package models

import (
	"github.com/astaxie/beego/orm"
)

type ProductImage struct {
	Id uint
	Product  *Product `orm:"rel(fk)"`
	Image string
}
func init() {
    orm.RegisterModel(new(ProductImage))
}

func (p *ProductImage) TableName() string {
    return "product_images"
}

