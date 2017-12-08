package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateUserTable_20171205_114942 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUserTable_20171205_114942{}
	m.Created = "20171205_114942"

	migration.Register("CreateUserTable_20171205_114942", m)
}

// Run the migrations
func (m *CreateUserTable_20171205_114942) Up() {
	m.SQL(`
		CREATE TABLE users
		(
			id SERIAL,
			name character varying(255) COLLATE pg_catalog."default" NOT NULL,
			email character varying(45) COLLATE pg_catalog."default" NOT NULL,
			password character varying(128) COLLATE pg_catalog."default" NOT NULL,
			created_at timestamp(0) without time zone DEFAULT now(),
			updated_at timestamp(0) without time zone DEFAULT now(),
			deleted_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
			CONSTRAINT users_pkey PRIMARY KEY (id)
		)
	`)
}

// Reverse the migrations
func (m *CreateUserTable_20171205_114942) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE users")
}
