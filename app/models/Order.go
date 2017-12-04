package models

import (
	"github.com/jinzhu/gorm"
)
type Order struct {
	gorm.Model
	UserID uint
	NameReceiver string
	Address string
	Status uint
}
