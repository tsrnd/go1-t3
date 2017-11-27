package models

import (
	"time"
)
type CartItem struct {
	Id int     				`db:"id" bson:"id"`
	CartId int	     		`db:"cart_id" bson:"cart_id"`
	ItemId int      		`db:"item_id" bson:"item_id"`
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
	Deleted   time.Time     `db:"deleted" bson:"deleted"`
}