package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// A Filter that runs before the controller
// Filter to protect user pages
var UserIsLogged = func(ctx *context.Context) {
    sess, _ := beego.GlobalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	// read the session from the request
	ses := sess.Get("auth")
    if ses != nil {
		flash := beego.NewFlash()
		flash.Warning("you are logged!")
		storeFlash(flash, ctx)
		ctx.Redirect(302, "/")
    }
}