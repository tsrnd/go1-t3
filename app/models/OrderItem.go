package models

import (
	"time"
)

type OrderItem struct {
	Id int
	OrderId int
	ItemId int
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
	Deleted   time.Time     `db:"deleted" bson:"deleted"`
}