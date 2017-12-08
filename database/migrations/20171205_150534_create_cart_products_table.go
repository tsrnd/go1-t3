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
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable("cart_products","InnoDB","utf8");
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(11)").SetUnsigned(true)
	m.NewCol("cart_id").SetDataType("INT(11)").SetNullable(false).SetUnsigned(true)
	m.NewCol("product_id").SetDataType("INT(11)").SetNullable(false).SetUnsigned(true)
	m.NewCol("quantity").SetDataType("INT(11)").SetNullable(false).SetUnsigned(true)
}

// Reverse the migrations
func (m *CreateCartProductsTable_20171205_150534) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE cart_products")
}
