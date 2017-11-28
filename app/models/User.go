package models

import (
	"github.com/goweb3/app/shared/database"
	"time"
)

type User struct {
	Id int					`db:"id" bson:"id"`
	Name string				`db:"name" bson:"name"`
	Email string			`db:"email" bson:"email"`
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
	DeletedAt time.Time     `db:"deleted_at" bson:"deleted_at"`
}

func(user *User) FindByName(name string) (error) {
	var err error
	err = database.SQL.QueryRow("SELECT id, name, email FROM users WHERE name = $1", name).Scan(&user.Id, &user.Name, &user.Email)
	return err
}

func(user *User) Create() (err error) {
	statement := "insert into users (name, email) values ($1, $2) returning id"
	stmt, err := database.SQL.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.Name, user.Email).Scan(&user.Id)
	return
}
