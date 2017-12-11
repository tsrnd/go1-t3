package models

import (
	// "github.com/goweb3/app/shared/database"
)

type Product struct {
	BaseModel
	Name          string `db:"name"`
	Description   string `db:"description"`
	Quantity      int    `db:"quantity"`
	Price         int    `db:"price"`
	ProductImages []ProductImage
}

/**
*
*
**/
func (product *Product) GetTopProducts() (products []Product, err error) {
	// err = database.SQL.Limit(9).Preload("ProductImages").Find(&products).Error
	return
}

/**
*	Find product by product id
**/
func (product *Product) FindByID(id uint) error {
	var err error
	// err = database.SQL.Where("id = ?", id).First(&product).Error
	return err
}

/**
*	Find product by product list id
**/
func (product *Product) FindByListID(id uint) error {
	var err error
	// err = database.SQL.Where("id = ?", id).First(&product).Error
	return err
}
