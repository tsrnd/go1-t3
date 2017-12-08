package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateProductCategoriesTable_20171205_150246 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateProductCategoriesTable_20171205_150246{}
	m.Created = "20171205_150246"

	migration.Register("CreateProductCategoriesTable_20171205_150246", m)
}

// Run the migrations
func (m *CreateProductCategoriesTable_20171205_150246) Up() {
	m.SQL(`
		CREATE TABLE product_categories
		(
			id SERIAL,
			product_id integer NOT NULL REFERENCES products(id),
			category_id integer NOT NULL REFERENCES categories(id),
			created_at timestamp(0) without time zone DEFAULT now(),
			updated_at timestamp(0) without time zone DEFAULT now(),
			deleted_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
			CONSTRAINT product_categories_pkey PRIMARY KEY (id)
		)
	`)
}

// Reverse the migrations
func (m *CreateProductCategoriesTable_20171205_150246) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE product_categories")
}
