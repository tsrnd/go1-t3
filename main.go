package main

import (
	_ "github.com/goweb3/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"	
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
    orm.RegisterDataBase("default", 
        "postgres",
		"user=default password=secret host=127.0.0.1 port=5432 dbname=default sslmode=disable");
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