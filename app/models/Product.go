package models

import (
	"github.com/jinzhu/gorm"
)
type Product struct {
	gorm.Model
	Name string	     		
	Description string
	Quantity int
	Price int
}
