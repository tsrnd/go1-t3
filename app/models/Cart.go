package models

import (
	"time"

	"github.com/goweb3/app/shared/database"
)

type Cart struct {
	ID        int       `db:"id" bson:"id"`
	UserID    int       `db:"user_id" bson:"user_id"`
	CreatedAt time.Time `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time `db:"updated_at" bson:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" bson:"deleted_at"`
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
