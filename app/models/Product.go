package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name          string `schema:"name"`
	Description   string `schema:"description"`
	Quantity      int    `schema:"quantity"`
	Price         int    `schema:"price"`
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

/**
*	Find product by product list id
**/
func (product *Product) FindByListID(id uint) error {
	var err error
	err = database.SQL.Where("id = ?", id).First(&product).Error
	return err
}
