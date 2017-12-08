package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateOderProductsTable_20171205_151232 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateOderProductsTable_20171205_151232{}
	m.Created = "20171205_151232"

	migration.Register("CreateOderProductsTable_20171205_151232", m)
}

// Run the migrations
func (m *CreateOderProductsTable_20171205_151232) Up() {
	m.SQL(`
		CREATE TABLE order_products
		(
			id SERIAL,
			order_id integer NOT NULL REFERENCES orders(id),
			product_id integer REFERENCES products(id),
			quantity integer,
			price integer,
			created_at timestamp(0) without time zone DEFAULT now(),
			updated_at timestamp(0) without time zone DEFAULT now(),
			create_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
			deleted_at timestamp without time zone,
			CONSTRAINT order_items_pkey PRIMARY KEY (id)
		)
	`)
}

// Reverse the migrations
func (m *CreateOderProductsTable_20171205_151232) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE oder_products")
}
