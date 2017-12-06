package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateCartsTable_20171205_150424 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCartsTable_20171205_150424{}
	m.Created = "20171205_150424"

	migration.Register("CreateCartsTable_20171205_150424", m)
}

// Run the migrations
func (m *CreateCartsTable_20171205_150424) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable("carts","InnoDB","utf8");
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(11)").SetUnsigned(true)
	m.NewCol("user_id").SetDataType("INT(11)").SetNullable(false).SetUnsigned(true)
}

// Reverse the migrations
func (m *CreateCartsTable_20171205_150424) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE carts")
}
