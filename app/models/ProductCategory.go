package models

import (
	"github.com/jinzhu/gorm"	
)

type ProductCategory struct {
	gorm.Model
	ProductID uint		`schema:"product_id"`
	CategoryID uint		`schema:"category_id"`
}
