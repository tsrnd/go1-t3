package models

import (
	"time"
)
type Cart struct {
	Id int	     			`db:"id" bson:"id"`
	UserId int				`db:"user_id" bson:"user_id"`
	ItemId int     			`db:"item_id" bson:"item_id"`
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
	Deleted   time.Time     `db:"deleted" bson:"deleted"`
}