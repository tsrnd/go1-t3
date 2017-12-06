package utils

import (
	"github.com/astaxie/beego/session"
	"github.com/astaxie/beego"	
)

type Session struct {
	SessionStore session.Store
}

func (s *Session) GetSession(controller *session.Store) {
	s.SessionStore = controller.StartSession()
}
