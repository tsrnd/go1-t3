package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/astaxie/beego/orm"
)

type Product struct {
	Id uint
	Name string	 		
	Description string
	Quantity int
	Price int
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
func (product *Product) FindByID(id uint) error {
	var err error
	err = database.SQL.Where("id = ?", id).First(&product).Error
	return err
}
