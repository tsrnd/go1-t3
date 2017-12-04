package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/jinzhu/gorm"
	
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Quantity uint
	Price uint
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

/** 
*	Find product by product id
**/
func (product *Product) FindByID(id uint) error {
	var err error
	err = database.SQL.Where("id = ?", id).First(&product).Error
	return err
}
