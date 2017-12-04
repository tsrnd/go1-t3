package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/jinzhu/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint `gorm:"index"`
}

func (cart *Cart) Create() (err error) {
	statement := "insert into carts (user_id) values ($1) returning id"
	stmt, err := database.SQL.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(cart.UserID).Scan(&cart.ID)
	return
}

func (cart *Cart) CheckExistCart(userID int) error {
	var err error
	err = database.SQL.QueryRow("SELECT id, user_id FROM carts WHERE user_id = $1", userID).Scan(&cart.ID, &cart.UserID)
	return err
}
