package models

import (
	"github.com/jinzhu/gorm"
)

type OrderProduct struct {
	gorm.Model
	OrderID uint
	ProductID uint
	Quantity uint
	Price uint
}
