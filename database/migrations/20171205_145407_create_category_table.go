package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateCategoryTable_20171205_145407 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCategoryTable_20171205_145407{}
	m.Created = "20171205_145407"

	migration.Register("CreateCategoryTable_20171205_145407", m)
}

// Run the migrations
func (m *CreateCategoryTable_20171205_145407) Up() {
	m.SQL(`
		CREATE TABLE categories
		(
			id SERIAL,
			name character varying(45) COLLATE pg_catalog."default" NOT NULL,
			created_at timestamp(0) without time zone DEFAULT now(),
			updated_at timestamp(0) without time zone DEFAULT now(),
			deleted_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
			CONSTRAINT categories_pkey PRIMARY KEY (id)
		)
	`)
}

// Reverse the migrations
func (m *CreateCategoryTable_20171205_145407) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE categories")
}
