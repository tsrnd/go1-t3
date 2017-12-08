package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	_ "github.com/goweb3/routers"

	"github.com/goweb3/bootstrap"
)

func init() {

	bootstrap.AutoloadEnv()

	orm.RegisterDriver(beego.AppConfig.String("database_driver"), orm.DRPostgres)
	host := beego.AppConfig.String("database_host")
	port, _ := beego.AppConfig.Int("database_port")
	username := beego.AppConfig.String("database_user")
	pass := beego.AppConfig.String("database_password")
	dbname := beego.AppConfig.String("database_name")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, pass, dbname)
	orm.RegisterDataBase("default", "postgres", psqlInfo)

	sessionconf := &session.ManagerConfig{
		CookieName: "begoosessionID",
		Gclifetime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()
}

func main() {
	beego.Run()
}
