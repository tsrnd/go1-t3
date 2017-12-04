package models

import (
	"github.com/jinzhu/gorm"
)
type Product struct {
	gorm.Model
	Name string	 		`schema:"name"`	
	Description string	`schema:"description"`
	Quantity int		`schema:"quantity"`
	Price int			`schema:"price"`
	ProductImages []ProductImage
}
