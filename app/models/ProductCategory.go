package models

import (
	"time"
)

type ProductCategory struct {
	Id int					`db:"id" bson:"id"`
	ProductId int			`db:"product_id" bson:"product_id"`
	CategoryId int			`db:"category_id" bson:"category_id"`
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
	DeletedAt time.Time     `db:"deleted_at" bson:"deleted_at"`
}
