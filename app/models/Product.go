package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Quantity    int
	Price       int
}

func (product *Product) FindById(id int) error {
	var err error
	err = database.SQL.QueryRow("SELECT id, name, description, quantity, price FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Description, &product.Quantity, &product.Price)
	return err
}
