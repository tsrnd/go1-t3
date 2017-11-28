package models

import (
	"time"
)
type CartProduct struct {
	Id int					`db:"id" bson:"id"`
	CartId int				`db:"cart_id" bson:"cart_id"`
	ProductId int			`db:"product_id" bson:"product_id"`
	Quantity int			`db:"quantity" bson:"quantity"`
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
	DeletedAt time.Time     `db:"deleted_at" bson:"deleted_at"`
}
