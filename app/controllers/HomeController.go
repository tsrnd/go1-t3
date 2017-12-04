package controller

import (
	"net/http"

	"github.com/goweb3/app/shared/view"
	"github.com/jianfengye/web-golang/web/session"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.SessionStart(r, w)
	userName := sess.Get("name")
	v := view.New(r)
	v.Vars["name"] = userName
	v.Name = "home/index"
	v.Render(w)
}
