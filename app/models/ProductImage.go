package models

import (
	"github.com/jinzhu/gorm"
)

type ProductImage struct {
	gorm.Model
	ProductID uint 		`schema:"product_id"`
	Image string		`schema:"image"`
}

