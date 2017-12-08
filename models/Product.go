package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/goweb3/app/shared/database"
)

type Product struct {
	Id            uint
	Name          string
	Description   string
	Quantity      int
	Price         int
	ProductImages []*ProductImage `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Product))
}
func (p *Product) TableName() string {
	return "products"
}

/**
*
*
**/
func (product *Product) GetTopProducts() (products []Product, err error) {
	err = database.SQL.Limit(9).Preload("ProductImages").Find(&products).Error
	return
}

/**
*	Find product by product id
**/
func (p *Product) FindByID() (err error) {
	o := orm.NewOrm()
	err = o.Read(p)
	return
}
