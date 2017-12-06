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
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable("product_categories","InnoDB","utf8");
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(11)").SetUnsigned(true)
	m.NewCol("product_id").SetDataType("INT(11)").SetNullable(false).SetUnsigned(true)
	m.NewCol("category_id").SetDataType("INT(11)").SetNullable(false).SetUnsigned(true)
}

// Reverse the migrations
func (m *CreateProductCategoriesTable_20171205_150246) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE product_categories")
}
