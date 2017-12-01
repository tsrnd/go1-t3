package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/jinzhu/gorm"
	
)
type Product struct {
	gorm.Model
	Name string	     		
	Description string
	Quantity int
	Price int
	ProductImages []ProductImage
}

/**
*
*
**/
func (product *Product) GetTopProducts() (products []Product, err error) {
	err = database.SQL.Limit(9).Preload("ProductImages").Find(&products).Error
	return
}
