package main

import (
	_ "github.com/goweb3/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
)
var globalSessions *session.Manager

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
    orm.RegisterDataBase("default", 
        "postgres",
		"user=default password=secret host=127.0.0.1 port=5432 dbname=default sslmode=disable");
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()
}

func main() {
	beego.Run()
}