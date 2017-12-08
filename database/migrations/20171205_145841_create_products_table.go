package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateProductsTable_20171205_145841 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateProductsTable_20171205_145841{}
	m.Created = "20171205_145841"

	migration.Register("CreateProductsTable_20171205_145841", m)
}

// Run the migrations
func (m *CreateProductsTable_20171205_145841) Up() {
	m.SQL(`
		CREATE TABLE products
		(
			id SERIAL,
			name character varying(45) COLLATE pg_catalog."default" NOT NULL,
			description text COLLATE pg_catalog."default",
			quantity integer DEFAULT 0,
			price integer DEFAULT 0,
			created_at timestamp(0) without time zone DEFAULT now(),
			updated_at timestamp(0) without time zone DEFAULT now(),
			deleted_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
			CONSTRAINT products_pkey PRIMARY KEY (id)
		)
	`)
}

// Reverse the migrations
func (m *CreateProductsTable_20171205_145841) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE products")
}
