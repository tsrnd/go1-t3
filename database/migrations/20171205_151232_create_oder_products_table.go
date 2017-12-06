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
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable("oder_products","InnoDB","utf8");
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(11)").SetUnsigned(true)
	m.NewCol("order_id").SetDataType("INT(11)").SetNullable(false).SetUnsigned(true)
	m.NewCol("product_id").SetDataType("INT(11)").SetNullable(false).SetUnsigned(true)	
	m.NewCol("quantity").SetDataType("INT(11)").SetNullable(false).SetUnsigned(true)
	m.NewCol("price").SetDataType("INT(11)").SetNullable(false).SetUnsigned(true)
}

// Reverse the migrations
func (m *CreateOderProductsTable_20171205_151232) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE oder_products")
}
