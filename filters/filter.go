package filters

import (
	"net/url"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)


func storeFlash(fd *beego.FlashData, ctx *context.Context) {
	var flashValue string
	for key, value := range fd.Data {
		flashValue += "\x00" + key + "\x23" + beego.BConfig.WebConfig.FlashSeparator + "\x23" + value + "\x00"
	}
	ctx.SetCookie(beego.BConfig.WebConfig.FlashName, url.QueryEscape(flashValue), 0, "/")
}