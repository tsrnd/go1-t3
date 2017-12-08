package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateOdersTable_20171205_150922 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateOdersTable_20171205_150922{}
	m.Created = "20171205_150922"

	migration.Register("CreateOdersTable_20171205_150922", m)
}

// Run the migrations
func (m *CreateOdersTable_20171205_150922) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable("oders","InnoDB","utf8");
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(11)").SetUnsigned(true)
	m.NewCol("user_id").SetDataType("INT(11)").SetNullable(false).SetUnsigned(true)
	m.NewCol("name_receiver").SetDataType("text").SetNullable(false)
	m.NewCol("address").SetDataType("text").SetNullable(false)
	m.NewCol("Status").SetDataType("INT(1)").SetNullable(false).SetUnsigned(true)
}

// Reverse the migrations
func (m *CreateOdersTable_20171205_150922) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE oders")
}
