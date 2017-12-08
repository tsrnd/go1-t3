package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateProductImagesTable_20171205_150729 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateProductImagesTable_20171205_150729{}
	m.Created = "20171205_150729"

	migration.Register("CreateProductImagesTable_20171205_150729", m)
}

// Run the migrations
func (m *CreateProductImagesTable_20171205_150729) Up() {
	m.SQL(`
		CREATE TABLE product_images
		(
			id SERIAL,
			product_id integer NOT NULL REFERENCES products(id),
			image character varying(45) COLLATE pg_catalog."default" DEFAULT NULL::character varying,
			created_at timestamp(0) without time zone DEFAULT now(),
			update_at timestamp(0) without time zone DEFAULT now(),
			deleted_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
			CONSTRAINT product_images_pkey PRIMARY KEY (id)
		)
	`)
}

// Reverse the migrations
func (m *CreateProductImagesTable_20171205_150729) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE product_images")
}
