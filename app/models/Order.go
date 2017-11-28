package models

import (
	"time"
)
type Order struct {
	Id int					`db:"id" bson:"id"`
	UserId int     			`db:"user_id" bson:"user_id"`
	NameReceiver string		`db:"name_receiver" bson:"name_receiver"`
	Address string			`db:"address" bson:"address"`
	Status int				`db:"status" bson:"status"`
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
	DeletedAt time.Time     `db:"deleted_at" bson:"deleted_at"`
}
