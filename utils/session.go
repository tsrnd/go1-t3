package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

type SessionUtil struct {
	SessionStore *session.Store
}
var (
	Session session.Store
)
func (s *SessionUtil) GetSession(controller beego.Controller) {
	sess := controller.StartSession()
	Session = sess
}
