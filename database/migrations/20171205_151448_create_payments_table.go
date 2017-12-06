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
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable("payments","InnoDB","utf8");
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(11)").SetUnsigned(true)
	m.NewCol("order_id").SetDataType("INT(11)").SetNullable(false).SetUnsigned(true)
	m.NewCol("account_number").SetDataType("varchar(255)").SetNullable(false)	
	m.NewCol("bank").SetDataType("varchar(255)").SetNullable(false)
}

// Reverse the migrations
func (m *CreatePaymentsTable_20171205_151448) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE payments")
}
