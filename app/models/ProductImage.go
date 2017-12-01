package models

import (
	"github.com/jinzhu/gorm"
)

type ProductImage struct {
	gorm.Model
	ProductID uint `gorm:"index"`
	Image string
}

