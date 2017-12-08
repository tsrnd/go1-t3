package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateOdersTable_20171205_150922 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateOdersTable_20171205_150922{}
	m.Created = "20171205_150922"

	migration.Register("CreateOdersTable_20171205_150922", m)
}

// Run the migrations
func (m *CreateOdersTable_20171205_150922) Up() {
	m.SQL(`
		CREATE TABLE orders
		(
			id SERIAL,
			user_id integer NOT NULL REFERENCES users(id),
			name_receiver character varying(45) COLLATE pg_catalog."default" DEFAULT NULL::character varying,
			address character varying(45) COLLATE pg_catalog."default" DEFAULT NULL::character varying,
			status integer,
			created_at timestamp(0) without time zone DEFAULT now(),
			updated_at timestamp(0) without time zone DEFAULT now(),
			deleted_at timestamp without time zone,
			CONSTRAINT orders_pkey PRIMARY KEY (id)
		)
	`)
}

// Reverse the migrations
func (m *CreateOdersTable_20171205_150922) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE oders")
}
