package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateCartProductsTable_20171205_150534 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCartProductsTable_20171205_150534{}
	m.Created = "20171205_150534"

	migration.Register("CreateCartProductsTable_20171205_150534", m)
}

// Run the migrations
func (m *CreateCartProductsTable_20171205_150534) Up() {
	m.SQL(`
		CREATE TABLE cart_products
		(
			id SERIAL,
			cart_id integer NOT NULL REFERENCES carts(id),
			product_id integer REFERENCES products(id),
			quantity integer,
			created_at timestamp(0) without time zone DEFAULT now(),
			updated_at timestamp(0) without time zone DEFAULT now(),
			deleted_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
			CONSTRAINT cart_items_pkey PRIMARY KEY (id)
		)
	`)
}

// Reverse the migrations
func (m *CreateCartProductsTable_20171205_150534) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE cart_products")
}
