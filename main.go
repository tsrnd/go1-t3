package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/goweb3/routers"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	host := beego.AppConfig.String("postgres_host")
	port, _ := beego.AppConfig.Int("postgres_port")
	username := beego.AppConfig.String("postgres_user")
	pass := beego.AppConfig.String("postgres_pass")
	dbname := beego.AppConfig.String("postgres_dbname")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, pass, dbname)
	orm.RegisterDataBase("default", "postgres", psqlInfo)
}

func main() {
	beego.Run()
}
