package models

import (
	"time"
)

type OrderProduct struct {
	Id int					`db:"id" bson:"id"`
	OrderId int				`db:"order_id" bson:"order_id"`
	ProductId int			`db:"product_id" bson:"product_id"`
	Quantity int			`db:"quantity" bson:"quantity"`
	Price int				`db:"price" bson:"price"`
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
	DeletedAt time.Time     `db:"deleted_at" bson:"deleted_at"`
}
