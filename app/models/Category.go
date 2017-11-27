package models

import (
	"time"
)
type Category struct {
	Id int					`db:"id" bson:"id"`
	Name string				`db:"name" bson:"name"`
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
	Deleted   time.Time     `db:"deleted" bson:"deleted"`
}