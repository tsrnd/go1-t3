package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateProductImagesTable_20171205_150729 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateProductImagesTable_20171205_150729{}
	m.Created = "20171205_150729"

	migration.Register("CreateProductImagesTable_20171205_150729", m)
}

// Run the migrations
func (m *CreateProductImagesTable_20171205_150729) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable("product_images","InnoDB","utf8");
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(11)").SetUnsigned(true)
	m.NewCol("product_id").SetDataType("INT(11)").SetNullable(false).SetUnsigned(true)
	m.NewCol("image").SetDataType("text").SetNullable(true)
}

// Reverse the migrations
func (m *CreateProductImagesTable_20171205_150729) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE product_images")
}
