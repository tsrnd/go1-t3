package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateCategoryTable_20171205_145407 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCategoryTable_20171205_145407{}
	m.Created = "20171205_145407"

	migration.Register("CreateCategoryTable_20171205_145407", m)
}

// Run the migrations
func (m *CreateCategoryTable_20171205_145407) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable("categories","InnoDB","utf8");
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(11)").SetUnsigned(true)
	m.NewCol("name").SetDataType("VARCHAR(45) COLLATE utf8_unicode_ci").SetNullable(false)
}

// Reverse the migrations
func (m *CreateCategoryTable_20171205_145407) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE categories")
}
