package main

import (
	"fmt"
	"github.com/astaxie/beego/migration"
	"github.com/astaxie/beego"	
)

// DO NOT MODIFY
type CreateUserTable_20171205_114942 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUserTable_20171205_114942{}
	m.Created = "20171205_114942"

	migration.Register("CreateUserTable_20171205_114942", m)
}

// Run the migrations
func (m *CreateUserTable_20171205_114942) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable("users","InnoDB","utf8");
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(11)").SetUnsigned(true)
	m.NewCol("name").SetDataType("VARCHAR(45) COLLATE utf8_unicode_ci").SetNullable(false)
	m.NewCol("email").SetDataType("VARCHAR(45) COLLATE utf8_unicode_ci").SetNullable(false)
	m.NewCol("password").SetDataType("VARCHAR(255) COLLATE utf8_unicode_ci").SetNullable(false)
	m.Migrate("create")
}

// Reverse the migrations
func (m *CreateUserTable_20171205_114942) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE users")
}
