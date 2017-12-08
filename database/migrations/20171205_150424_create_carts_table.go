package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateCartsTable_20171205_150424 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCartsTable_20171205_150424{}
	m.Created = "20171205_150424"

	migration.Register("CreateCartsTable_20171205_150424", m)
}

// Run the migrations
func (m *CreateCartsTable_20171205_150424) Up() {
	m.SQL(`
		CREATE TABLE carts
		(
			id SERIAL,
			user_id integer NOT NULL REFERENCES users(id),
			created_at timestamp(0) without time zone DEFAULT now(),
			updated_at timestamp(0) without time zone DEFAULT now(),
			deleted_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
			CONSTRAINT carts_pkey PRIMARY KEY (id)
		)
	`)
}

// Reverse the migrations
func (m *CreateCartsTable_20171205_150424) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE carts")
}
