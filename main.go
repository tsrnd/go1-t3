package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/goweb3/routers"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default",
		"postgres",
		"user=default password=secret host=127.0.0.1 port=5432 dbname=default sslmode=disable")
}

func main() {
	beego.Run()
}
