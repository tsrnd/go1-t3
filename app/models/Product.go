package models

import (
	"time"

	"github.com/goweb3/app/shared/database"
)

type Product struct {
	ID          int       `db:"id" bson:"id"`
	Name        string    `db:"name" bson:"name"`
	Description string    `db:"description" bson:"description"`
	Quantity    int       `db:"quantity" bson:"quantity"`
	Price       int       `db:"price" bson:"price"`
	CreatedAt   time.Time `db:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" bson:"updated_at"`
	DeletedAt   time.Time `db:"deleted_at" bson:"deleted_at"`
}

func (product *Product) FindById(id int) error {
	var err error
	err = database.SQL.QueryRow("SELECT id, name, description, quantity, price FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Description, &product.Quantity, &product.Price)
	return err
}
