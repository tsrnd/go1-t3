package models

import (
	"github.com/astaxie/beego/orm"
	
)
type Category struct {
	Id uint
	Name string
}

func init() {
    orm.RegisterModel(new(Category))
}
