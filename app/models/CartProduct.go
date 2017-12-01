package models

import (
	"github.com/jinzhu/gorm"
)
type CartProduct struct {
	gorm.Model
	CartID uint `gorm:"index"`
	ProductID uint
	Quantity uint
}
