package models

import (
	"time"
)
type Product struct {
	Id int	     			`db:"id" bson:"id"`
	Name string	     		`db:"name" bson:"name"`
	Description string	    `db:"description" bson:"description"`
	Quantity int	    	`db:"quantity" bson:"quantity"`
	Price int				`db:"price" bson:"price"`	
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
	DeletedAt time.Time     `db:"deleted_at" bson:"deleted_at"`
}
