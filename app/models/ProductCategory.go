package models

import (
	"github.com/jinzhu/gorm"	
)

type ProductCategory struct {
	gorm.Model
	ProductID uint
	CategoryID uint
}
