package utils

import (
	"github.com/astaxie/beego"
)

type SessionUtil struct {}

var (
	Controller *beego.Controller
)
func (s *SessionUtil) SetControler(c beego.Controller) {
	Controller = &c
}
