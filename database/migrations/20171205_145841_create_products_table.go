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
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable("products","InnoDB","utf8");
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(11)").SetUnsigned(true)
	m.NewCol("name").SetDataType("VARCHAR(45) COLLATE utf8_unicode_ci").SetNullable(false)
	m.NewCol("description").SetDataType("VARCHAR(45) COLLATE utf8_unicode_ci").SetNullable(true)
	m.NewCol("quantity").SetDataType("INT(11)").SetNullable(false)	
	m.NewCol("Price").SetDataType("INT(11)").SetNullable(false)
}

// Reverse the migrations
func (m *CreateProductsTable_20171205_145841) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE products")
}
