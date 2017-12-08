package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreatePaymentsTable_20171205_151448 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreatePaymentsTable_20171205_151448{}
	m.Created = "20171205_151448"
	migration.Register("CreatePaymentsTable_20171205_151448", m)
}

// Run the migrations
func (m *CreatePaymentsTable_20171205_151448) Up() {
	m.SQL(`
		CREATE TABLE payments
		(
			id SERIAL,
			order_id integer NOT NULL REFERENCES orders(id),
			account_number character varying(45) COLLATE pg_catalog."default",
			bank character varying(45) COLLATE pg_catalog."default",
			created_at timestamp without time zone,
			updated_at timestamp without time zone,
			deleted_at timestamp without time zone,
			CONSTRAINT payments_pkey PRIMARY KEY (id)
		)
	`)
}

// Reverse the migrations
func (m *CreatePaymentsTable_20171205_151448) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE payments")
}
