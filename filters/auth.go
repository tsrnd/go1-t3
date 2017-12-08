package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// A Filter that runs before the controller
// Filter to protect user pages
var ProtectUserPages = func(ctx *context.Context) {
    sess, _ := beego.GlobalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
    defer sess.SessionRelease(ctx.ResponseWriter)
    // read the session from the request
    ses := sess.Get("auth")
    if ses == nil {
		ctx.Redirect(302, "/login")
    }
}